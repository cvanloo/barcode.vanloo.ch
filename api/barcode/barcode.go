package barcode

import (
	"errors"
	"fmt"
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

func Scale(img image.Image, width, height int) (image.Image, error) {
	// TODO: change API so that these assertions will always hold
	grayImage, ok := img.(*image.Gray16)
	if !ok {
		return nil, errors.New("not a grayscale image")
	}
	oldWidth := grayImage.Bounds().Dx()
	if width < oldWidth {
		return nil, errors.New("new width must be greater or equal to the old width")
	}

	scaledImage := image.NewGray16(image.Rectangle{image.Point{0,0}, image.Point{width, height}})

	// scale width
	scale := width / oldWidth
	fmt.Printf("scale: %d\n", scale)
	for x := 0; x < oldWidth; x++ {
		for s := 0; s < scale; s++ {
			scaledImage.SetGray16(x*scale+s, 0, grayImage.Gray16At(x, 0))
		}
	}

	// scale height
	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			scaledImage.SetGray16(x, y, scaledImage.Gray16At(x, 0))
		}
	}

	return scaledImage, nil
}
