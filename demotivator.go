package demotivator

type demotivator struct {
	fontPath, textFirstLine, textSecondLine string
	imageBlob                               []byte

	drawer *drawer
}

func New(textFirstLine, textSecondLine, fontPath string, imageBlob []byte) *demotivator {
	return &demotivator{
		textFirstLine:  textFirstLine,
		textSecondLine: textSecondLine,
		fontPath:       fontPath,
		imageBlob:      imageBlob,
		drawer:         newDrawer(),
	}
}

func (d *demotivator) Generate() {
	d.drawer.LoadInImage(d.imageBlob)
	d.drawer.ConfigureFrameSizes()
	d.drawer.CreateTopTemplate()
	d.drawer.MergeInImageToTopTemplate()
	d.drawer.CreateBottomTemplate(d.textFirstLine, d.textSecondLine, d.fontPath)
	d.drawer.MergeTopAndBottomTemplates()
}

func (d *demotivator) GetBlob() []byte {
	return d.drawer.GetBlob()
}

func (d *demotivator) SaveImage(outputPath string) {
	d.drawer.outImage.WriteImage(outputPath)
}
