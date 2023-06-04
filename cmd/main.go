package main

import (
	"demotivator/internal/demotivator"
	"flag"
)

func main() {
	var imageInPath, imageOutPath, fontPath, text1, text2 string
	var debug bool
	flag.StringVar(&imageInPath, "i", "image.jpg", "input image path")
	flag.StringVar(&imageOutPath, "o", "image-out.jpg", "output image path")
	flag.StringVar(&fontPath, "font", "", "font path")
	flag.StringVar(&text1, "text1", "demotivator", "first line")
	flag.StringVar(&text2, "text2", "demotivator", "second line")
	flag.BoolVar(&debug, "debug", false, "show image without save")
	flag.Parse()

	dem := demotivator.New(imageInPath, fontPath)
	dem.Generate(text1, text2, imageOutPath, debug)
}
