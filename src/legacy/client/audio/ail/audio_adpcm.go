//go:build !server

package ail

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

func newADPCMDecoder(r *wavReader, hr io.Reader, wf waveFormat, hdr []byte) (*adpcmDecoder, error) {
	d := &adpcmDecoder{
		r:          r,
		sampleRate: int(wf.samplesPerSec),
		channels:   int(wf.channels),
		blockSize:  int(wf.blockAlign),
	}
	var buf [12]byte
	_, err := io.ReadFull(hr, buf[:12])
	if err != nil {
		return nil, err
	}
	if str := string(buf[0:4]); str != "fact" {
		return nil, fmt.Errorf("invalid ADPCM file section: %q", str)
	}
	if sz := binary.LittleEndian.Uint32(buf[4:8]); sz != 4 {
		return nil, fmt.Errorf("unsupported ADPCM size: %q", sz)
	}
	d.samples = int(binary.LittleEndian.Uint32(buf[8:12]))
	audioLog.Printf("ADPCM stream: %q, %d channels", r.Name(), wf.channels)
	return d, nil
}

type adpcmDecoder struct {
	r          *wavReader
	blockSize  int
	position   int
	samples    int
	sampleRate int
	channels   int

	buffered int
	buffer   [16 * 1024 * 3]byte
}

func (*adpcmDecoder) Format() string {
	return "ADPCM"
}

func (d *adpcmDecoder) Channels() int {
	return d.channels
}

func (d *adpcmDecoder) SampleRate() int {
	return d.sampleRate
}

func (d *adpcmDecoder) Close() {}

func (d *adpcmDecoder) Decode(out []int16) (int, bool) {
	blockSize := d.blockSize

	samples := 0
	if d.position >= d.samples {
		return 0, false
	}

	hasErr := false
	if d.buffered < blockSize {
		remaining, _ := d.r.ChunkSize()
		if remaining == 0 {
			return 0, false
		}
		if remaining > blockSize-d.buffered {
			remaining = blockSize - d.buffered
		}
		n, err := d.r.ReadFull(d.buffer[d.buffered : d.buffered+remaining])
		d.buffered += n
		hasErr = err != nil
	}

	if d.channels == 1 {
		out = decodeADPCMMono(out[:0], d.buffer[:blockSize])
	} else {
		out = decodeADPCMStereo(out[:0], d.buffer[:blockSize])
	}
	samples = len(out) // TODO(dennwc): is it true for stereo?

	d.buffered -= blockSize
	if d.buffered != 0 {
		copy(d.buffer[0:], d.buffer[blockSize:blockSize+d.buffered])
	}
	if samples/d.channels+d.position >= d.samples {
		samples = (d.samples - d.position) * d.channels
	}
	d.position += samples / d.channels
	if samples == 0 && hasErr {
		return 0, false
	}
	return samples, true
}

func (d *adpcmDecoder) Seek(position int) {
	d.buffered = 0
	blockSize := d.blockSize
	samplesPerBlock := (blockSize/d.channels - 4) * 2
	blocks := position / samplesPerBlock
	d.position = 0

	for d.position < position {
		csz, _ := d.r.NextChunk()
		if csz == 0 {
			break
		}
		chunkBlocks := csz / blockSize
		if blocks < chunkBlocks {
			d.r.Skip(blocks * blockSize)
			d.position += blocks * samplesPerBlock
			break
		} else {
			d.r.Skip(csz)
			d.position += chunkBlocks * samplesPerBlock
		}
	}

	// TODO seek within an ADPCM block
}

func (d *adpcmDecoder) Position() int {
	return d.position
}

var (
	imaIndexTable = [16]int{-1, -1, -1, -1, 2, 4, 6, 8, -1, -1, -1, -1, 2, 4, 6, 8}
	imaStepTable  = [89]int{
		7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 19, 21, 23,
		25, 28, 31, 34, 37, 41, 45, 50, 55, 60, 66, 73, 80,
		88, 97, 107, 118, 130, 143, 157, 173, 190, 209, 230, 253, 279,
		307, 337, 371, 408, 449, 494, 544, 598, 658, 724, 796, 876, 963,
		1060, 1166, 1282, 1411, 1552, 1707, 1878, 2066, 2272, 2499, 2749, 3024, 3327,
		3660, 4026, 4428, 4871, 5358, 5894, 6484, 7132, 7845, 8630, 9493, 10442, 11487,
		12635, 13899, 15289, 16818, 18500, 20350, 22385, 24623, 27086, 29794, 32767,
	}
)

func decodeADPCMMono(out []int16, data []byte) []int16 {
	predictor := int16(data[0]) | (int16(data[1]) << 8)
	index := int(data[2])

	out = append(out, predictor)
	for i := 4; i < len(data); i++ {
		predictor = decodeNibble(predictor, data[i], index)
		index = satindex(index + imaIndexTable[data[i]&15])
		out = append(out, predictor)

		predictor = decodeNibble(predictor, data[i]>>4, index)
		index = satindex(index + imaIndexTable[data[i]>>4])
		out = append(out, predictor)
	}
	return out
}

func sat16(x int) int16 {
	if x < math.MinInt16 {
		return math.MinInt16
	} else if x > math.MaxInt16 {
		return math.MaxInt16
	}
	return int16(x)
}

func decodeNibble(predictor int16, nibble byte, index int) int16 {
	diff := 0
	step := imaStepTable[index]

	diff = step >> 3
	if nibble&4 != 0 {
		diff += step
	}
	if nibble&2 != 0 {
		diff += step >> 1
	}
	if nibble&1 != 0 {
		diff += step >> 2
	}
	if nibble&8 != 0 {
		predictor = sat16(int(predictor) - diff)
	} else {
		predictor = sat16(int(predictor) + diff)
	}
	return predictor
}

func satindex(x int) int {
	if x < 0 {
		return 0
	} else if x > 88 {
		return 88
	}
	return x
}

func decodeADPCMStereo(out []int16, data []byte) []int16 {
	lpredictor := int16(data[0]) | (int16(data[1]) << 8)
	lindex := int(data[2])
	rpredictor := int16(data[4]) | (int16(data[5]) << 8)
	rindex := int(data[6])

	out = append(out, lpredictor, rpredictor)
	for i := 8; i < len(data); i += 8 {
		for j := 0; j < 4; j++ {
			lpredictor = decodeNibble(lpredictor, data[i+j], lindex)
			lindex = satindex(lindex + imaIndexTable[data[i+j]&15])
			out = append(out, lpredictor)

			rpredictor = decodeNibble(rpredictor, data[i+j+4], rindex)
			rindex = satindex(rindex + imaIndexTable[data[i+j+4]&15])
			out = append(out, rpredictor)

			lpredictor = decodeNibble(lpredictor, data[i+j]>>4, lindex)
			lindex = satindex(lindex + imaIndexTable[data[i+j]>>4])
			out = append(out, lpredictor)

			rpredictor = decodeNibble(rpredictor, data[i+j+4]>>4, rindex)
			rindex = satindex(rindex + imaIndexTable[data[i+j+4]>>4])
			out = append(out, rpredictor)
		}
	}
	return out
}
