package main

import (
	"github.com/rustzz/demotivator"
	"image"
	"log"
)


func main() {
	url := "https://sun7-8.userapi.com/impg/HzygGAv_y7IAOGKS2jA5fiS7zZvkyUEx3IwC3A/J0RpHYK6xaY.jpg" +
		"?size=640x640&quality=96&proxy=1&sign=6a1277be354b123381867f534b0cf73c&type=album"
	imageReader := demotivator.LoadSrcImageFromURL(url)
	im, _, err := image.Decode(&imageReader)
	if err != nil {
		log.Fatal(err)
		return
	}

	dem := demotivator.New()
	dem.Make(im, []string{
		"cum", "cum",
	}, "/home/rustzz/out.png")
}
