package demotivator

import (
	"bytes"
	"image"
)

func New(srcImage image.Image, texts []string) *Demotivator {
	_tmp := &Demotivator{
		TemplateConfig: &Template{
			FontConfig: &Font{},
			TextConfig: &Text{},
			FrameConfig: &Frame{},
		},
		SrcImageConfig: &SrcImage{},
	}
	_tmp.SetConfig(srcImage, texts)
	return _tmp
}

func (dem *Demotivator) SetConfig(srcImage image.Image, texts []string) {
	dem.TemplateConfig.FrameConfig.Padding = 4
	dem.TemplateConfig.FrameConfig.Width = 4
	if srcImage != nil {
		dem.SrcImageConfig.Image = srcImage
		dem.SrcImageConfig.Width = srcImage.Bounds().Size().X
		dem.SrcImageConfig.Height = srcImage.Bounds().Size().Y
		dem.TemplateConfig.PaddingTop = srcImage.Bounds().Size().Y / 35
		dem.TemplateConfig.PaddingBottom = (srcImage.Bounds().Size().Y / 35) * 6
		dem.TemplateConfig.PaddingLeft = srcImage.Bounds().Size().X / 35
		dem.TemplateConfig.PaddingRight = srcImage.Bounds().Size().X / 35
		dem.TemplateConfig.TextConfig.VerticalSpacing = dem.TemplateConfig.PaddingTop / 4
		dem.TemplateConfig.Width = dem.TemplateConfig.PaddingLeft + dem.TemplateConfig.PaddingRight +
			(dem.TemplateConfig.FrameConfig.Width * 2) + (dem.TemplateConfig.FrameConfig.Padding * 2) +
			srcImage.Bounds().Size().X
		dem.TemplateConfig.Height = dem.TemplateConfig.PaddingTop + dem.TemplateConfig.PaddingBottom +
			(dem.TemplateConfig.FrameConfig.Width * 2) + (dem.TemplateConfig.FrameConfig.Padding * 2) +
			srcImage.Bounds().Size().Y + (dem.TemplateConfig.TextConfig.VerticalSpacing * 2)
	}
	if texts != nil { dem.TemplateConfig.TextConfig.Texts = texts }

	if srcImage != nil && texts != nil { dem.configsConfigured = true }
}

func (dem *Demotivator) GetImageBuffer() (imageBuffer *bytes.Buffer, err error) {
	imageBuffer = &bytes.Buffer{}
	if err = dem.TemplateConfig.Image.EncodePNG(imageBuffer); err != nil { return }
	return
}

func (dem *Demotivator) Make(srcImage image.Image, texts []string) (imageBuffer *bytes.Buffer, err error) {
	if !dem.configsConfigured { dem.SetConfig(srcImage, texts) }
	dem.TemplateConfig.RenderTemplate()
	dem.RenderSrcImage()
	if err = dem.TemplateConfig.RenderTexts(); err != nil { return }

	imageBuffer, err = dem.GetImageBuffer()
	if err != nil { return }
	return
}
