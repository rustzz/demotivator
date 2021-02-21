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
