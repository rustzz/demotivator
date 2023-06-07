package demotivator

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type drawer struct {
	inImage        *imagick.MagickWand
	outImage       *imagick.MagickWand
	topTemplate    *imagick.MagickWand
	bottomTemplate *imagick.MagickWand

	margin      uint
	padding     uint
	borderWidth uint
}

func newDrawer() *drawer {
	return &drawer{
		inImage:        imagick.NewMagickWand(),
		outImage:       imagick.NewMagickWand(),
		topTemplate:    imagick.NewMagickWand(),
		bottomTemplate: imagick.NewMagickWand(),
	}
}

func (d *drawer) LoadInImage(imageBlob []byte) {
	logrus.Info("Загружаем изображение")
	if err := d.inImage.ReadImageBlob(imageBlob); err != nil {
		logrus.Fatal(err)
	}
}

func (d *drawer) getInImageWidth() uint {
	return d.inImage.GetImageWidth()
}

func (d *drawer) getInImageHeight() uint {
	return d.inImage.GetImageHeight()
}

func (d *drawer) ConfigureFrameSizes() {
	logrus.Info("Формируем размеры рамки")
	d.margin = uint(float64(d.getInImageWidth()+d.getInImageHeight()) * float64(0.01))
	d.borderWidth = uint(float64(d.getInImageWidth()+d.getInImageHeight()) * float64(0.004))
	d.padding = uint(float64(d.getInImageWidth()+d.getInImageHeight()) * float64(0.002))
}

func (d *drawer) CreateTopTemplate() {
	logrus.Info("Создаем верхнюю часть изображения")
	pw := imagick.NewPixelWand()
	pw.SetColor("black")
	err := d.topTemplate.NewImage(
		uint(d.getInImageWidth()+((d.margin+d.borderWidth+d.padding)*2)),
		uint(d.getInImageHeight()+((d.margin+d.borderWidth+d.padding)*2)),
		pw,
	)
	if err != nil {
		logrus.Fatal(err)
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
		logrus.Fatal(err)
	}
}

func (d *drawer) MergeInImageToTopTemplate() {
	logrus.Info("Сливаем верхнюю часть и входящее изображение")
	err := d.topTemplate.CompositeImage(
		d.inImage, imagick.COMPOSITE_OP_OVER, true,
		int(d.margin+d.borderWidth+d.padding),
		int(d.margin+d.borderWidth+d.padding),
	)
	if err != nil {
		logrus.Fatal(err)
	}
}

func (d *drawer) CreateBottomTemplate(ftext, stext, fontPath string) {
	if len(ftext) == 0 || len(stext) == 0 {
		return
	}
	logrus.Info("Создаем нижнюю часть изображения")
	fontSize := 100.0

	pw := imagick.NewPixelWand()
	pw.SetColor("black")
	err := d.bottomTemplate.NewImage(d.topTemplate.GetImageWidth(), d.getInImageHeight(), pw)
	if err != nil {
		logrus.Error(err)
	}
	pw.Destroy()

	pw = imagick.NewPixelWand()
	pw.SetColor("white")
	dw := imagick.NewDrawingWand()
	if len(fontPath) != 0 {
		err = dw.SetFont(fontPath)
		if err != nil {
			logrus.Error(err)
		}
	}
	dw.SetFontSize(fontSize)
	dw.SetFillColor(pw)
	pw.Destroy()

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			logrus.Fatalf("%v, %s", panicInfo, string(debug.Stack()))
		}
	}()
	// error on set emoji font : QueryFontMetrics
	metrics1 := d.bottomTemplate.QueryFontMetrics(dw, ftext)
	metrics2 := d.bottomTemplate.QueryFontMetrics(dw, stext)

	for metrics1.TextWidth > float64(d.topTemplate.GetImageWidth()) ||
		metrics2.TextWidth > float64(d.topTemplate.GetImageWidth()) {
		fontSize -= 6
		dw.SetFontSize(fontSize)
		metrics1 = d.bottomTemplate.QueryFontMetrics(dw, ftext)
		metrics2 = d.bottomTemplate.QueryFontMetrics(dw, stext)
	}

	if len(ftext) == 0 {
		metrics1.TextHeight = 0
	}
	if len(stext) == 0 {
		metrics2.TextHeight = 0
	}

	err = d.bottomTemplate.ResizeImage(
		d.topTemplate.GetImageWidth(),
		uint(d.margin*2)+uint(metrics1.TextHeight)+uint(metrics2.TextHeight),
		imagick.FILTER_LANCZOS2,
	)
	if err != nil {
		logrus.Fatal(err)
	}

	dw.Annotation(
		float64(d.bottomTemplate.GetImageWidth())/2-metrics1.TextWidth/2,
		metrics1.TextHeight-float64(d.margin/2),
		ftext,
	)
	dw.Annotation(
		float64(d.bottomTemplate.GetImageWidth())/2-metrics2.TextWidth/2,
		metrics1.TextHeight-float64(d.margin/2)+metrics2.TextHeight,
		stext,
	)
	err = d.bottomTemplate.DrawImage(dw)
	if err != nil {
		logrus.Fatal(err)
	}
	dw.Destroy()
}

func (d *drawer) MergeTopAndBottomTemplates() {
	logrus.Info("Сливаем верхнюю и нижнюю части изображения")
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

func (d *drawer) GetBlob() []byte {
	return d.outImage.GetImageBlob()
}

func (d *drawer) SaveImage(outputPath string) {
	logrus.Infof("Сохраняем изображение по пути: %s", outputPath)
	if err := d.outImage.WriteImage(outputPath); err != nil {
		logrus.Fatal(err)
	}
}
