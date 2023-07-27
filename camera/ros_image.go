package camera

import (
	"image"
	"image/color"
)

type RosImage struct {
	width  int
	height int
	step   int
	data   []byte
}

func (rosImage *RosImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (rosImage *RosImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: rosImage.height, Y: rosImage.width},
	}
}

func (rosImage *RosImage) At(x, y int) color.RGBA {
	bytesPerPixel := rosImage.step / rosImage.width
	pixelOffset := rosImage.width*x + y
	byteOffset := bytesPerPixel * pixelOffset

	return color.RGBA{
		R: rosImage.data[byteOffset+2],
		G: rosImage.data[byteOffset+1],
		B: rosImage.data[byteOffset],
		A: 0,
	}
}
