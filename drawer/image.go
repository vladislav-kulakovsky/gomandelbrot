package drawer

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type Image struct {
	BitMap [][]color.Color
	Width  int
	Height int
}

func NewImage(w, h int) *Image {
	bitmap := make([][]color.Color, h, h)
	for y := 0; y < h; y++ {
		bitmap[y] = make([]color.Color, w, w)
	}
	return &Image{
		BitMap: bitmap,
		Width:  w,
		Height: h,
	}
}

func (i *Image) Set(x, y int, c color.Color) {
	i.BitMap[y][x] = c
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i *Image) At(x, y int) color.Color {
	return i.BitMap[y][x]
}

func (i *Image) Draw(writer io.Writer) error {
	return png.Encode(writer, i)
}
