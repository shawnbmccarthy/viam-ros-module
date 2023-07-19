package main

import (
	"fmt"
	"github.com/shawnbmccarthy/viam-ros-module/utils"
	"image"
	"image/color"
	"image/jpeg"
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
	color
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
		A: 255,
	}
}

func main() {
	di := DumbImage{width: utils.Width, height: utils.Height, data: utils.Data, step: utils.Step}

	for x := 0; x < utils.Height; x++ {
		for y := 0; y < utils.Width; y++ {
			di.At(x, y)
			//fmt.Printf("%v\n", c)
		}
	}
	fmt.Printf("ros image len: %d\n", len(utils.Data))
	fmt.Printf("ros image height: %d, width: %d, step: %d\n", utils.Height, utils.Width, utils.Step)
	//img, _, err := image.Decode(bytes.NewReader(utils.Data))
	//if err != nil {
	//	fmt.Printf("err: %v", err)
	//}

	out, _ := os.Create("./dummy_a255.jpg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 0

	err := jpeg.Encode(out, &di, &opts)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

}
