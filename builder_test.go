package thumbor

import (
	"fmt"
)

func ExampleNoSecret() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/foo.png
}

func ExampleWithSecret() {
	b := NewBuilder("https://images.foo.bar", "abc123")
	fmt.Println(b.Image("foo.png"))
	// Output: https://images.foo.bar/SV5Hoh-Z7Gz7CpiT1bjuZuByZTLqlc5FfpxYTUxZxHU=/foo.png
}

func ExampleTrim() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.Trim().Image("foo.png"))
	// Output: http://example.com/unsafe/trim/foo.png
}

func ExampleTrimWithSourceTopLeft() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.TrimWithSource("top-left").Image("foo.png"))
	// Output: http://example.com/unsafe/trim:top-left/foo.png
}

func ExampleTrimWithSourceBottomRight() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.TrimWithSource("bottom-right").Image("foo.png"))
	// Output: http://example.com/unsafe/trim:bottom-right/foo.png
}

func ExampleTrimWithSourceUnknown() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.TrimWithSource("foo-bar").Image("foo.png"))
	// Output: http://example.com/unsafe/trim/foo.png
}

func ExampleCrop() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.Crop(1, 2, 3, 4).Image("foo.png"))
	// Output: http://example.com/unsafe/1x2:3x4/foo.png
}

func ExampleFullFitIn() {
	b := NewBuilder("http://example.com", "")
	fmt.Println(b.FullFitIn(10, 20).Image("foo.png"))
	// Output: http://example.com/unsafe/full-fit-in/10x20/foo.png
}

func ExampleFitIn() {
	b := NewBuilder("http://example.com", "")
	b.FullFitIn(10, 20)
	b.FitIn(125, 125)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/fit-in/125x125/foo.png
}

func ExampleResize() {
	b := NewBuilder("http://example.com", "")
	b.Resize(100, 100)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/100x100/foo.png
}

func ExampleResizeProportionalWidth() {
	b := NewBuilder("http://example.com", "")
	b.Resize(0, 200)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/0x200/foo.png
}

func ExampleResizeProportionalHeight() {
	b := NewBuilder("http://example.com", "")
	b.Resize(150, 0)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/150x0/foo.png
}

func ExampleResizeUseOrig() {
	b := NewBuilder("http://example.com", "")
	b.Resize(-1, -1)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/origxorig/foo.png
}

func ExampleHAlignLeft() {
	b := NewBuilder("http://example.com", "")
	b.HAlign("left")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/left/foo.png
}

func ExampleHAlignCenter() {
	b := NewBuilder("http://example.com", "")
	b.HAlign("center")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/center/foo.png
}

func ExampleHAlignRight() {
	b := NewBuilder("http://example.com", "")
	b.HAlign("right")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/right/foo.png
}

func ExampleHAlignIgnoreUnknown() {
	b := NewBuilder("http://example.com", "")
	b.HAlign("test")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/foo.png
}

func ExampleVAlignTop() {
	b := NewBuilder("http://example.com", "")
	b.VAlign("top")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/top/foo.png
}

func ExampleVAlignMiddle() {
	b := NewBuilder("http://example.com", "")
	b.VAlign("middle")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/middle/foo.png
}

func ExampleVAlignBottom() {
	b := NewBuilder("http://example.com", "")
	b.VAlign("bottom")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/bottom/foo.png
}

func ExampleVAlignIgnoreUnknown() {
	b := NewBuilder("http://example.com", "")
	b.VAlign("test")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/foo.png
}

func ExampleSmartCropOn() {
	b := NewBuilder("http://example.com", "")
	b.Resize(75, 75)
	b.SmartCrop(true)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/75x75/smart/foo.png
}

func ExampleSmartCropOff() {
	b := NewBuilder("http://example.com", "")
	b.Trim()
	b.SmartCrop(true)
	b.SmartCrop(false)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/trim/foo.png
}

func ExampleAddFilter() {
	b := NewBuilder("http://example.com", "")
	b.Trim()
	b.AddFilter("brightness(50)")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/trim/filters:brightness(50)/foo.png
}

func ExampleAddMultipleFilters() {
	b := NewBuilder("http://example.com", "")
	b.AddFilter("brightness(50)")
	b.AddFilter("blur(20, 20)")
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/filters:brightness(50):blur(20, 20)/foo.png
}

func ExampleMetaOn() {
	b := NewBuilder("http://example.com", "")
	b.Resize(75, 75)
	b.MetaOnly(true)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/meta/75x75/foo.png
}

func ExampleMetaOff() {
	b := NewBuilder("http://example.com", "")
	b.Trim()
	b.MetaOnly(true)
	b.MetaOnly(false)
	fmt.Println(b.Image("foo.png"))
	// Output: http://example.com/unsafe/trim/foo.png
}

func ExampleChain() {
	b := NewBuilder("http://example.com", "")
	b.Crop(10, 20, 30, 40).HAlign("left").VAlign("middle").SmartCrop(true).MetaOnly(true)
	b.AddFilter("brightness(42").AddFilter("foo(bar)")

	fmt.Println(b.Image("foo.png"))
	fmt.Println(b.Image("foo.png").Trim())
	fmt.Println(b.TrimWithSource("top-left").Image("foo.png"))
	fmt.Println(b.FullFitIn(10, 20).Image("foo.png"))
	fmt.Println(b.FitIn(30, 40).Image("foo.png"))
	fmt.Println(b.Resize(200, -1).Image("foo.png"))

	// Output:
	// http://example.com/unsafe/meta/10x20:30x40/left/middle/smart/filters:brightness(42:foo(bar)/foo.png
	// http://example.com/unsafe/meta/trim/10x20:30x40/left/middle/smart/filters:brightness(42:foo(bar)/foo.png
	// http://example.com/unsafe/meta/trim:top-left/10x20:30x40/left/middle/smart/filters:brightness(42:foo(bar)/foo.png
	// http://example.com/unsafe/meta/trim:top-left/10x20:30x40/full-fit-in/10x20/left/middle/smart/filters:brightness(42:foo(bar)/foo.png
	// http://example.com/unsafe/meta/trim:top-left/10x20:30x40/fit-in/30x40/left/middle/smart/filters:brightness(42:foo(bar)/foo.png
	// http://example.com/unsafe/meta/trim:top-left/10x20:30x40/200xorig/left/middle/smart/filters:brightness(42:foo(bar)/foo.png
}
