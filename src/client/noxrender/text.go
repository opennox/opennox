package noxrender

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/opennox/libs/datapath"
	"github.com/opennox/libs/ifs"
	"github.com/opennox/libs/noxfont"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/opennox/opennox/v1/legacy/common/alloc/handles"
)

type fontFile struct {
	Name    string
	File    string
	FileAlt string
	Size    int
	Font    font.Face
	Ptr     unsafe.Pointer
}

var noxFontFiles = []fontFile{
	{Name: noxfont.DefaultName, File: noxfont.DefaultFile, Size: 10},
	{Name: noxfont.LargeName, File: noxfont.LargeFile, FileAlt: noxfont.DefaultFile, Size: 16},
	{Name: noxfont.NumbersName, File: noxfont.NumbersFile, Size: 7},
	{Name: noxfont.SmallName, File: noxfont.SmallFile, Size: 9},
}

type RenderFonts struct {
	byName map[string]*fontFile
	byPtr  map[unsafe.Pointer]*fontFile
	def    font.Face
}

func (r *RenderFonts) Load(lang int) error {
	r.byName = make(map[string]*fontFile)
	r.byPtr = make(map[unsafe.Pointer]*fontFile)
	for i := range noxFontFiles {
		f := &noxFontFiles[i]
		r.byName[f.Name] = f
		fname := f.File
		switch lang {
		case 6, 8:
			if f.FileAlt != "" {
				fname = f.FileAlt
			}
		}
		fnt, err := loadFont(datapath.Data(fname), f.Size)
		if err != nil {
			return fmt.Errorf("cannot load font: %w", err)
		}
		f.Font = fnt
		f.Ptr = handles.NewPtr()
		r.byPtr[f.Ptr] = f
	}
	r.def = r.FontByName(noxfont.DefaultName)
	return nil
}

func (r *RenderFonts) Free() {
	for i := range noxFontFiles {
		f := &noxFontFiles[i]
		if f.Font != nil {
			f.Font.Close()
			f.Font = nil
		}
		f.Ptr = nil
	}
}

func (r *RenderFonts) DefaultFont() font.Face {
	return r.def
}

func (r *RenderFonts) FontByName(name string) font.Face {
	name = strings.ToLower(name)
	f := r.byName[name]
	if f == nil {
		Log.Printf("font not found: %q", name)
		return nil
	}
	return f.Font
}

func (r *RenderFonts) AsFont(ptr unsafe.Pointer) font.Face {
	if ptr == nil {
		return nil
	}
	fnt := r.byPtr[ptr]
	if fnt == nil {
		return nil
	}
	return fnt.Font
}

func (r *RenderFonts) FontPtrByName(name string) unsafe.Pointer {
	name = strings.ToLower(name)
	f := r.byName[name]
	if f == nil {
		Log.Printf("font not found: %q", name)
		return nil
	}
	return f.Ptr
}

func loadFont(path string, size int) (font.Face, error) {
	f, err := ifs.Open(path + ".ttf")
	if os.IsNotExist(err) {
		f, err = ifs.Open(path + ".otf")
	}
	if err == nil {
		defer f.Close()
		// Keep the font data in memory instead of retaining an open file.
		data, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", filepath.Base(f.Name()), err)
		}
		fnt, err := opentype.Parse(data)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", filepath.Base(f.Name()), err)
		}
		face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
			Size: float64(size), DPI: 72,
			Hinting: font.HintingNone,
		})
		if err != nil {
			return nil, err
		}
		return face, nil
	}
	f, err = ifs.Open(path + noxfont.Ext)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fnt, err := noxfont.Decode(f)
	if err != nil {
		return nil, err
	}
	return fnt, nil
}

type noxRenderText struct {
	face *renderFace
	font.Drawer
	tabWidth  int
	lineStart int
	advance   int
	useBold   bool
	smooth    bool
}

func (r *NoxRender) initText() {
	r.text.tabWidth = 64
	r.text.face = &renderFace{r: r}
	r.text.Face = r.text.face
	r.text.useBold = true
}

func (r *NoxRender) SetTextSmooting(enabled bool) {
	r.text.smooth = enabled
}

func (r *NoxRender) SetBold(enable bool) {
	r.text.useBold = enable
}

func (r *NoxRender) TabWidth() int {
	return r.text.tabWidth
}

func (r *NoxRender) SetTabWidth(w int) {
	r.text.tabWidth = w
}

func (r *NoxRender) fontOrDefault(fnt font.Face) font.Face {
	if fnt != nil {
		return fnt
	}
	return r.Fonts.DefaultFont()
}

func (r *NoxRender) FontHeight(fnt font.Face) int {
	fnt = r.fontOrDefault(fnt)
	if fnt == nil {
		return 0
	}
	return fnt.Metrics().CapHeight.Round()
}

func (r *NoxRender) GetFonts() *RenderFonts {
	return &r.Fonts
}

