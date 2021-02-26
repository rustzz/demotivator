package demotivator

import (
	"github.com/fogleman/gg"
)

func (template *Template) makeTemplateImage() {
	template.Image = gg.NewContext(template.Width, template.Height)
}

func (template *Template) fillBackground() {
	template.Image.SetHexColor("#000000")

	template.Image.DrawRectangle(
		0, 0,
		float64(template.Width), float64(template.Height),
	)
	template.Image.Fill()
}

func (template *Template) drawFrame() {
	template.Image.SetHexColor("#ffffff")
	template.Image.DrawRectangle(
		float64(template.PaddingLeft),
		float64(template.PaddingTop),
		float64(template.Width - (template.PaddingLeft + template.PaddingRight)),
		float64(template.Height -
			(template.PaddingTop + template.PaddingBottom +
				(template.TextConfig.VerticalSpacing * 2)),
		),
	)
	template.Image.Fill()

	template.Image.SetHexColor("#000000")
	template.Image.DrawRectangle(
		float64(template.PaddingLeft + template.FrameConfig.Width),
		float64(template.PaddingTop + template.FrameConfig.Width),
		float64(template.Width -
			(template.PaddingLeft + template.PaddingRight + template.FrameConfig.Width + template.FrameConfig.Padding),
		),
		float64(template.Height -
			(template.PaddingTop + template.PaddingBottom + template.FrameConfig.Width + template.FrameConfig.Padding +
				(template.TextConfig.VerticalSpacing * 2)),
		),
	)
	template.Image.Fill()
}

func (template *Template) RenderTemplate() {
	template.makeTemplateImage()
	template.fillBackground()
	template.drawFrame()
}

func (dem *Demotivator) RenderSrcImage() {
	dem.TemplateConfig.Image.DrawImage(
		dem.SrcImageConfig.Image,
		dem.TemplateConfig.PaddingLeft + dem.TemplateConfig.FrameConfig.Padding + dem.TemplateConfig.FrameConfig.Width,
		dem.TemplateConfig.PaddingTop + dem.TemplateConfig.FrameConfig.Padding + dem.TemplateConfig.FrameConfig.Width,
	)
	return
}
