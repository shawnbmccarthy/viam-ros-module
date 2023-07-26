package main

import (
	"bytes"
	"fmt"
	"github.com/shawnbmccarthy/viam-ros-module/utils"
	"image"
	"image/color"
	"image/png"
	"os"
)

type DumbImage struct {
	data   []byte
	width  int
	height int
	step   int
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
	pixelOffset := di.width*x + y
	byteOffset := bytesPerPixel * pixelOffset

	return color.RGBA{
		R: di.data[byteOffset+2],
		G: di.data[byteOffset+1],
		B: di.data[byteOffset],
		A: 0,
	}
}

func (di *DumbImage) Write(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		di.data = append(di.data, p[i])
	}
	fmt.Printf("Wrote: %d bytes", len(di.data))
	return len(di.data), nil
}

func main() {
	di := DumbImage{width: utils.Width, height: utils.Height, data: utils.Data, step: utils.Step}
	fmt.Printf("ros image len: %d\n", len(utils.Data))
	fmt.Printf("ros image height: %d, width: %d, step: %d\n", utils.Height, utils.Width, utils.Step)

	f, _ := os.Create("./di2.png")


	buffer := bytes.Buffer{}
	pngEncoder := png.Encoder{CompressionLevel: png.BestCompression}

	_ = pngEncoder.Encode(&buffer, &di)
	_, _ = f.Write(buffer.Bytes())
}
