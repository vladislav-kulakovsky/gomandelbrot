package algo

import (
	"image/color"
	"mandelbrot/drawer"
	"math/cmplx"
	"sync"
)

type Mandelbrot struct {
	iterations  int
	scaleFactor float32
	offsetX     float32
	offsetY     float32
}

func NewMandelbrot(iterations int, scaleFactor float32, offsetX, offsetY float32) drawer.MandelbrotGenerator {
	return Mandelbrot{iterations, scaleFactor, offsetX, offsetY}
}

func (m Mandelbrot) Generate(canvas *drawer.Image) error {
	for iy := 0; iy < canvas.Height; iy++ {
		y := m.scaleFactor*(float32(iy)-float32(canvas.Height)/2.0) - m.offsetY

		for ix := 0; ix < canvas.Width; ix++ {
			x := m.scaleFactor*(float32(ix)-float32(canvas.Width)/2.0) - m.offsetX
			canvas.Set(ix, iy, m.GetColor(complex(float64(x), float64(y))))
		}
	}

	return nil
}

func (m Mandelbrot) GenerateParallel(canvas *drawer.Image) error {
	var waitGroup sync.WaitGroup

	for iy := 0; iy < canvas.Height; iy++ {
		y := m.scaleFactor*(float32(iy)-float32(canvas.Height)/2.0) - m.offsetY

		for ix := 0; ix < canvas.Width; ix++ {
			x := m.scaleFactor*(float32(ix)-float32(canvas.Width)/2.0) - m.offsetX
			waitGroup.Add(1)

			go func(ix, iy int) {
				canvas.Set(ix, iy, m.GetColor(complex(float64(x), float64(y))))
				waitGroup.Done()
			}(ix, iy)
		}
	}
	waitGroup.Wait()
	return nil
}

func (m Mandelbrot) GetColor(c complex128) color.Color {
	z := c
	i := 0

	for ; i < m.iterations && cmplx.Abs(z) <= 2; i++ {
		z = cmplx.Pow(z, 2) + c
	}

	if cmplx.Abs(z) > 2 {
		return color.RGBA { 255 - 15 * uint8(i), 15 * uint8(i), 255 - 15 * uint8(i), 255 }
	}

	return color.Black
}
