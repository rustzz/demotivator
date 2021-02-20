package demotivator

import "image"


func (dem *Demotivator) CheckSrcImage(srcImage image.Image) bool {
	if srcImage.Bounds().Size().X < 150 || srcImage.Bounds().Size().Y < 150 {
		return false
	}
	return true
}