func (r *NoxRender) GetBag() *RenderSprites {
	return &r.Bag
}

func (r *NoxRender) DrawString(font font.Face, str string, pos image.Point) int { // nox_xxx_drawString_43F6E0
	if !r.p.ShouldDrawText() {
		return pos.X
	}
	font = r.fontOrDefault(font)
	if font == nil {
		return pos.X
	}
	r.text.lineStart = pos.X
	return r.drawStringLine(font, str, pos)
}

func (r *NoxRender) DrawStringStyle(font font.Face, str string, pos image.Point) int { // nox_xxx_drawStringStyle_43F7B0
	if !r.p.ShouldDrawText() {
		return pos.X
	}
	font = r.fontOrDefault(font)
	if font == nil {
		return pos.X
	}
	r.text.lineStart = pos.X
	if !r.text.useBold {
		return r.drawStringLine(font, str, pos)
	}
	r.text.advance = 1
	r.drawStringLine(font, str, pos)
	x := r.drawStringLine(font, str, pos.Add(image.Point{X: 1}))
	r.text.advance = 0
	return x
}

func (r *NoxRender) DrawStringHL(font font.Face, str string, pos image.Point) int { // nox_draw_drawStringHL_43F730
	if !r.p.ShouldDrawText() {
		return pos.X
	}
	font = r.fontOrDefault(font)
	if font == nil {
		return pos.X
	}
	r.text.lineStart = pos.X
	old := r.p.TextColor()
	r.p.SetTextColor(r.p.Color())
	r.drawStringLine(font, str, pos.Add(image.Point{X: 1, Y: 1}))
	r.p.SetTextColor(old)
	return r.drawStringLine(font, str, pos)
}

func (r *NoxRender) drawStringLine(font font.Face, str string, pos image.Point) int {
	str = strings.NewReplacer(
		"\n", "",
		"\r", "",
	).Replace(str)
	return r.drawString(font, str, pos)
}

func (r *NoxRender) drawString(fnt font.Face, s string, pos image.Point) int {
	// FIXME: handle tab characters properly
	if fnt == nil {
		return pos.X
	}
	// TODO: handle invalid chars by blinking the replacement char
	dy := fnt.Metrics().CapHeight.Round()
	r.text.face.f = fnt
	r.text.Src = image.NewUniform(r.p.TextColor())
	r.text.Dst = r.pix.SubImage(r.p.ClipRect())
	r.text.Dot = fixed.P(pos.X, pos.Y+dy)
	r.text.DrawString(s)
	return r.text.Dot.X.Round()
}

func (r *NoxRender) DrawStringWrapped(font font.Face, s string, rect image.Rectangle) int { // nox_xxx_drawStringWrap_43FAF0
	if !r.p.ShouldDrawText() {
		return rect.Min.X
	}
	dot, _ := r.doStringWrapped(font, s, rect, r.drawString, nil)
	return dot.X
}

func (r *NoxRender) DrawStringWrappedHL(font font.Face, s string, rect image.Rectangle) int { // nox_xxx_drawStringWrapHL_43FD00
	if !r.p.ShouldDrawText() {
		return rect.Min.X
	}
	r.text.lineStart = rect.Min.X
	old := r.p.TextColor()
	r.p.SetTextColor(r.p.Color())
	pt := image.Point{X: 1, Y: 1}
	rect2 := rect
	rect2.Min = rect2.Min.Add(pt)
	rect2.Max = rect2.Max.Add(pt)
	r.DrawStringWrapped(font, s, rect2)
	r.p.SetTextColor(old)
	return r.DrawStringWrapped(font, s, rect)
}

func (r *NoxRender) getStringRuneAdvance(fnt font.Face, c rune) int {
	r.text.face.f = fnt
	dx, ok := r.text.face.GlyphAdvance(c)
	if !ok {
		dx, ok = r.text.face.GlyphAdvance('?')
	}
	return dx.Round()
}

type doStringFunc func(f font.Face, s string, p image.Point) int

