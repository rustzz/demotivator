package demotivator

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func (dem *Demotivator) settingFont(outImage *gg.Context, text string) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		font, _ := truetype.Parse(goregular.TTF)
		face := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
		outImage.SetFontFace(face)

		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > (dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) {
			fontSize -= 1
		}
		if int(heightText) < (dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) {
			fontSize += 1
		}
		if ((dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) - int(heightText)) < 5 &&
			((dem.TemplateConfig.PaddingBottom / 2 - dem.TextConfig.TextSpacing) - int(heightText)) > -5 {
			for ; outImage.Width() < int(widthText); {
				fontSize -= 1
				face = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
				outImage.SetFontFace(face)
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (dem *Demotivator) setTexts(outImage *gg.Context, texts []string) (*gg.Context, error) {
	font, _ := truetype.Parse(goregular.TTF)
	outImage.SetHexColor("#ffffff")

	fontSize, err := dem.settingFont(outImage, texts[0])
	if err != nil {
		return outImage, err
	}
	if fontSize < 10 {
		fontSize = 0
	}

	face := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
	outImage.SetFontFace(face)

	widthUpperText, _ := outImage.MeasureString(texts[0])
	outImage.DrawString(
		texts[0],
		float64((outImage.Width() / 2 ) - int(widthUpperText / 2)),
		float64(outImage.Height() - ((dem.TemplateConfig.PaddingBottom / 2) + dem.TextConfig.TextSpacing)),
	)

	fontSize, err = dem.settingFont(outImage, texts[1])
	if err != nil {
		return outImage, err
	}
	fontSize -= 30
	if fontSize < 10 {
		fontSize = 0
	}

	face = truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
	outImage.SetFontFace(face)

	widthLowerText, _ := outImage.MeasureString(texts[1])
	outImage.DrawString(
		texts[1],
		float64((outImage.Width() / 2) - int(widthLowerText / 2)),
		float64(outImage.Height() -
			dem.TemplateConfig.PaddingBottom/2+ int(float64(dem.TextConfig.TextSpacing) * 1.5)),
	)
	return outImage, nil
}
