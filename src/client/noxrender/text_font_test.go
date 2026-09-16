package noxrender

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/image/font/gofont/goregular"
)

func TestLoadFontClosesFile(t *testing.T) {
	for _, ext := range []string{".ttf", ".otf"} {
		t.Run(ext, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "default")
			require.NoError(t, os.WriteFile(path+ext, goregular.TTF, 0600))
			face, err := loadFont(path, 10)
			require.NoError(t, err)
			defer face.Close()

			// Windows refuses removal while the source file is still open.
			require.NoError(t, os.Remove(path+ext))
			advance, ok := face.GlyphAdvance('A')
			require.True(t, ok, "the font must remain usable after removing its source")
			require.Positive(t, advance)
		})
	}
}
