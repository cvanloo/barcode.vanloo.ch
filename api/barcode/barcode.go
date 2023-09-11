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
