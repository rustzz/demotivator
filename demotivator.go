package demotivator

import (
	"bytes"
	"github.com/fogleman/gg"
	"image"
	"log"
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

func saveImage(outImage *gg.Context, path string) bytes.Buffer {
	if len(path) != 0 {
		err := outImage.SavePNG(path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var imageBuffer bytes.Buffer
		err := outImage.EncodePNG(&imageBuffer)
		if err != nil {
			log.Fatal(err)
		}
		return imageBuffer
	}
	return bytes.Buffer{}
}

func (dem *Demotivator) Make(srcImage image.Image, texts []string, outPath string) bytes.Buffer {
	if !CheckSrcImage(srcImage) {
		return bytes.Buffer{}
	}
	dem.setConfigs(srcImage)
	outImage := dem.createTemplate(srcImage)
	outImage = dem.drawFrame(outImage)
	outImage = dem.placeSrcImage(outImage, srcImage)
	outImage = dem.setTexts(outImage, texts)

	dem.OutImage = outImage
	return saveImage(outImage, outPath)
}
