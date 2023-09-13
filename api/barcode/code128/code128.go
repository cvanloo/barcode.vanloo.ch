// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"reflect"
)

type Code128 struct {
	*image.Gray16
}

func (c Code128) Scale(width, height int) (image.Image, error) {
	oldWidth := c.Bounds().Dx()
	if width < oldWidth {
		return nil, errors.New("unable to shrink image, new width too small")
	}

	scaledImage := image.NewGray16(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	// TODO: don't generate a quiet zone in Encode. Do this here.
	//   Never make a quiet zone bigger than necessary, use as much space as
	//   possible for the barcode.

	// scale width
	scale := width / oldWidth
	qz := (width % oldWidth) / 2
	for x := 0; x < qz; x++ { // extend quiet zone start
		scaledImage.SetGray16(x, 0, color.White)
	}
	for x := 0; x < oldWidth; x++ { // copy pixels, scale them
		for s := 0; s < scale; s++ {
			scaledImage.SetGray16(qz+s+x*scale, 0, c.Gray16At(x, 0))
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

func Encode(text string) (Code128, error) {
	runes := []rune(text)

	// - each symbol is encoded using 6 modules
	// - a module is 1, 2, 3, or 4 units in size
	// - the six modules making up a symbol have a total size of 11 units
	// - we define a unit as 1 pixel
	//
	// quiet 10px + start 11px + 11px*len + checksum 11px + stop 13px + quiet 10px = 55px
	width, height := 11*len(runes)+55, 1
	img := image.NewGray16(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	var (
		xPos  int
		table TableIndex
		cksm  *Checksum
	)

	{ // draw quiet space
		for i := 0; i < 10; i++ {
			img.SetGray16(xPos, 0, color.White)
			xPos++
		}
	}

	{ // draw start symbol
		table = determineTable(runes, LookupNone)
		var startSym int
		switch table {
		case LookupA:
			startSym = START_A
		case LookupB:
			startSym = START_B
		case LookupC:
			startSym = START_C
		}
		bits := Bitpattern[startSym-SpecialOffset]
		drawBits(img, bits[3:9], &xPos)
		cksm = NewChecksum(startSym - SpecialOffset)
	}

	{ // draw data symbols
		for idx, r := range runes {
			nextTable := determineTable(runes[idx:], table)

			if nextTable != table {
				// TODO: Shift B/Shift A
				var code int
				switch nextTable {
				case LookupA:
					code = CODE_A
				case LookupB:
					code = CODE_B
				case LookupC:
					code = CODE_C
				}
				bits := Bitpattern[code-SpecialOffset]
				drawBits(img, bits[3:9], &xPos)
				cksm.Add(code - SpecialOffset)

				table = nextTable
			}

			bits, val := must2(lookup(r, table))
			drawBits(img, bits[3:9], &xPos)
			cksm.Add(val)
		}
	}

	{ // draw checksum
		bits := Bitpattern[cksm.Sum()]
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

	return Code128{img}, nil
}

func determineTable(nextText []rune, currentTable TableIndex) TableIndex {
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
		return LookupB
	}
	if isNumber(nextText[0]) {
		if len(nextText) > 1 && isNumber(nextText[1]) {
			return LookupC
		}
		return LookupB
	}
	if isSpecial(nextText[0]) {
		return LookupA
	}

	panic("unreachable (hopefully)")
}

func lookup(r rune, table TableIndex) (bits []int, val int, err error) {
	for i, bits := range Bitpattern {
		if bits[table] == int(r) {
			return bits, i, nil
		}
	}
	return nil, -1, fmt.Errorf("invalid rune %U (`%s') in table %s", r, string(r), string(rune(table+0x41)))
}

func drawBits(img *image.Gray16, bits []int, startX *int) {
	for i, w := range bits {
		if i%2 == 0 { // draw bar
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

type Checksum struct {
	Value int
	Idx   int
}

func NewChecksum(initial int) *Checksum {
	return &Checksum{Value: initial, Idx: 1}
}

func (c *Checksum) Add(val int) {
	c.Value += val * c.Idx
	c.Idx++
}

func (c *Checksum) Sum() int {
	c.Value %= 103
	return c.Value
}

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

// BarColorTolerance determines which colors count as a bar.
// The r, g, b color channels (multiplied by a) are summed and normalized
// between 0 and 1.
// A pixel is a bar-pixel when the resulting value is less than or equal to
// BarColorTolerance.
var BarColorTolerance = 0.7

func Decode(img image.Image) (bs []byte, err error) {
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
	case SYM_START_A:
		decodeTable = DecodeTableA
		charTable = CharTableA
	case SYM_START_B:
		decodeTable = DecodeTableB
		charTable = CharTableB
	case SYM_START_C:
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
		case SYM_CODE_A:
			decodeTable = DecodeTableA
			charTable = CharTableA
		case SYM_CODE_B:
			decodeTable = DecodeTableB
			charTable = CharTableB
		case SYM_CODE_C:
			decodeTable = DecodeTableC
			charTable = CharTableC
		case SYM_FNC3:
		case SYM_FNC2:
		case SYM_SHIFT_B:
		case SYM_FNC1:
		case SYM_START_A:
		case SYM_START_B:
		case SYM_START_C:
		case SYM_STOP:
		case SYM_REVERSE_STOP:
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
	sym := DecodeTableA[startSym[0]][startSym[1]][startSym[2]][startSym[3]][startSym[4]][startSym[5]]
	if sym == SYM_REVERSE_STOP {
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
