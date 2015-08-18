package thumbor

import (
	"fmt"
	"strconv"
)

func (b *Builder) Image(img string) *Builder {
	b.ImagePath = img
	return b
}

func (b *Builder) Trim() *Builder {
	b.Commands.Trim = "trim"
	return b
}

func (b *Builder) TrimWithSource(source string) *Builder {
	if source != "top-left" && source != "bottom-right" {
		return b.Trim()
	}

	b.Commands.Trim = fmt.Sprintf("trim:%s", source)
	return b
}

func (b *Builder) Crop(topLeftX, topLeftY, bottomRightX, bottomRightY int) *Builder {
	b.Commands.Crop = fmt.Sprintf(
		"%dx%d:%dx%d",
		topLeftX,
		topLeftY,
		bottomRightX,
		bottomRightY,
	)
	return b
}

func (b *Builder) FullFitIn(width, height int) *Builder {
	b.Commands.Resize = fmt.Sprintf("full-fit-in/%dx%d", width, height)
	return b
}

func (b *Builder) FitIn(width, height int) *Builder {
	b.Commands.Resize = fmt.Sprintf("fit-in/%dx%d", width, height)
	return b
}

func (b *Builder) Resize(width, height int) *Builder {
	widthStr := "orig"
	if width >= 0 {
		widthStr = strconv.Itoa(width)
	}

	heightStr := "orig"
	if height >= 0 {
		heightStr = strconv.Itoa(height)
	}

	b.Commands.Resize = fmt.Sprintf("%sx%s", widthStr, heightStr)
	return b
}

func (b *Builder) HAlign(align string) *Builder {
	if align != "left" && align != "center" && align != "right" {
		return b
	}

	b.Commands.HAlign = align
	return b
}

func (b *Builder) VAlign(align string) *Builder {
	if align != "top" && align != "middle" && align != "bottom" {
		return b
	}

	b.Commands.VAlign = align
	return b
}

func (b *Builder) SmartCrop(crop bool) *Builder {
	b.Commands.SmartCrop = crop
	return b
}

func (b *Builder) AddFilter(filter string) *Builder {
	b.Commands.Filters = append(b.Commands.Filters, filter)
	return b
}

func (b *Builder) MetaOnly(meta bool) *Builder {
	b.Commands.Meta = meta
	return b
}
