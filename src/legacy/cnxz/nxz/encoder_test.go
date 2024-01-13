//go:build 386

package nxz

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	maps := noxtest.DataPath(t, "maps")
	files, err := os.ReadDir(maps)
	require.NoError(t, err)
	for _, fi := range files {
		mname := filepath.Join(maps, fi.Name(), fi.Name()+".map")
		zname := filepath.Join(maps, fi.Name(), fi.Name()+".nxz")
		if _, err = ifs.Stat(zname); err != nil {
			continue
		}
		t.Run(fi.Name(), func(t *testing.T) {
			mexp, mexpN := hashFile(t, zname)
			gotc, gotcN := compress(t, mname)
			require.Equal(t, mexpN, gotcN)
			require.Equal(t, mexp, gotc)
		})
	}
}

func compress(t testing.TB, path string) (string, int) {
	out, err := os.CreateTemp("", "nxzmap_*.nxz")
	require.NoError(t, err)
	defer func() {
		out.Close()
		_ = os.Remove(out.Name())
	}()
	err = CompressFile(path, out.Name())
	require.NoError(t, err)
	return hashFile(t, out.Name())
}
