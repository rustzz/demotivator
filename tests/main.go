package main

import (
	"bytes"
	"fmt"
	"github.com/rustzz/demotivator"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	url = "https://2ch.hk/makaba/templates/img/anon.jpg"
	texts = [2]string{"cum CUM", "cum CUM"}
)

func GetImage(url string) (outImage image.Image, err error) {
	resp, err := http.Get(url)
	if err != nil { return }
	defer resp.Body.Close()

	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil { return }

	imageBuffer := bytes.NewBuffer(imageBytes)
	outImage, _, err = image.Decode(imageBuffer)
	if err != nil { return }
	return
}

func FromConstructor() (imgBytes []byte, err error) {
	srcImage, err := GetImage(url)
	if err != nil { return }
	dem := demotivator.New(srcImage, texts)
	imgBytes, err = dem.Make(nil, [2]string{})
	if err != nil { return }
	return
}

func FromObject() (imgBytes []byte, err error) {
	srcImage, err := GetImage(url)
	if err != nil { return }
	dem := &demotivator.Demotivator{
		TemplateConfig: &demotivator.Template{
			TextConfig: &demotivator.Text{ FontConfig: &demotivator.Font{} },
			FrameConfig: &demotivator.Frame{},
		},
		SrcImageConfig: &demotivator.SrcImage{},
	}
	imgBytes, err = dem.Make(srcImage, texts)
	if err != nil { return }
	return
}

func main() {
	imgBytes, err := FromConstructor()
	//imgBytes, err := FromObject()
	if err != nil { log.Fatal(err) }

	homeDir, err := os.UserHomeDir()
	file, err := os.Create(fmt.Sprintf("%s/out.png", homeDir))
	if err != nil { log.Fatal(err) }
	defer file.Close()

	imgBuffer := bytes.NewBuffer(imgBytes)
	im, _, err := image.Decode(imgBuffer)
	if err != nil { log.Fatal(err) }
	if err = png.Encode(file, im); err != nil { log.Fatal(err) }
}
