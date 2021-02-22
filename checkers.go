package demotivator

import (
	"github.com/rustzz/demotivator/errors"
	"image"
)


func CheckSrcImage(srcImage image.Image) error {
	if srcImage.Bounds().Size().X < 150 || srcImage.Bounds().Size().Y < 150 {
		return &errors.ImageSizeError{}
	}
	return nil
}
