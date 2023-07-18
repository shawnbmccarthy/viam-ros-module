package main

import (
	"fmt"
	"image"
	"image/color"
	"github.com/shawnbmccarthy/viam-ros-module/utils"
)

type DumbImage struct {
	data []byte
	width int
	height int
	step int
}

func (di *DumbImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (di *DumbImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: di.height, Y: di.width},
	}
}

func (di *DumbImage) At(x, y int) color.Color {
	bytesPerPixel := di.step / di.width
	pixelOffset := di.width * x + y
	byteOffset := bytesPerPixel * pixelOffset

	fmt.Printf("(%d,%d) -> Bytes/Pixel: %d, Pixel Offset: %d, Byte Offset: %d, ", x, y, bytesPerPixel, pixelOffset, byteOffset)
	return color.RGBA{
		R: di.data[byteOffset],
		G: di.data[byteOffset+1],
		B: di.data[byteOffset+2],
		A: 0,
	}
}

func main() {
	di := DumbImage{width: utils.Width, height: utils.Height, data: utils.Data, step: utils.Step}

	for x := 0; x < utils.Height; x++ {
		for y := 0; y < utils.Width; y++ {
			color := di.At(x, y)
			fmt.Printf("%v\n", color)
		}
	}
	fmt.Printf("ros image len: %d\n", len(utils.Data))
	fmt.Printf("ros image height: %d, width: %d, step: %d\n", utils.Height, utils.Width, utils.Step)
}
