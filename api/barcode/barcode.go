package barcode

import (
	"github.com/cvanloo/barcode/code128"
	"image"
)

type Barcode interface {
	Encode(bs []byte) (image.Image, error)
	Decode(img image.Image) (bs []byte, err error)
}

var (
	Code128 Barcode = code128.Code128{}
)

func Scale(img image.Image, width, height int) image.Image {
	_ = width // TODO
	grayImg := img.(*image.Gray16)
	scaledImage := image.NewGray16(image.Rectangle{image.Point{0,0}, image.Point{grayImg.Bounds().Dx(), height}})
	for y := 0; y < height; y++ {
		for x := 0; x < grayImg.Bounds().Dx(); x++ {
			scaledImage.SetGray16(x, y, grayImg.Gray16At(x, 0))
		}
	}
	return scaledImage
}
