package demotivator

import (
	"bytes"
	"github.com/fogleman/gg"
	"image"
)

type TemplateConfig struct {
	PaddingTop		int
	PaddingBottom	int
	PaddingLeft		int
	PaddingRight	int
}

type FrameConfig struct {
	MarginBottom	int
	Padding			int
	FrameWidth		int
}

type TextConfig struct {
	TextSpacing		int
}

type Demotivator struct {
	OutImage			*gg.Context
	TemplateConfig		TemplateConfig
	FrameConfig			FrameConfig
	TextConfig			TextConfig
}

func (dem *Demotivator) setConfigs(srcImage image.Image) {
	dem.TemplateConfig.PaddingTop = srcImage.Bounds().Size().Y/35
	dem.TemplateConfig.PaddingBottom = (srcImage.Bounds().Size().Y/20)*3
	dem.TemplateConfig.PaddingLeft = srcImage.Bounds().Size().X/35
	dem.TemplateConfig.PaddingRight = srcImage.Bounds().Size().X/35
	dem.FrameConfig.Padding = 2
	dem.FrameConfig.FrameWidth = 2
	dem.TextConfig.TextSpacing = dem.TemplateConfig.PaddingBottom / 6
	return
}

func saveImage(outImage *gg.Context, path string) (imageReader *bytes.Reader, err error) {
	if len(path) != 0 {
		err = outImage.SavePNG(path)
		if err != nil {
			return
		}
	} else {
		imageBuffer := new(bytes.Buffer)
		err = outImage.EncodePNG(imageBuffer)
		if err != nil {
			return
		}
		imageReader = bytes.NewReader(imageBuffer.Bytes())
		return
	}
	return
}

func (dem *Demotivator) Make(srcImage image.Image, texts []string, outPath string) (imageReader *bytes.Reader, err error) {
	if !CheckSrcImage(srcImage) {
		return
	}
	dem.setConfigs(srcImage)
	outImage := dem.createTemplate(srcImage)
	outImage = dem.placeSrcImage(outImage, srcImage)
	outImage, err = dem.setTexts(outImage, texts)
	if err != nil {
		return
	}

	dem.OutImage = outImage
	imageReader, err = saveImage(outImage, outPath)
	if err != nil {
		return
	}
	return
}
