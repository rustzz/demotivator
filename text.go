package demotivator

import (
	"fmt"
	"github.com/fogleman/gg"
	"path/filepath"
	"runtime"
)

// get root path of project
var (
	_, b, _, _	= runtime.Caller(0)
	basePath	= filepath.Dir(b)
	fontName	= "times.ttf"
)
// ========================

func (dem *Demotivator) settingFont(outImage *gg.Context, text string) (fontSize int, err error) {
	fontSize = 10
	for ;; {
		if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
			err != nil {
				return
		}
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
				if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
					err != nil {
						return
				}
				widthText, heightText = outImage.MeasureString(text)
			}
			return
		}
	}
}

func (dem *Demotivator) setTexts(outImage *gg.Context, texts []string) (*gg.Context, error) {
	outImage.SetHexColor("#ffffff")
	fontSize, err := dem.settingFont(outImage, texts[0])
	if err != nil {
		return outImage, err
	}
	if fontSize < 10 {
		fontSize = 0
	}

	if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize));
		err != nil {
		return outImage, err
	}

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

	if err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/%s", basePath, fontName), float64(fontSize)); err != nil {
		return outImage, err
	}

	widthLowerText, _ := outImage.MeasureString(texts[1])
	outImage.DrawString(
		texts[1],
		float64((outImage.Width() / 2) - int(widthLowerText / 2)),
		float64(outImage.Height() -
			int(dem.TemplateConfig.PaddingBottom / 2) + int(float64(dem.TextConfig.TextSpacing) * 1.5)),
	)
	return outImage, nil
}
