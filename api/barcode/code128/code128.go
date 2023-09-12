// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"reflect"
)

// BarColorTolerance determines which colors count as a bar.
// The r, g, b color channels (multiplied by a) are summed and normalized
// between 0 and 1.
// A pixel is a bar-pixel when the resulting value is less than or equal to
// BarColorTolerance.
var BarColorTolerance = 0.7

type Code128 struct{}

func (Code128) Decode(img image.Image) (bs []byte, err error) {
	widths, err := modules(img)
	if err != nil {
		return nil, err
	}

	_ = reverse(widths)
	qs, sta, d, c, stp, qe := segments(widths)
	_, _, _ = qs, qe, stp

	if len(d)%6 != 0 {
		return nil, errors.New("invalid data segment")
	}

	decodeTable := DecodeTableA
	charTable := CharTableA
	current, posMul := 5, 1

	defer func() {
		if r := recover(); r != nil {
			table := "?"
			if reflect.ValueOf(decodeTable).Pointer() == reflect.ValueOf(DecodeTableA).Pointer() {
				table = "A"
			}
			if reflect.ValueOf(decodeTable).Pointer() == reflect.ValueOf(DecodeTableB).Pointer() {
				table = "B"
			}
			if reflect.ValueOf(decodeTable).Pointer() == reflect.ValueOf(DecodeTableC).Pointer() {
				table = "C"
			}
			seq := ""
			for i := -5; current+i < len(d) && i <= 0; i++ {
				seq += fmt.Sprintf("%d", d[current+i])
			}
			err = fmt.Errorf("table %s: invalid sequence: %s", table, seq)
		}
	}()

	staSym := decodeTable[sta[0]][sta[1]][sta[2]][sta[3]][sta[4]][sta[5]]
	switch staSym {
	default:
		return nil, fmt.Errorf("invalid start symbol: %s -- %+v", charTable[staSym], sta)
	case START_A:
		decodeTable = DecodeTableA
		charTable = CharTableA
	case START_B:
		decodeTable = DecodeTableB
		charTable = CharTableB
	case START_C:
		decodeTable = DecodeTableC
		charTable = CharTableC
	}

	checksum := staSym

	for current < len(d) {
		sym := decodeTable[d[current-5]][d[current-4]][d[current-3]][d[current-2]][d[current-1]][d[current-0]]
		checksum += sym * posMul

		switch sym {
		default:
			bs = append(bs, []byte(charTable[sym])...)
		case CODE_A:
			decodeTable = DecodeTableA
			charTable = CharTableA
		case CODE_B:
			decodeTable = DecodeTableB
			charTable = CharTableB
		case CODE_C:
			decodeTable = DecodeTableC
			charTable = CharTableC
		case FNC3:
		case FNC2:
		case SHIFT_B:
		case FNC1:
		case START_A:
		case START_B:
		case START_C:
		case STOP:
		case REVERSE_STOP:
			return bs, fmt.Errorf("symbol %+v (%s) invalid in this position", sym, charTable[sym])
		}

		posMul++
		current += 6
	}

	checksum = checksum % 103
	cksmVal := DecodeTableA[c[0]][c[1]][c[2]][c[3]][c[4]][c[5]]
	cksmOK := cksmVal == checksum
	if !cksmOK {
		return bs, fmt.Errorf("invalid checksum: want: %d, got: %d", cksmVal, checksum)
	}

	return bs, nil
}

func (Code128) Encode(bs []byte) (image.Image, error) {
	return nil, errors.New("Not implemented")
}

func modules(img image.Image) (widths []int, err error) {
	var (
		isBar             = false // bar or space; start out expecting spaces (quiet zone)
		run               = 0     // length of current bar or space
		div               = 1     // divisor to normalize module widths
		divFound          = false // has the divisor been determined yet?
		quietSpaceMissing = false // many barcodes ignore the spec and omit quiet space
	)
	for x := 0; x < img.Bounds().Dx(); x++ {
		c := img.At(x, 0)
		r, g, b, _ := c.RGBA()
		l := float64(r + g + b) / 0x2FFFD // 0xFFFF * 3 = 0x2FFFD

		if !divFound && len(widths) == 2 {
			divFound = true

			// start symbol must start with 2-wide module
			div = widths[1] / 2

			// fixup previous runs
			widths[0] = widths[0] / div
			widths[1] = widths[1] / div
		}

		if l <= BarColorTolerance { // bar
			if isBar {
				run++
			} else { // new bar run begins; finish space run
				if run == 0 {
					// barcode didn't start with a quiet space!
					quietSpaceMissing = true
				}
				widths = append(widths, run/div)
				isBar = true
				run = 1
			}
		} else { // space
			if !isBar {
				run++
			} else { // new space run begins; finish bar run
				widths = append(widths, run/div)
				isBar = false
				run = 1
			}
		}
	}

	// don't forget to record last run!
	widths = append(widths, run/div)
	if quietSpaceMissing {
		widths = append(widths, 0)
	}
	return widths, nil
}

func reverse(widths []int) (isReversed bool) {
	startSym := widths[1:7]
	sym := DecodeTableA[startSym[0]][startSym[1]][startSym[2]][startSym[3]][startSym[4]][startSym[5]]
	if sym == REVERSE_STOP {
		isReversed = true
		for i, j := 0, len(widths)-1; i < j; i, j = i+1, j-1 {
			widths[i], widths[j] = widths[j], widths[i]
		}
	}
	return
}

func segments(widths []int) (quietStart int, startSym []int, data []int, checkSym []int, stopPat []int, quietEnd int) {
	quietStart = widths[0]
	startSym = widths[1:7]
	data = widths[7 : len(widths)-14]
	checkSym = widths[len(widths)-14 : len(widths)-8]
	stopPat = widths[len(widths)-8 : len(widths)-1]
	quietEnd = widths[len(widths)-1]
	return
}
