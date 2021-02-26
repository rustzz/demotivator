package demotivator

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func (template *Template) configureFont(text string) (err error) {
	template.FontConfig.Font, err = truetype.Parse(goregular.TTF)
	if err != nil { return }
	template.TextConfig.reachedImageBorder = false
	template.FontConfig.FontSize = 10
	for ;; {
		fontFace := truetype.NewFace(
			template.FontConfig.Font, &truetype.Options{Size: float64(template.FontConfig.FontSize)},
		)
		template.Image.SetFontFace(fontFace)

		widthText, heightText := template.Image.MeasureString(text)
		if int(heightText) > (template.PaddingBottom / 2) { template.FontConfig.FontSize -= 1 }
		if int(heightText) < (template.PaddingBottom / 2) { template.FontConfig.FontSize += 1 }
		if ((template.PaddingBottom / 2) - int(heightText)) < 5 &&
			((template.PaddingBottom / 2) - int(heightText)) > -5 {
			for ; template.Image.Width() < int(widthText); {
				template.TextConfig.reachedImageBorder = true
				template.FontConfig.FontSize -= 1
				fontFace = truetype.NewFace(
					template.FontConfig.Font, &truetype.Options{Size: float64(template.FontConfig.FontSize)},
				)
				template.Image.SetFontFace(fontFace)
				widthText, heightText = template.Image.MeasureString(text)
			}
			return
		}
	}
}

func (template *Template) RenderTexts() (err error) {
	template.Image.SetHexColor("#ffffff")

	if err = template.configureFont(template.TextConfig.Texts[0]); err != nil { return }
	fontSizeUpper := template.FontConfig.FontSize
	if fontSizeUpper < 10 { fontSizeUpper = 0 }

	if err = template.configureFont(template.TextConfig.Texts[1]); err != nil { return }
	fontSizeLower := template.FontConfig.FontSize
	if !template.TextConfig.reachedImageBorder { fontSizeLower -= 35 }
	if fontSizeLower < 10 { fontSizeLower = 0 }

	fontFaceUpper := truetype.NewFace(template.FontConfig.Font, &truetype.Options{Size: float64(fontSizeUpper)})
	fontFaceLower := truetype.NewFace(template.FontConfig.Font, &truetype.Options{Size: float64(fontSizeLower)})

	template.Image.SetFontFace(fontFaceUpper)
	widthUpperText, heightUpperText := template.Image.MeasureString(template.TextConfig.Texts[0])
	template.Image.DrawString(
		template.TextConfig.Texts[0],
		float64((template.Image.Width() / 2 ) - int(widthUpperText / 2)),
		float64(template.Image.Height() - (template.PaddingBottom + template.TextConfig.VerticalSpacing) +
			int(heightUpperText)),
	)

	template.Image.SetFontFace(fontFaceLower)
	widthLowerText, heightLowerText := template.Image.MeasureString(template.TextConfig.Texts[1])
	template.Image.DrawString(
		template.TextConfig.Texts[1],
		float64((template.Image.Width() / 2) - int(widthLowerText / 2)),
		float64(template.Image.Height() - (template.PaddingBottom) + template.TextConfig.VerticalSpacing +
			int(heightUpperText) + int(heightLowerText)),
	)
	return
}
