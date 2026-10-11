//go:build !server

package ail

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/opennox/libs/ifs"
	"github.com/opennox/libs/noxtest"
	"github.com/shoenig/test/must"
)

type memFile struct {
	*bytes.Reader
}

func (memFile) Close() error { return nil }

func testBlock(ch, blockSize, idx, hdrIdx int) []byte {
	b := make([]byte, blockSize)
	seed := uint32(idx)*2654435761 + 12345
	for i := range b {
		seed = seed*1664525 + 1013904223
		b[i] = byte(seed >> 24)
	}
	for c := 0; c < ch; c++ {
		b[c*4+2] = byte(hdrIdx) // step index in header
		b[c*4+3] = 0
	}
	return b
}

// buildADPCMWav builds a synthetic ADPCM WAV. Data is split into chunks of the given sizes (in bytes);
// the last chunk takes everything that remains.
func buildADPCMWav(ch, blockSize, nblocks, hdrIdx int, chunkSizes []int, truncate int) []byte {
	var data []byte
	for i := 0; i < nblocks; i++ {
		data = append(data, testBlock(ch, blockSize, i, hdrIdx)...)
	}
	perBlock := 1 + (blockSize-4*ch)*2/ch
	total := nblocks * perBlock

	le := binary.LittleEndian
	var body bytes.Buffer
	body.WriteString("WAVE")
	body.WriteString("fmt ")
	var f [20]byte
	le.PutUint16(f[0:], 0x11)
	le.PutUint16(f[2:], uint16(ch))
	le.PutUint32(f[4:], 22050)
	le.PutUint32(f[8:], 11000)
	le.PutUint16(f[12:], uint16(blockSize))
	le.PutUint16(f[14:], 4)
	le.PutUint16(f[16:], 2)
	le.PutUint16(f[18:], uint16(perBlock))
	var sz [4]byte
	le.PutUint32(sz[:], 20)
	body.Write(sz[:])
	body.Write(f[:])
	body.WriteString("fact")
	le.PutUint32(sz[:], 4)
	body.Write(sz[:])
	le.PutUint32(sz[:], uint32(total))
	body.Write(sz[:])
	for _, cs := range append(chunkSizes, len(data)) {
		if cs > len(data) {
			cs = len(data)
		}
		body.WriteString("data")
		le.PutUint32(sz[:], uint32(cs))
		body.Write(sz[:])
		body.Write(data[:cs])
		data = data[cs:]
	}
	out := append([]byte("RIFF\x00\x00\x00\x00"), body.Bytes()...)
	le.PutUint32(out[4:], uint32(len(out)-8))
	if truncate > 0 {
		out = out[:len(out)-truncate]
	}
	return out
}

func openMemWav(t testing.TB, b []byte) (*wavReader, *adpcmDecoder) {
	r, err := newWavReader(memFile{bytes.NewReader(b)}, "mem.wav")
	must.NoError(t, err)
	d, ok := r.dec.(*adpcmDecoder)
	must.True(t, ok)
	return r, d
}

func decodeAll(t testing.TB, r *wavReader) []int16 {
	var all []int16
	buf := make([]int16, 64*1024)
	for i := 0; i < 100000; i++ {
		n, ok := r.Decode(buf)
		all = append(all, buf[:n]...)
		if !ok {
			return all
		}
	}
	t.Fatal("decoder does not terminate")
	return nil
}

func TestADPCMBadHeaderIndex681(t *testing.T) {
	for _, c := range []struct{ ch, bs int }{{1, 256}, {2, 512}} {
		b := buildADPCMWav(c.ch, c.bs, 4, 253, nil, 0)
		r, d := openMemWav(t, b)
		out := decodeAll(t, r)
		perBlock := c.ch * (1 + (c.bs-4*c.ch)*2/c.ch)
		must.EqOp(t, 4*perBlock, len(out))
		must.False(t, d.failed)
		r.Close()
	}
}

func TestADPCMSplitChunks(t *testing.T) {
	for _, c := range []struct{ ch, bs int }{{1, 256}, {2, 512}} {
		one := buildADPCMWav(c.ch, c.bs, 8, 10, nil, 0)
		// split inside blocks
		split := buildADPCMWav(c.ch, c.bs, 8, 10, []int{c.bs + 37, c.bs*2 + 5, 1}, 0)
		r1, _ := openMemWav(t, one)
		exp := decodeAll(t, r1)
		r2, d2 := openMemWav(t, split)
		got := decodeAll(t, r2)
		must.False(t, d2.failed)
		must.SliceNotEmpty(t, exp)
		must.Eq(t, exp, got)
	}
}

