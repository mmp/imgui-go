package imgui

// #include "wrapper/Font.h"
// #include "wrapper/Types.h"
import "C"

// Font describes one loaded font in an atlas.
type Font uintptr

// FontGlyph represents a single font glyph from a font atlas.
type FontGlyph uintptr

// DefaultFont can be used to refer to the default font of the current font atlas without
// having the actual font reference.
const DefaultFont Font = 0

// PushFont adds the given font on the stack. Use DefaultFont to refer to the default font.
func PushFont(font Font) {
	C.iggPushFont(font.handle())
}

// PopFont removes the previously pushed font from the stack.
func PopFont() {
	C.iggPopFont()
}

// FontSize returns the current font size (= height in pixels) of the current font with the current scale applied.
func FontSize() float32 {
	return float32(C.iggGetFontSize())
}

// CalcTextSize calculates the size of the text.
func CalcTextSize(text string, hideTextAfterDoubleHash bool, wrapWidth float32) Vec2 {
	CString := newStringBuffer(text)
	defer CString.free()

	var vec2 Vec2
	valueArg, returnFunc := vec2.wrapped()

	C.iggCalcTextSize((*C.char)(CString.ptr), C.int(CString.size)-1, castBool(hideTextAfterDoubleHash), C.float(wrapWidth), valueArg)
	returnFunc()

	return vec2
}

func (font Font) handle() C.IggFont {
	return C.IggFont(font)
}

func (font Font) FindGlyph(ch rune) FontGlyph {
	return FontGlyph(C.iggFindGlyph(font.handle(), C.int(ch)))
}

func (glyph FontGlyph) handle() C.IggFontGlyph {
	return C.IggFontGlyph(glyph)
}

func (glyph FontGlyph) Colored() bool {
	return C.iggFontGlyphColored(glyph.handle()) != 0
}

func (glyph FontGlyph) Visible() bool {
	return C.iggFontGlyphVisible(glyph.handle()) != 0
}

func (glyph FontGlyph) Codepoint() int {
	return int(C.iggFontGlyphCodepoint(glyph.handle()))
}

func (glyph FontGlyph) AdvanceX() float32 {
	return float32(C.iggFontGlyphAdvanceX(glyph.handle()))
}

func (glyph FontGlyph) X0() float32 {
	return float32(C.iggFontGlyphX0(glyph.handle()))
}

func (glyph FontGlyph) Y0() float32 {
	return float32(C.iggFontGlyphY0(glyph.handle()))
}

func (glyph FontGlyph) X1() float32 {
	return float32(C.iggFontGlyphX1(glyph.handle()))
}

func (glyph FontGlyph) Y1() float32 {
	return float32(C.iggFontGlyphY1(glyph.handle()))
}

func (glyph FontGlyph) U0() float32 {
	return float32(C.iggFontGlyphU0(glyph.handle()))
}

func (glyph FontGlyph) V0() float32 {
	return float32(C.iggFontGlyphV0(glyph.handle()))
}

func (glyph FontGlyph) U1() float32 {
	return float32(C.iggFontGlyphU1(glyph.handle()))
}

func (glyph FontGlyph) V1() float32 {
	return float32(C.iggFontGlyphV1(glyph.handle()))
}
