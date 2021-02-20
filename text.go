package demotivator

import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"path/filepath"
	"runtime"
)

// get root path of project
var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)
// ========================

func (dem *Demotivator) settingFont(outImage *gg.Context, text string) int {
	fontSize := 10
	for ;; {
		err := outImage.LoadFontFace(fmt.Sprintf("%s/fonts/arial.ttf", basePath), float64(fontSize))
		if err != nil {
			log.Fatal(err)
			return fontSize
		}
		widthText, heightText := outImage.MeasureString(text)
		if int(heightText) > (dem.TemplateConfig.PaddingBottom/2-dem.TextConfig.TextSpacing) {
			fontSize -= 1
		}
		if int(heightText) < (dem.TemplateConfig.PaddingBottom/2-dem.TextConfig.TextSpacing) {
			fontSize += 1
		}
		if ((dem.TemplateConfig.PaddingBottom/2-dem.TextConfig.TextSpacing)-int(heightText)) < 5 &&
			((dem.TemplateConfig.PaddingBottom/2-dem.TextConfig.TextSpacing)-int(heightText)) > -5 {
			for ; outImage.Width() < int(widthText); {
				fontSize -= 1
				err := outImage.LoadFontFace(fmt.Sprintf("%s/fonts/arial.ttf", basePath), float64(fontSize))
				if err != nil {
					return fontSize
				}
				widthText, heightText = outImage.MeasureString(text)
			}
			return fontSize
		}
	}
	return fontSize
}

func (dem *Demotivator) setTexts(outImage *gg.Context, texts []string) *gg.Context {
	fontSize := dem.settingFont(outImage, texts[0])
	if fontSize < 10 {
		fontSize = 0
	}
	err := outImage.LoadFontFace(fmt.Sprintf("%s/fonts/arial.ttf", basePath), float64(fontSize))
	if err != nil {
		log.Fatal(err)
		return outImage
	}
	widthUpperText, _ := outImage.MeasureString(texts[0])
	outImage.SetHexColor("#ffffff")
	outImage.DrawString(
		texts[0],
		float64((outImage.Width()/2)-int(widthUpperText/2)),
		float64(outImage.Height()-((dem.TemplateConfig.PaddingBottom/2)+dem.TextConfig.TextSpacing)),
	)

	fontSize = dem.settingFont(outImage, texts[1])-15
	if fontSize < 10 {
		fontSize = 0
	}
	err = outImage.LoadFontFace(fmt.Sprintf("%s/fonts/arial.ttf", basePath), float64(fontSize))
	if err != nil {
		log.Fatal(err)
		return outImage
	}
	widthLowerText, _ := outImage.MeasureString(texts[1])
	outImage.SetHexColor("#ffffff")
	outImage.DrawString(
		texts[1],
		float64((outImage.Width()/2)-int(widthLowerText/2)),
		float64(outImage.Height()-
			(dem.TemplateConfig.PaddingBottom/2)+(dem.TextConfig.TextSpacing)),
	)
	return outImage
}
