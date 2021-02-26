package demotivator

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
)

type Font struct {
	Font		*truetype.Font
	FontSize	int
}

type Text struct {
	Texts				[]string
	VerticalSpacing		int

	reachedImageBorder	bool
}

type Frame struct {
	Width			int
	Padding			int
}

type SrcImage struct {
	Image	image.Image
	Width	int
	Height	int
}

type Template struct {
	FontConfig		*Font
	TextConfig		*Text
	FrameConfig		*Frame
	Image			*gg.Context

	Width			int
	Height			int
	PaddingTop		int
	PaddingLeft		int
	PaddingRight	int
	PaddingBottom	int
}

type Demotivator struct {
	TemplateConfig		*Template
	SrcImageConfig		*SrcImage

	configsConfigured	bool
}
