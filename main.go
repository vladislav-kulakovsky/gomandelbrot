package main

import (
	log "github.com/sirupsen/logrus"
	"mandelbrot/algo"
	"mandelbrot/drawer"
	"os"
)

func main() {
	drawer.Init()

	if drawer.GlobalConfig.Algorithm.ScaleFactor == 0 {
		drawer.GlobalConfig.Algorithm.ScaleFactor = drawer.AutoFitScaleFactor(drawer.GlobalConfig.Image.Resolution)
	}

	log.Info("starting, trying to open file")
	f, err := os.OpenFile("mandelbrot-sample.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.WithError(err).Fatal("cannot create/open file for writing")
	}
	defer f.Close()

	log.Info("starting calculations")
	y := int(float32(drawer.GlobalConfig.Image.Resolution) / 1.75)
	log.WithField("height", y).Info("calculated height of image")
	i := drawer.NewImage(drawer.GlobalConfig.Image.Resolution, y)
	log.WithField("scale-factor", drawer.GlobalConfig.Algorithm.ScaleFactor).Debug("chosen factor")
	m := algo.NewMandelbrot(
		drawer.GlobalConfig.Algorithm.Iterations,
		drawer.GlobalConfig.Algorithm.ScaleFactor,
		drawer.GlobalConfig.Image.Offset.X,
		drawer.GlobalConfig.Image.Offset.Y,
	)

	if drawer.GlobalConfig.Algorithm.Parallel {
		m.GenerateParallel(i)
	} else {
		m.Generate(i)
	}
	log.Info("done calculation, writing to file")
	err = i.Draw(f)
	if err != nil {
		log.WithError(err).Fatal("failed generate image")
	}
	log.Info("done!")
}
