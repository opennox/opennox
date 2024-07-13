//go:build !server

package ail

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/shoenig/test/must"
	"golang.org/x/exp/maps"
)

type testFile struct {
	hash    string
	size    int
	sizeExp int
}

func TestAudioDecode(t *testing.T) {
	cases := []struct {
		name string
		exp  map[string]testFile
	}{
		{
			name: "dialog",
			exp:  dialogHashes,
		},
		{
			name: "music",
			exp:  musicHashes,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path := noxtest.DataPath(t, c.name)
			list, err := ifs.ReadDir(path)
			must.NoError(t, err)
			all := make(map[string]testFile)
			t.Cleanup(func() {
				keys := maps.Keys(all)
				sort.Strings(keys)
				var buf strings.Builder
				for _, name := range keys {
					v := all[name]
					fmt.Fprintf(&buf, "\t%q: {hash: %q, size: %d},\n", name, v.hash, v.size)
				}
				t.Logf("\n%s", buf.String())
			})
			for _, fi := range list {
				if fi.IsDir() {
					continue
				}
				name := strings.ToLower(fi.Name())
				t.Run(name, func(t *testing.T) {
					path := filepath.Join(path, name)
					f, err := ifs.Open(path)
					must.NoError(t, err)
					f.Close()
					path = f.Name()

					r, err := openAudio(path)
					must.NoError(t, err)
					defer r.Close()

					h := sha256.New()
					rate := r.SampleRate()
					channels := r.Channels()
					t.Logf("%s, rate: %d, channels: %d", r.Format(), rate, channels)
					buf := make([]int16, rate)
					total := 0
					for {
						fail := false
						if total > 5*60*rate*channels {
							fail = true
						}
						n, ok := r.Decode(buf)
						total += n
						if n != 0 {
							hashPCM(h, buf[:n])
						}
						if !ok {
							break
						}
						if fail {
							t.Fatal("too many samples")
						}
					}
					got := hex.EncodeToString(h.Sum(nil))
					gotSz := total * 2
					expf := c.exp[name]
					sizeFF := expf.size
					if sizeFF == 0 {
						_, sizeFF = hashPCMFile(t, path)
					}
					if expf.hash != got || expf.size != sizeFF {
						all[name] = testFile{hash: got, size: sizeFF}
					}
					if expf.hash != got {
						t.Errorf("hash missmatch: exp %s, got %s", expf.hash, got)
					}
					sizeExp := expf.sizeExp
					if sizeExp == 0 {
						sizeExp = sizeFF
					}
					if sizeExp != 0 {
						must.EqOp(t, sizeExp, gotSz)
					}
				})
			}
		})
	}
}

func hashPCM(h hash.Hash, buf []int16) int {
	for _, v := range buf {
		var b [2]byte
		binary.LittleEndian.PutUint16(b[:], uint16(v))
		h.Write(b[:])
	}
	return 2 * len(buf)
}

func writePCM(t testing.TB, path string, buf []int16) {
	f, err := os.Create(path)
	must.NoError(t, err)
	defer f.Close()

	bw := bufio.NewWriter(f)
	defer bw.Flush()

	for _, v := range buf {
		var b [2]byte
		binary.LittleEndian.PutUint16(b[:], uint16(v))
		bw.Write(b[:])
	}
}

func hashPCMFile(t testing.TB, path string) (string, int) {
	bin, err := exec.LookPath("ffmpeg")
	if err != nil {
		return "", 0
	}
	f, err := os.CreateTemp("", "noxaudio_*.s16le")
	must.NoError(t, err)
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	cmd := exec.Command(bin, "-y", "-i", path, "-f", "s16le", f.Name())
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	must.NoError(t, err)
	h := sha256.New()
	n, err := io.Copy(h, f)
	must.NoError(t, err)
	return hex.EncodeToString(h.Sum(nil)), int(n)
}
