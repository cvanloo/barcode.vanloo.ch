// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"reflect"
	"image/color"

	"github.com/cvanloo/barcode/code128/encoding"
	"github.com/cvanloo/barcode/code128/decoding"
)

// BarColorTolerance determines which colors count as a bar.
// The r, g, b color channels (multiplied by a) are summed and normalized
// between 0 and 1.
// A pixel is a bar-pixel when the resulting value is less than or equal to
// BarColorTolerance.
var BarColorTolerance = 0.7

type Code128 struct{}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func must2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v1, v2
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
		l := 0.0
		for y := 0; y < img.Bounds().Dy(); y++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			v := float64(r+g+b) / 0x2FFFD // 0xFFFF * 3 = 0x2FFFD

			// calculate incremental average
			// v/1.1 -- add a bias towards lower values
			// @todo: use something better: https://stackoverflow.com/questions/48395434/how-to-crop-or-remove-white-background-from-an-image
			//   https://www.imageprocessingplace.com/downloads_V3/root_downloads/tutorials/contour_tracing_Abeer_George_Ghuneim/alg.html
			//   https://en.wikipedia.org/wiki/Edge_detection
			l = l + (v/1.1-l)/float64(y+1)
		}

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
	sym := decoding.DecodeTableA[startSym[0]][startSym[1]][startSym[2]][startSym[3]][startSym[4]][startSym[5]]
	if sym == decoding.REVERSE_STOP {
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

	decodeTable := decoding.DecodeTableA
	charTable := decoding.CharTableA
	current, posMul := 5, 1

	defer func() {
		if r := recover(); r != nil {
			table := "?"
			if reflect.ValueOf(decodeTable).Pointer() == reflect.ValueOf(decoding.DecodeTableA).Pointer() {
				table = "A"
			}
			if reflect.ValueOf(decodeTable).Pointer() == reflect.ValueOf(decoding.DecodeTableB).Pointer() {
				table = "B"
			}
			if reflect.ValueOf(decodeTable).Pointer() == reflect.ValueOf(decoding.DecodeTableC).Pointer() {
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
	case decoding.START_A:
		decodeTable = decoding.DecodeTableA
		charTable = decoding.CharTableA
	case decoding.START_B:
		decodeTable = decoding.DecodeTableB
		charTable = decoding.CharTableB
	case decoding.START_C:
		decodeTable = decoding.DecodeTableC
		charTable = decoding.CharTableC
	}

	checksum := staSym
	fmt.Printf("CKSM: %d\n", checksum)

	for current < len(d) {
		sym := decodeTable[d[current-5]][d[current-4]][d[current-3]][d[current-2]][d[current-1]][d[current-0]]
		checksum += sym * posMul
		fmt.Printf("CKSM: %d\n", checksum)

		switch sym {
		default:
			bs = append(bs, []byte(charTable[sym])...)
		case decoding.CODE_A:
			decodeTable = decoding.DecodeTableA
			charTable = decoding.CharTableA
		case decoding.CODE_B:
			decodeTable = decoding.DecodeTableB
			charTable = decoding.CharTableB
		case decoding.CODE_C:
			decodeTable = decoding.DecodeTableC
			charTable = decoding.CharTableC
		case decoding.FNC3:
		case decoding.FNC2:
		case decoding.SHIFT_B:
		case decoding.FNC1:
		case decoding.START_A:
		case decoding.START_B:
		case decoding.START_C:
		case decoding.STOP:
		case decoding.REVERSE_STOP:
			return bs, fmt.Errorf("symbol %+v (%s) invalid in this position", sym, charTable[sym])
		}

		posMul++
		current += 6
	}

	fmt.Println(string(bs))

	checksum = checksum % 103
	cksmVal := decoding.DecodeTableA[c[0]][c[1]][c[2]][c[3]][c[4]][c[5]]
	cksmOK := cksmVal == checksum
	if !cksmOK {
		return bs, fmt.Errorf("invalid checksum: want: %d, got: %d", cksmVal, checksum)
	}

	return bs, nil
}

func lookup(r rune, table encoding.TableIndex) (bits []int, idx int, err error) {
	for i, bits := range encoding.Bitpatterns {
		if bits[table] == int(r) {
			return bits, i, nil
		}
	}
	return nil, -1, fmt.Errorf("invalid rune %U (`%s') in table %s", r, string(r), string(rune(table+0x41)))
}

func determineTable(nextText []rune, currentTable encoding.TableIndex) encoding.TableIndex {
	// ~$ man 7 ascii
	isAsciiPrintable := func(r rune) bool {
		return r >= 0x20 /* space */ && r <= 0x7F /* DEL */
	}
	isNumber := func(r rune) bool {
		return r >= 0x30 /* 0 */ && r <= 0x39 /* 9 */
	}
	isSpecial := func(r rune) bool {
		return r >= 0x00 /* NUL */ && r <= 0x1F /* US */
	}

	// TODO: improve algorithm for more efficient encoding (minimize table switching)
	//   E.g., Does it make sense to switch to C, or should we just stay in A/B?
	//   Should we use a Shift B/Shift A or Code B/Code A?
	if isAsciiPrintable(nextText[0]) {
		return encoding.LookupB
	}
	if isNumber(nextText[0]) {
		if len(nextText) > 1 && isNumber(nextText[1]) {
			return encoding.LookupC
		}
		return encoding.LookupB
	}
	if isSpecial(nextText[0]) {
		return encoding.LookupA
	}

	panic("unreachable (hopefully)")
}

func drawBits(img *image.Gray16, bits []int, startX *int) {
	for i, w := range bits {
		if i % 2 == 0 { // draw bar
			for j := 0; j < w; j++ {
				img.SetGray16(*startX+j, 0, color.Black)
			}
		} else { // draw space
			for j := 0; j < w; j++ {
				img.SetGray16(*startX+j, 0, color.White)
			}
		}
		*startX += w
	}
}

func (Code128) Encode(bs []byte) (image.Image, error) {
	text := string(bs)

	// 11*len(bs) -- each symbol encoded with 6 modules, where all 6 modules
	// together must be 11 units (in our case 1 unit = 1 pixel) wide.
	//
	// quietspace(10px) + start(11px) + checksum(11px) + stoppattern(13px) + quietspace(10px) = 55px
	width, height := 11*len(bs)+55, 1
	img := image.NewGray16(image.Rectangle{image.Point{0,0}, image.Point{width, height}})

	var (
		xPos, checksum, ckIdx int
		table encoding.TableIndex
	)

	{ // draw quiet space
		for i := 0; i < 10; i++ {
			img.SetGray16(xPos, 0, color.White)
			xPos++
		}
	}

	{ // draw start symbol
		table = determineTable([]rune(text[:]), encoding.LookupUninit)
		var startSym int
		switch table {
		case encoding.LookupA:
			startSym = encoding.START_A
		case encoding.LookupB:
			startSym = encoding.START_B
		case encoding.LookupC:
			startSym = encoding.START_C
		}
		bits := encoding.Bitpatterns[startSym-32]
		drawBits(img, bits[3:9], &xPos)
		checksum = startSym - 32
		fmt.Printf("CKSM: %d\n", checksum)
		ckIdx = 1
	}

	{ // draw data symbols
		for idx, r := range text {
			nextTable := determineTable([]rune(text[idx:]), table)

			if nextTable != table {
				// TODO: Shift B/Shift A
				var code int
				switch nextTable {
				case encoding.LookupA:
					code = encoding.CODE_A
				case encoding.LookupB:
					code = encoding.CODE_B
				case encoding.LookupC:
					code = encoding.CODE_C
				}
				bits := encoding.Bitpatterns[code-32]
				drawBits(img, bits[3:9], &xPos)
				checksum += (code-32) * ckIdx
				fmt.Printf("CKSM: %d\n", checksum)
				ckIdx++

				table = nextTable
			}

			bits, val := must2(lookup(r, table))
			drawBits(img, bits[3:9], &xPos)
			checksum += val * ckIdx
			fmt.Printf("CKSM: %d\n", checksum)
			ckIdx++
		}
	}

	{ // draw checksum
		checksum %= 103
		bits := encoding.Bitpatterns[checksum]
		drawBits(img, bits[3:9], &xPos)
	}

	{ // draw stop pattern
		drawBits(img, []int{2, 3, 3, 1, 1, 1, 2}, &xPos)
	}

	{ // draw quiet space
		for i := 0; i < 10; i++ {
			img.SetGray16(xPos, 0, color.White)
			xPos++
		}
	}

	return img, nil
}
