package main

import (
	"fmt"
	"github.com/rustzz/demotivator"
	"image"
	"log"
	"os"
)


func main() {
	url := "https://2ch.hk/makaba/templates/img/anon.jpg"
	url = "https://sun7-6.userapi.com/impg/1wSphkz44J2c4Sufvl120SXr9qXRb4A7utSzRQ/OiGXoDsp7QQ.jpg?size=318x318&quality=96&proxy=1&sign=0250805d18276a51ddbccf677336da8d&type=album"
	imageReader, err := demotivator.LoadSrcImageFromURL(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	im, _, err := image.Decode(imageReader)
	if err != nil {
		log.Fatal(err)
		return
	}

	homeDir, _ := os.UserHomeDir()
	dem := &demotivator.Demotivator{}
	if _, err = dem.Make(im, []string{
		"cum", "cum",
	}, fmt.Sprintf("%s/out.png", homeDir)); err != nil {
		log.Fatal(err)
	}
}
