package barcode

import (
	"image/color"
	"errors"
	"fmt"
	"image"
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
	qz := (width - oldWidth) / 2
	fmt.Printf("scale: %d\n", scale)
	for x := 0; x < qz; x++ { // extend quiet zone start
		scaledImage.SetGray16(x, 0, color.White)
	}
	for x := 0; x < oldWidth; x++ {
		for s := 0; s < scale; s++ {
			scaledImage.SetGray16(qz+s+x*scale, 0, grayImage.Gray16At(x, 0))
		}
	}
	for x := 0; x < qz; x++ { // extend quiet zone end
		scaledImage.SetGray16(width-x-1, 0, color.White)
	}

	// scale height
	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			scaledImage.SetGray16(x, y, scaledImage.Gray16At(x, 0))
		}
	}

	return scaledImage, nil
}
