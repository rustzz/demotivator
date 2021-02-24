package main

import (
	"fmt"
	"github.com/rustzz/demotivator"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	url := "https://2ch.hk/makaba/templates/img/anon.jpg"
	imageReader, err := demotivator.LoadSrcImageFromURL(url)
	if err != nil {
		log.Fatal(err)
	}
	im, _, err := image.Decode(imageReader)
	if err != nil {
		log.Fatal(err)
	}

	homeDir, _ := os.UserHomeDir()
	dem := demotivator.New()
	if _, err = dem.Make(&im, []string{
		"cum", "CUM",
	}, fmt.Sprintf("%s/out.png", homeDir)); err != nil {
		log.Fatal(err)
	}
}
