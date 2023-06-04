package demotivator

import (
	"os"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"gopkg.in/gographics/imagick.v3/imagick"
)

var (
	log = logrus.New()
)

type demotivator struct {
	inFontPath string

	inImage        *imagick.MagickWand
	outImage       *imagick.MagickWand
	topTemplate    *imagick.MagickWand
	bottomTemplate *imagick.MagickWand

	margin      uint
	padding     uint
	borderWidth uint
}

func New(imagePath string, fontPath string) *demotivator {
	inImage := imagick.NewMagickWand()
	err := inImage.ReadImage(imagePath)
	if err != nil {
		log.Error(err)
	}
	return &demotivator{
		inFontPath:     fontPath,
		inImage:        inImage,
		outImage:       imagick.NewMagickWand(),
		topTemplate:    imagick.NewMagickWand(),
		bottomTemplate: imagick.NewMagickWand(),
	}
}

func (d *demotivator) getInImageWidth() uint {
	return d.inImage.GetImageWidth()
}

func (d *demotivator) getInImageHeight() uint {
	return d.inImage.GetImageHeight()
}

func (d *demotivator) confFrameSizes() {
	d.margin = uint(float64(d.getInImageWidth()+d.getInImageHeight()) * float64(0.01))
	d.borderWidth = uint(float64(d.getInImageWidth()+d.getInImageHeight()) * float64(0.004))
	d.padding = uint(float64(d.getInImageWidth()+d.getInImageHeight()) * float64(0.002))
}

func (d *demotivator) createTopTemplate() {
	pw := imagick.NewPixelWand()
	pw.SetColor("black")
	err := d.topTemplate.NewImage(
		uint(d.getInImageWidth()+((d.margin+d.borderWidth+d.padding)*2)),
		uint(d.getInImageHeight()+((d.margin+d.borderWidth+d.padding)*2)),
		pw,
	)
	if err != nil {
		log.Error(err)
	}
	pw.Destroy()

	pw = imagick.NewPixelWand()
	pw.SetColor("white")
	dw := imagick.NewDrawingWand()
	dw.SetFillColor(pw)
	dw.SetStrokeWidth(float64(d.borderWidth))
	dw.Rectangle(
		float64(d.margin), float64(d.margin),
		float64(d.margin+((d.borderWidth+d.padding)*2)+d.getInImageWidth()),
		float64(d.margin+((d.borderWidth+d.padding)*2)+d.getInImageHeight()),
	)
	pw.SetColor("black")
	dw.SetFillColor(pw)
	dw.Rectangle(
		float64(d.margin+d.borderWidth), float64(d.margin+d.borderWidth),
		float64(d.margin+d.borderWidth+(d.padding*2)+d.getInImageWidth()),
		float64(d.margin+d.borderWidth+(d.padding*2)+d.getInImageHeight()),
	)
	pw.Destroy()
	err = d.topTemplate.DrawImage(dw)
	if err != nil {
		log.Error(err)
	}
}

func (d *demotivator) mergeInImageToTopTemplate() {
	err := d.topTemplate.CompositeImage(
		d.inImage, imagick.COMPOSITE_OP_OVER, true,
		int(d.margin+d.borderWidth+d.padding),
		int(d.margin+d.borderWidth+d.padding),
	)
	if err != nil {
		log.Error(err)
	}
}

func (d *demotivator) createBottomTemplate(text1, text2 string) {
	fontSize := 100.0

	pw := imagick.NewPixelWand()
	pw.SetColor("black")
	err := d.bottomTemplate.NewImage(d.topTemplate.GetImageWidth(), d.getInImageHeight(), pw)
	if err != nil {
		log.Error(err)
	}
	pw.Destroy()

	pw = imagick.NewPixelWand()
	pw.SetColor("white")
	dw := imagick.NewDrawingWand()
	err = dw.SetFont(d.inFontPath)
	if err != nil {
		log.Error(err)
	}
	dw.SetFontSize(fontSize)
	dw.SetFillColor(pw)
	pw.Destroy()

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			log.Errorf("%v, %s", panicInfo, string(debug.Stack()))
		}
	}()
	// error on set emoji font : QueryFontMetrics
	metrics1 := d.bottomTemplate.QueryFontMetrics(dw, text1)
	metrics2 := d.bottomTemplate.QueryFontMetrics(dw, text2)

	for metrics1.TextWidth > float64(d.topTemplate.GetImageWidth()) ||
		metrics2.TextWidth > float64(d.topTemplate.GetImageWidth()) {
		fontSize -= 6
		dw.SetFontSize(fontSize)
		metrics1 = d.bottomTemplate.QueryFontMetrics(dw, text1)
		metrics2 = d.bottomTemplate.QueryFontMetrics(dw, text2)
	}

	if len(text1) == 0 {
		metrics1.TextHeight = 0
	}
	if len(text2) == 0 {
		metrics2.TextHeight = 0
	}

	err = d.bottomTemplate.ResizeImage(
		d.topTemplate.GetImageWidth(),
		uint(d.margin*2)+uint(metrics1.TextHeight)+uint(metrics2.TextHeight),
		imagick.FILTER_LANCZOS2,
	)
	if err != nil {
		log.Error(err)
	}

	dw.Annotation(
		float64(d.bottomTemplate.GetImageWidth())/2-metrics1.TextWidth/2,
		metrics1.TextHeight-float64(d.margin/2),
		text1,
	)
	dw.Annotation(
		float64(d.bottomTemplate.GetImageWidth())/2-metrics2.TextWidth/2,
		metrics1.TextHeight-float64(d.margin/2)+metrics2.TextHeight,
		text2,
	)
	err = d.bottomTemplate.DrawImage(dw)
	if err != nil {
		log.Error(err)
	}
	dw.Destroy()
}

func (d *demotivator) mergeTopAndBottomTemplates() {
	pw := imagick.NewPixelWand()
	pw.SetColor("black")
	d.outImage.NewImage(
		d.topTemplate.GetImageWidth(),
		d.topTemplate.GetImageHeight()+d.bottomTemplate.GetImageHeight(),
		pw,
	)
	pw.Destroy()

	d.outImage.CompositeImage(d.topTemplate, imagick.COMPOSITE_OP_OVER, true, 0, 0)
	d.outImage.CompositeImage(d.bottomTemplate, imagick.COMPOSITE_OP_OVER, true, 0, int(d.topTemplate.GetImageHeight()))
}

func (d *demotivator) testShow(w *imagick.MagickWand) {
	err := w.ResizeImage(w.GetImageWidth()/2, w.GetImageHeight()/2, imagick.FILTER_LANCZOS2)
	if err != nil {
		log.Error(err)
	}
	err = w.DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		log.Error(err)
	}
}

func (d *demotivator) saveImage(imageOutPath string) {
	d.outImage.WriteImage(imageOutPath)
}

func (d *demotivator) Generate(text1, text2, imageOutPath string, debug bool) {
	d.confFrameSizes()
	d.createTopTemplate()
	d.mergeInImageToTopTemplate()
	d.createBottomTemplate(text1, text2)
	d.mergeTopAndBottomTemplates()
	if debug {
		d.testShow(d.outImage)
		return
	}
	d.saveImage(imageOutPath)
}