func (r *NoxRender) doStringWrapped(font font.Face, s string, rect image.Rectangle, wordFnc doStringFunc, lineFnc func(s string)) (dot image.Point, bb image.Rectangle) {
	dot = rect.Min
	bb.Min = rect.Min
	bb.Max = rect.Min
	font = r.fontOrDefault(font)
	if font == nil {
		return
	}
	if s == "" {
		return
	}

	dy := font.Metrics().CapHeight.Round()
	r.text.lineStart = rect.Min.X
	str := []rune(s)

	var (
		wordP image.Point // relative to rect.Min
		word  []rune      // current word
		line  []rune      // current line
		wordN int         // words in the current line
		addX  int         // relative to wordX
	)

	procWord := func() {
		cp := rect.Min.Add(wordP)
		cp.X = wordFnc(font, string(word), cp)
		if cp.X > bb.Max.X {
			bb.Max.X = cp.X
		}
		line = append(line, word...)
		word = word[:0]
		wordP.X = cp.X - rect.Min.X
		addX = 0
		wordN++
	}

	maxW := rect.Dx()
	maxH := rect.Dy()

	for i, c := range str {
		switch c {
		case '\t':
			procWord()
			line = append(line, c)
			tab := r.text.tabWidth
			wordP.X += tab
			wordP.X -= wordP.X % tab
			continue
		case '\r', '\n':
			if i != 0 && str[i-1] == '\r' {
				continue
			}
			procWord()
			wordP.X = 0
			wordN = 0
			wordP.Y += dy
			if lineFnc != nil {
				lineFnc(string(line))
				line = line[:0]
			}
			if maxH > 0 && wordP.Y >= maxH {
				bb.Max.Y = rect.Max.Y + dy
				return rect.Min.Add(wordP), bb
			}
			continue
		}
		if maxW > 0 {
			dx := r.getStringRuneAdvance(font, c)
			if dx == 0 {
				continue
			}
			addX += dx
			if wordN != 0 && wordP.X+addX > maxW {
				// word is too long, but not the first in a line
				if lineFnc != nil {
					lineFnc(string(line))
					line = line[:0]
				}
				wordP.X = 0
				wordN = 0
				wordP.Y += dy
				if maxH > 0 && wordP.Y >= maxH {
					bb.Max.Y = rect.Max.Y + dy
					return rect.Min.Add(wordP), bb
				}
			}
		}
		word = append(word, c)
		if c == ' ' {
			procWord()
		}
	}
	procWord()
	if lineFnc != nil && len(line) != 0 {
		lineFnc(string(line))
		line = line[:0]
	}
	bb.Max.Y += wordP.Y + dy
	return rect.Min.Add(wordP), bb
}

func (r *NoxRender) DrawStringWrappedStyle(font font.Face, s string, rect image.Rectangle) int { // nox_xxx_bookDrawString_43FA80_43FD80
	if !r.p.ShouldDrawText() {
		return rect.Min.X
	}
	r.text.lineStart = rect.Min.X
	if !r.text.useBold {
		return r.DrawStringWrapped(font, s, rect)
	}
	// draw bold by repeating the call twice with a swift
	r.text.advance = 1
	r.DrawStringWrapped(font, s, rect)
	rect.Min.X += 1
	x := r.DrawStringWrapped(font, s, rect)
	r.text.advance = 0
	return x
}

func (r *NoxRender) SplitStringWrapped(fnt font.Face, s string, maxW int) []string {
	rect := image.Rect(0, 0, maxW, 0)
	var lines []string
	r.doStringWrapped(fnt, s, rect, r.advanceString, func(line string) {
		lines = append(lines, line)
	})
	return lines
}

func (r *NoxRender) advanceString(fnt font.Face, s string, p image.Point) int {
	x := p.X
	for _, c := range s {
		x += r.getStringRuneAdvance(fnt, c)
	}
	return x
}

func (r *NoxRender) GetStringSizeWrapped(fnt font.Face, s string, maxW int) image.Point { // nox_xxx_drawGetStringSize_43F840
	rect := image.Rect(0, 0, maxW, 0)
	_, bb := r.doStringWrapped(fnt, s, rect, r.advanceString, nil)
	return image.Point{X: bb.Dx(), Y: bb.Dy()}
}

func (r *NoxRender) GetStringSizeWrappedStyle(fnt font.Face, s string, maxW int) image.Point { // nox_xxx_bookGetStringSize_43FA80
	if !r.text.useBold {
		return r.GetStringSizeWrapped(fnt, s, maxW)
	}
	r.text.advance = 1
	sz := r.GetStringSizeWrapped(fnt, s, maxW)
	r.text.advance = 0
	return sz
}

type renderFace struct {
	r *NoxRender
	f font.Face
}

func (f *renderFace) Close() error {
	return f.f.Close()
}

func (f *renderFace) Glyph(dot fixed.Point26_6, r rune) (dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int26_6, ok bool) {
	dr, mask, maskp, advance, ok = f.f.Glyph(dot, r)
	if ok {
		advance += fixed.I(f.r.text.advance)
	}
	return
}

func (f *renderFace) GlyphBounds(r rune) (bounds fixed.Rectangle26_6, advance fixed.Int26_6, ok bool) {
	bounds, advance, ok = f.f.GlyphBounds(r)
	if ok {
		advance += fixed.I(f.r.text.advance)
	}
	return
}

func (f *renderFace) GlyphAdvance(r rune) (advance fixed.Int26_6, ok bool) {
	advance, ok = f.f.GlyphAdvance(r)
	if ok {
		advance += fixed.I(f.r.text.advance)
	}
	return
}

func (f *renderFace) Kern(r0, r1 rune) fixed.Int26_6 {
	return f.f.Kern(r0, r1)
}

func (f *renderFace) Metrics() font.Metrics {
	return f.f.Metrics()
}
