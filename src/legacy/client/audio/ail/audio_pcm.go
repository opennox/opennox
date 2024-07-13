//go:build !server

package ail

import (
	"encoding/binary"
	"io"
)

func newPCMDecoder(r *wavReader, hr io.Reader, wf waveFormat, hdr []byte) (*pcmDecoder, error) {
	d := &pcmDecoder{
		r:          r,
		sampleRate: int(wf.samplesPerSec),
		channels:   int(wf.channels),
	}
	audioLog.Printf("PCM stream: %q, %d channels", r.Name(), wf.channels)
	return d, nil
}

type pcmDecoder struct {
	r          *wavReader
	position   int
	sampleRate int
	buf        []byte
	channels   int
}

func (*pcmDecoder) Format() string {
	return "PCM"
}

func (d *pcmDecoder) Channels() int {
	return d.channels
}

func (d *pcmDecoder) SampleRate() int {
	return d.sampleRate
}

func (d *pcmDecoder) Close() {}

func (d *pcmDecoder) Decode(out []int16) (int, bool) {
	n := len(out) * 2
	if cap(d.buf) < n {
		d.buf = make([]byte, n)
	} else {
		d.buf = d.buf[:n]
	}
	n, err := d.r.Read(d.buf)
	d.buf = d.buf[:n]
	for i := range out[:n/2] {
		out[i] = int16(binary.LittleEndian.Uint16(d.buf[2*i:]))
	}
	d.position += n / 2 / d.channels
	if n == 0 && err != nil {
		return 0, false
	}
	return n / 2, true
}

func (d *pcmDecoder) Seek(position int) {
	// TODO
	d.position = 0
}

func (d *pcmDecoder) Position() int {
	return d.position
}
