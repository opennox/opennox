//go:build !server

package ail

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/noxworld-dev/opennox-lib/ifs"
)

func openWav(path string) (*wavReader, error) {
	f, err := ifs.Open(path)
	if err != nil {
		return nil, err
	}
	return newWavReader(f, path)
}

func newWavReader(f io.ReadSeekCloser, name string) (*wavReader, error) {
	r := &wavReader{
		file:     f,
		filename: name,
	}
	if err := r.open(); err != nil {
		_ = f.Close()
		return nil, err
	}
	return r, nil
}

type audioDecoder interface {
	Format() string
	Channels() int
	SampleRate() int
	Position() int
	Seek(pos int)
	Decode(out []int16) (int, bool)
	Close()
}

type waveFormat struct {
	format         uint16
	channels       uint16
	samplesPerSec  uint32
	avgBytesPerSec uint32
	blockAlign     uint16
}

type wavReader struct {
	file      io.ReadSeekCloser
	filename  string
	fileSize  int64
	chunk     *bufio.Reader
	chunkSize int
	chunkPos  int
	dec       audioDecoder
}

func (r *wavReader) Name() string {
	return r.filename
}

func (r *wavReader) Format() string {
	return "WAV+" + r.dec.Format()
}

func (r *wavReader) Channels() int {
	return r.dec.Channels()
}

func (r *wavReader) SampleRate() int {
	return r.dec.SampleRate()
}

func (r *wavReader) Position() int {
	return r.dec.Position()
}

func (r *wavReader) Close() {
	r.dec.Close()
	r.file.Close()
}

func (r *wavReader) open() error {
	var tmp [256]byte
	_, err := io.ReadFull(r.file, tmp[:20])
	if err != nil {
		return err
	}
	if magic := string(tmp[0:4]); magic != "RIFF" {
		return fmt.Errorf("invalid audio file magic: %q", magic)
	}
	r.fileSize = int64(binary.LittleEndian.Uint32(tmp[4:8]))
	if str := string(tmp[8:12]); str != "WAVE" {
		return fmt.Errorf("invalid audio file section: %q", str)
	}
	// +12
	size := int(binary.LittleEndian.Uint32(tmp[16:20]))
	if str := string(tmp[12:16]); str != "fmt " {
		return fmt.Errorf("invalid audio file section 2: %q", str)
	}
	if size < 14 || size > cap(tmp) {
		return fmt.Errorf("invalid audio header size: %d", size)
	}
	_, err = io.ReadFull(r.file, tmp[:size])
	if err != nil {
		return err
	}
	hdr := tmp[:size]
	wf := waveFormat{
		format:         binary.LittleEndian.Uint16(hdr[0:2]),
		channels:       binary.LittleEndian.Uint16(hdr[2:4]),
		samplesPerSec:  binary.LittleEndian.Uint32(hdr[4:8]),
		avgBytesPerSec: binary.LittleEndian.Uint32(hdr[8:12]),
		blockAlign:     binary.LittleEndian.Uint16(hdr[12:14]),
	}
	var dec audioDecoder
	switch wf.format {
	case 0x01: // PCM
		dec, err = newPCMDecoder(r, r.file, wf, hdr)
	case 0x11: // ADPCM
		dec, err = newADPCMDecoder(r, r.file, wf, hdr)
	case 0x55: // MP3
		dec, err = newMP3Decoder(r, r.file, wf, hdr)
	default:
		return fmt.Errorf("unsupported format: 0x%x", wf.format)
	}
	if err != nil {
		return err
	}
	r.dec = dec
	return nil
}

func (r *wavReader) openChunk(size int) {
	r.chunkSize = size
	cr := io.LimitReader(r.file, int64(size))
	if r.chunk == nil {
		r.chunk = bufio.NewReader(cr)
	} else {
		r.chunk.Reset(cr)
	}
}

func (r *wavReader) NextChunk() (int, error) {
	var tmp [8]byte

	r.chunkSize = 0
	r.chunkPos = 0

	for {
		_, err := io.ReadFull(r.file, tmp[:])
		if err != nil {
			if err != io.EOF {
				audioLog.Println("find", err)
			}
			return 0, err
		}
		size := int(binary.LittleEndian.Uint32(tmp[4:]))
		if string(tmp[:4]) == "data" {
			r.openChunk(size)
			return size, nil
		}
		_, err = r.file.Seek(int64(size), io.SeekCurrent)
		if err != nil {
			audioLog.Println(err)
			return 0, err
		}
	}
}

func (r *wavReader) ChunkSize() (int, error) {
	if sz := r.Remaining(); sz != 0 {
		return sz, nil
	}
	return r.NextChunk()
}

func (r *wavReader) Read(data []byte) (int, error) {
	sz, err := r.ChunkSize()
	if err != nil {
		return 0, err
	} else if sz == 0 {
		return 0, io.EOF
	}
	if len(data) > sz {
		data = data[:sz]
	}
	n, err := r.chunk.Read(data)
	r.chunkPos += n
	return n, err
}

func (r *wavReader) ReadFull(data []byte) (int, error) {
	return io.ReadFull(r, data)
}

func (r *wavReader) Remaining() int {
	n := r.chunkSize - r.chunkPos
	if n < 0 {
		n = 0
	}
	return n
}

func (r *wavReader) Seek(offset int) {
	r.chunkSize = 0
	r.chunkPos = 0
	r.file.Seek(12, io.SeekStart)
	r.dec.Seek(offset)
}

func (r *wavReader) Skip(offset int) {
	if offset <= 0 {
		return
	}
	r.chunkPos += offset
	r.file.Seek(int64(offset), io.SeekCurrent)
}

func (r *wavReader) Decode(out []int16) (int, bool) {
	return r.dec.Decode(out)
}