func TestADPCMTruncated(t *testing.T) {
	for _, c := range []struct{ ch, bs int }{{1, 256}, {2, 512}} {
		b := buildADPCMWav(c.ch, c.bs, 4, 10, nil, 17)
		r, d := openMemWav(t, b)
		out := decodeAll(t, r)
		perBlock := c.ch * (1 + (c.bs-4*c.ch)*2/c.ch)
		must.EqOp(t, 3*perBlock, len(out))
		must.True(t, d.buffered >= 0)
		must.False(t, d.failed)
		// repeated decode after the end must stay clean
		n, ok := r.Decode(make([]int16, 4096))
		must.EqOp(t, 0, n)
		must.False(t, ok)
		must.True(t, d.buffered >= 0)
	}
}

func TestADPCMInvalidFormat(t *testing.T) {
	for _, c := range []struct{ ch, bs int }{
		{1, 0}, {1, 4}, {1, 65535}, {2, 0}, {2, 8}, {2, 13}, {2, 65535}, {3, 256}, {0, 256},
	} {
		b := buildADPCMWav(1, 256, 2, 10, nil, 0)
		binary.LittleEndian.PutUint16(b[12+8+2:], uint16(c.ch))
		binary.LittleEndian.PutUint16(b[12+8+12:], uint16(c.bs))
		_, err := newWavReader(memFile{bytes.NewReader(b)}, "bad.wav")
		must.Error(t, err, must.Sprintf("ch=%d bs=%d", c.ch, c.bs))
	}
}

func TestADPCMSeek(t *testing.T) {
	for _, c := range []struct{ ch, bs int }{{1, 256}, {2, 512}} {
		// chunks of 3 blocks so that seek must cross chunk boundaries
		b := buildADPCMWav(c.ch, c.bs, 12, 10, []int{3 * c.bs, 3 * c.bs, 3 * c.bs}, 0)
		r, d := openMemWav(t, b)
		exp := decodeAll(t, r)
		perBlock := c.ch * (1 + (c.bs-4*c.ch)*2/c.ch)
		must.EqOp(t, 12*perBlock, len(exp))

		r, d = openMemWav(t, b)
		buf := make([]int16, 64*1024)
		for i := 0; i < 2; i++ { // playback already started
			_, ok := r.Decode(buf)
			must.True(t, ok)
		}
		spb := (c.bs/c.ch - 4) * 2
		const skip = 7
		r.Seek(skip * spb)
		must.EqOp(t, skip*spb, d.Position())
		n, ok := r.Decode(buf)
		must.True(t, ok)
		must.Eq(t, exp[skip*perBlock:skip*perBlock+n], buf[:n])
	}
}

func TestADPCMRealFile(t *testing.T) {
	path := noxtest.DataPath(t, "music", "town1.wav")
	if _, err := ifs.Stat(path); err != nil {
		t.Skip("no game file")
	}
	r, err := openWav(path)
	must.NoError(t, err)
	defer r.Close()
	d, ok := r.dec.(*adpcmDecoder)
	if !ok {
		t.Skipf("not ADPCM: %s", r.Format())
	}
	exp := decodeAll(t, r)
	t.Logf("TOWN1.WAV: %d channels, rate %d, %d samples (%d frames), declared %d frames",
		d.channels, d.sampleRate, len(exp), len(exp)/d.channels, d.samples)
	must.SliceNotEmpty(t, exp)

	// block size in output samples
	r2, err := openWav(path)
	must.NoError(t, err)
	defer r2.Close()
	d2 := r2.dec.(*adpcmDecoder)
	buf := make([]int16, 64*1024)
	perBlock, _ := r2.Decode(buf)
	for i := 0; i < 5; i++ {
		_, ok := r2.Decode(buf)
		must.True(t, ok)
	}
	spb := (d2.blockSize/d2.channels - 4) * 2
	const skip = 20
	r2.Seek(skip * spb)
	n, ok := r2.Decode(buf)
	must.True(t, ok)
	must.Eq(t, exp[skip*perBlock:skip*perBlock+n], buf[:n])
}
