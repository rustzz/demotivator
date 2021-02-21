package demotivator

import (
	"github.com/fogleman/gg"
	"image"
)

func (dem *Demotivator) makeTemplateImage(srcImage image.Image) (templateImage *gg.Context) {
	templateWidth := srcImage.Bounds().Size().X+
		dem.TemplateConfig.PaddingLeft+
		dem.TemplateConfig.PaddingRight+
		(dem.FrameConfig.FrameWidth*2)+
		(dem.FrameConfig.Padding*2)
	templateHeight := srcImage.Bounds().Size().Y+
		dem.TemplateConfig.PaddingTop+
		dem.TemplateConfig.PaddingBottom+
		dem.TextConfig.TextSpacing+
		(dem.FrameConfig.FrameWidth*2)+
		(dem.FrameConfig.Padding*2)
	templateImage = gg.NewContext(
		templateWidth,
		templateHeight,
	)
	return
}

func (dem *Demotivator) fillBackground(template *gg.Context) *gg.Context {
	template.SetHexColor("#000000")
	template.DrawRectangle(
		0, 0,
		float64(template.Width()), float64(template.Height()),
	)
	template.Fill()
	return template
}

func (dem *Demotivator) drawFrame(template *gg.Context) *gg.Context {
	template.SetHexColor("#ffffff")
	template.DrawRectangle(
		float64(dem.TemplateConfig.PaddingLeft),
		float64(dem.TemplateConfig.PaddingTop),
		float64(template.Width()-(dem.TemplateConfig.PaddingLeft+dem.TemplateConfig.PaddingRight)),
		float64(template.Height()-(dem.TemplateConfig.PaddingTop+dem.TemplateConfig.PaddingBottom+dem.TextConfig.TextSpacing)),
	)
	template.Fill()
	template.SetHexColor("#000000")
	template.DrawRectangle(
		float64(dem.TemplateConfig.PaddingLeft+dem.FrameConfig.FrameWidth),
		float64(dem.TemplateConfig.PaddingTop+dem.FrameConfig.FrameWidth),
		float64(template.Width()-(dem.TemplateConfig.PaddingLeft+dem.TemplateConfig.PaddingRight+(dem.FrameConfig.FrameWidth*2))),
		float64(
			template.Height()-
				(dem.TemplateConfig.PaddingTop+dem.TemplateConfig.PaddingBottom+(dem.FrameConfig.FrameWidth*2)+dem.TextConfig.TextSpacing),
		),
	)
	template.Fill()
	return template
}

func (dem *Demotivator) placeSrcImage(outImage *gg.Context, srcImage image.Image) *gg.Context {
	outImage.DrawImage(
		srcImage,
		dem.TemplateConfig.PaddingLeft+dem.FrameConfig.FrameWidth+dem.FrameConfig.Padding,
		dem.TemplateConfig.PaddingTop+dem.FrameConfig.FrameWidth+dem.FrameConfig.Padding,
	)
	return outImage
}

func (dem *Demotivator) createTemplate(srcImage image.Image) (template *gg.Context) {
	template = dem.makeTemplateImage(srcImage)
	template = dem.fillBackground(template)
	template = dem.drawFrame(template)
	return
}
