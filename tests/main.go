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
	texts = []string{"cum CUM", "cum CUM"}
)

func GetImage() (outImage image.Image, err error) {
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

func FromConstructor() (imageBuffer *bytes.Buffer, err error) {
	srcImage, err := GetImage()
	if err != nil { return }
	dem := demotivator.New(srcImage, texts)
	imageBuffer, err = dem.Make(nil, nil)
	if err != nil { return }
	return
}

func FromObject() (imageBuffer *bytes.Buffer, err error) {
	srcImage, err := GetImage()
	if err != nil { return }
	dem := &demotivator.Demotivator{
		TemplateConfig: &demotivator.Template{
			FontConfig: &demotivator.Font{},
			TextConfig: &demotivator.Text{},
			FrameConfig: &demotivator.Frame{},
		},
		SrcImageConfig: &demotivator.SrcImage{},
	}
	imageBuffer, err = dem.Make(srcImage, texts)
	if err != nil { return }
	return
}

func main() {
	//imageBuffer, err := FromConstructor()
	imageBuffer, err := FromObject()
	if err != nil { log.Fatal(err) }

	homeDir, err := os.UserHomeDir()
	file, err := os.Create(fmt.Sprintf("%s/out1.png", homeDir))
	if err != nil { log.Fatal(err) }
	defer file.Close()

	im, _, err := image.Decode(imageBuffer)
	if err != nil { log.Fatal(err) }
	if err = png.Encode(file, im); err != nil { log.Fatal(err) }
}
