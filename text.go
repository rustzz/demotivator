package demotivator

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	font, _ = truetype.Parse(goregular.TTF)
)

func (dem *Demotivator) settingFont(outImage *gg.Context, text string) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		fontFace := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
		outImage.SetFontFace(fontFace)

		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > (dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) { fontSize -= 1 }
		if int(heightText) < (dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) { fontSize += 1 }
		if ((dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) - int(heightText)) < 5 &&
			((dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) - int(heightText)) > -5 {
			for ; outImage.Width() < int(widthText); {
				fontSize -= 1
				fontFace = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
				outImage.SetFontFace(fontFace)
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (dem *Demotivator) setTexts(outImage *gg.Context, texts []string) (*gg.Context, error) {
	outImage.SetHexColor("#ffffff")

	fontSizeUpper, err := dem.settingFont(outImage, texts[0])
	if err != nil { return outImage, err }
	if fontSizeUpper < 10 { fontSizeUpper = 0 }

	fontSizeLower, err := dem.settingFont(outImage, texts[1])
	if err != nil { return outImage, err }
	fontSizeLower -= 25
	if fontSizeLower < 10 { fontSizeLower = 0 }

	fontFaceUpper := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeUpper)})
	outImage.SetFontFace(fontFaceUpper)
	fontFaceLower := truetype.NewFace(font, &truetype.Options{Size: float64(fontSizeLower)})
	outImage.SetFontFace(fontFaceLower)

	widthUpperText, _ := outImage.MeasureString(texts[0])
	outImage.DrawString(
		texts[0],
		float64((outImage.Width() / 2 ) - int(widthUpperText / 2)),
		float64(outImage.Height() - ((dem.TemplateConfig.PaddingBottom / 2) + dem.TextConfig.TextSpacing)),
	)

	widthLowerText, _ := outImage.MeasureString(texts[1])
	outImage.DrawString(
		texts[1],
		float64((outImage.Width() / 2) - int(widthLowerText / 2)),
		float64(outImage.Height() - dem.TemplateConfig.PaddingBottom/2 +
			int(float64(dem.TextConfig.TextSpacing) * 1.5)),
	)
	return outImage, nil
}
