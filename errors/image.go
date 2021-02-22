package errors

func (err *ImageSizeError) Error() string {
	return "Изображение слишком мелкое"
}
