// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"reflect"
)

const (
	SPACE             = 0x00
	EXCLAMATION       = 0x01
	DOUBLE_QUOTE      = 0x02
	POUND             = 0x03
	DOLLAR            = 0x04
	PERCENT           = 0x05
	AMPERSAND         = 0x06
	SINGLE_QUOTE      = 0x07
	OPEN_PARENTHESIS  = 0x08
	CLOSE_PARENTHESIS = 0x09
	ASTERISK          = 0x0a
	PLUS              = 0x0b
	COMMA             = 0x0c
	HYPHEN            = 0x0d
	PERIOD            = 0x0e
	SLASH             = 0x0f
	ZERO              = 0x10
	ONE               = 0x11
	TWO               = 0x12
	THREE             = 0x13
	FOUR              = 0x14
	FIVE              = 0x15
	SIX               = 0x16
	SEVEN             = 0x17
	EIGHT             = 0x18
	NINE              = 0x19
	COLON             = 0x1a
	SEMICOLON         = 0x1b
	LESS_THAN         = 0x1c
	EQUAL             = 0x1d
	GREATER_THAN      = 0x1e
	QUESTION          = 0x1f
	AT                = 0x20
	A                 = 0x21
	B                 = 0x22
	C                 = 0x23
	D                 = 0x24
	E                 = 0x25
	F                 = 0x26
	G                 = 0x27
	H                 = 0x28
	I                 = 0x29
	J                 = 0x2a
	K                 = 0x2b
	L                 = 0x2c
	M                 = 0x2d
	N                 = 0x2e
	O                 = 0x2f
	P                 = 0x30
	Q                 = 0x31
	R                 = 0x32
	S                 = 0x33
	T                 = 0x34
	U                 = 0x35
	V                 = 0x36
	W                 = 0x37
	X                 = 0x38
	Y                 = 0x39
	Z                 = 0x3a

	CODE_C = 0x63
	CODE_B = 0x64
	CODE_A = 0x65

	START_A = 0x67
	START_B = 0x68
	START_C = 0x69
	STOP    = 0x6a

	REVERSE_STOP = -1
)

var DecodeTableA = [][][][][][]int{
	1: {
		1: {
			1: {
				3: {
					2: {
						3: A,
					},
				},
			},
			2: {
				1: {
					3: {
						3: J,
					},
				},
				2: {
					3: {
						2: COMMA,
					},
				},
				3: {
					1: {
						3: D,
					},
					3: {
						1: K,
					},
				},
			},
			3: {
				1: {
					2: {
						3: M,
					},
				},
				2: {
					2: {
						2: SLASH,
					},
				},
				3: {
					2: {
						1: N,
					},
				},
			},
		},
		2: {
			1: {
				2: {
					2: {
						3: POUND,
					},
				},
				3: {
					2: {
						2: DOLLAR,
					},
				},
			},
			2: {
				1: {
					3: {
						2: HYPHEN,
					},
				},
				2: {
					1: {
						3: AMPERSAND,
					},
					3: {
						1: PERIOD,
					},
				},
				3: {
					1: {
						2: SINGLE_QUOTE,
					},
				},
			},
			3: {
				1: {
					2: {
						2: ZERO,
					},
				},
				2: {
					2: {
						1: ONE,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: B,
					},
				},
				2: {
					2: {
						2: PERCENT,
					},
				},
				3: {
					2: {
						1: C,
					},
				},
			},
			2: {
				1: {
					1: {
						3: E,
					},
					3: {
						1: L,
					},
				},
				2: {
					1: {
						2: OPEN_PARENTHESIS,
					},
				},
				3: {
					1: {
						1: F,
					},
				},
			},
			3: {
				1: {
					2: {
						1: O,
					},
				},
			},
		},
	},
	2: {
		1: {
			1: {
				1: {
					3: {
						3: REVERSE_STOP,
					},
				},
				2: {
					1: {
						4: START_B,
					},
					3: {
						2: START_C,
					},
				},
				3: {
					1: {
						3: G,
					},
					3: {
						1: Q,
					},
				},
				4: {
					1: {
						2: START_A,
					},
				},
			},
			2: {
				1: {
					2: {
						3: GREATER_THAN,
					},
				},
				2: {
					2: {
						2: SPACE,
					},
				},
				3: {
					2: {
						1: QUESTION,
					},
				},
			},
			3: {
				1: {
					1: {
						3: S,
					},
					3: {
						1: U,
					},
				},
				2: {
					1: {
						2: FIVE,
					},
				},
				3: {
					1: {
						1: T,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					3: {
						2: THREE,
					},
				},
				2: {
					1: {
						3: CLOSE_PARENTHESIS,
					},
					3: {
						1: FOUR,
					},
				},
				3: {
					1: {
						2: ASTERISK,
					},
				},
			},
			2: {
				1: {
					2: {
						2: EXCLAMATION,
					},
				},
				2: {
					2: {
						1: DOUBLE_QUOTE,
					},
				},
			},
			3: {
				1: {
					1: {
						2: SIX,
					},
				},
				2: {
					1: {
						1: TWO,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						3: H,
					},
					3: {
						1: R,
					},
				},
				2: {
					1: {
						2: PLUS,
					},
				},
				3: {
					1: {
						1: I,
					},
				},
			},
			2: {
				1: {
					2: {
						1: AT,
					},
				},
			},
		},
	},
	3: {
		1: {
			1: {
				1: {
					2: {
						3: V,
					},
				},
				2: {
					2: {
						2: EIGHT,
					},
				},
				3: {
					2: {
						1: W,
					},
				},
			},
			2: {
				1: {
					1: {
						3: Y,
					},
					3: {
						1: SEVEN,
					},
				},
				2: {
					1: {
						2: SEMICOLON,
					},
				},
				3: {
					1: {
						1: Z,
					},
				},
			},
			3: {
				1: {
					2: {
						1: P,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						2: NINE,
					},
				},
				2: {
					2: {
						1: COLON,
					},
				},
			},
			2: {
				1: {
					1: {
						2: LESS_THAN,
					},
				},
				2: {
					1: {
						1: EQUAL,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						1: X,
					},
				},
			},
		},
	},
}

var DecodeTableB = [][][][][][]int{
	1: {
		1: {
			1: {
				3: {
					2: {
						3: A,
					},
				},
			},
			2: {
				1: {
					3: {
						3: J,
					},
				},
				3: {
					1: {
						3: D,
					},
					3: {
						1: K,
					},
				},
			},
			3: {
				1: {
					2: {
						3: M,
					},
				},
				3: {
					2: {
						1: N,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: B,
					},
				},
				3: {
					2: {
						1: C,
					},
				},
			},
			2: {
				1: {
					1: {
						3: E,
					},
					3: {
						1: L,
					},
				},
				3: {
					1: {
						1: F,
					},
				},
			},
			3: {
				1: {
					2: {
						1: O,
					},
				},
			},
		},
	},
	2: {
		1: {
			1: {
				1: {
					3: {
						3: REVERSE_STOP,
					},
				},
				2: {
					1: {
						4: START_B,
					},
					3: {
						2: START_C,
					},
				},
				3: {
					1: {
						3: G,
					},
					3: {
						1: Q,
					},
				},
				4: {
					1: {
						2: START_A,
					},
				},
			},
			3: {
				1: {
					1: {
						3: S,
					},
					3: {
						1: U,
					},
				},
				3: {
					1: {
						1: T,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						3: H,
					},
					3: {
						1: R,
					},
				},
				3: {
					1: {
						1: I,
					},
				},
			},
		},
	},
	3: {
		1: {
			1: {
				1: {
					2: {
						3: V,
					},
				},
				3: {
					2: {
						1: W,
					},
				},
			},
			2: {
				1: {
					1: {
						3: Y,
					},
				},
				3: {
					1: {
						1: Z,
					},
				},
			},
			3: {
				1: {
					2: {
						1: P,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						1: X,
					},
				},
			},
		},
	},
}

var DecodeTableC = [][][][][][]int{}

var ASCIITable = map[int]string{
	SPACE:             " ",
	EXCLAMATION:       "!",
	DOUBLE_QUOTE:      `"`,
	POUND:             "#",
	DOLLAR:            "$",
	PERCENT:           "%",
	AMPERSAND:         "&",
	SINGLE_QUOTE:      "'",
	OPEN_PARENTHESIS:  "(",
	CLOSE_PARENTHESIS: ")",
	ASTERISK:          "*",
	PLUS:              "+",
	COMMA:             ",",
	HYPHEN:            "-",
	PERIOD:            ".",
	SLASH:             "/",
	ZERO:              "0",
	ONE:               "1",
	TWO:               "2",
	THREE:             "3",
	FOUR:              "4",
	FIVE:              "5",
	SIX:               "6",
	SEVEN:             "7",
	EIGHT:             "8",
	NINE:              "9",
	COLON:             ":",
	SEMICOLON:         ";",
	LESS_THAN:         "<",
	EQUAL:             "=",
	GREATER_THAN:      ">",
	QUESTION:          "?",
	AT:                "@",
	A:                 "A",
	B:                 "B",
	C:                 "C",
	D:                 "D",
	E:                 "E",
	F:                 "F",
	G:                 "G",
	H:                 "H",
	I:                 "I",
	J:                 "J",
	K:                 "K",
	L:                 "L",
	M:                 "M",
	N:                 "N",
	O:                 "O",
	P:                 "P",
	Q:                 "Q",
	R:                 "R",
	S:                 "S",
	T:                 "T",
	U:                 "U",
	V:                 "V",
	W:                 "W",
	X:                 "X",
	Y:                 "Y",
	Z:                 "Z",

	CODE_C: "<CODE_C>",
	CODE_B: "<CODE_B>",
	CODE_A: "<CODE_A>",

	START_A: "<START_A>",
	START_B: "<START_B>",
	START_C: "<START_C>",
	STOP:    "<STOP>",

	REVERSE_STOP: "<REVERSE_STOP>",
}

func Widths(img image.Image) (widths []int, err error) {
	bars := false
	run := 0
	div := 1
	divFound := false
	quietSpaceMissing := false
	for x := 0; x < img.Bounds().Dx(); x++ {
		c := img.At(x, 0)
		r, g, b, _ := c.RGBA()
		_, _ = g, b

		if !divFound && len(widths) == 2 {
			divFound = true
			div = widths[1] / 2

			fmt.Printf("determined div as %d\n", div)

			// fixup previous runs
			widths[0] = widths[0] / div
			widths[1] = widths[1] / div
		}

		if r == 0x0000 {
			if bars {
				run++
			} else {
				// finish space run
				if run == 0 && !divFound {
					// barcode didn't start with a quiet space!
					widths = append(widths, 0)
					quietSpaceMissing = true
				}
				if run != 0 {
					widths = append(widths, run/div)
				}
				bars = true
				run = 1
			}
		} else if r == 0xFFFF {
			if !bars {
				run++
			} else {
				// finish bar run
				if run != 0 { // @fixme: check unnecessary
					widths = append(widths, run/div)
				}
				bars = false
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

func Reverse(widths []int) (nws []int, rev bool) {
	nws = widths
	startSym := widths[1:7]
	sym := DecodeTableA[startSym[0]][startSym[1]][startSym[2]][startSym[3]][startSym[4]][startSym[5]]
	if sym == REVERSE_STOP {
		rev = true
		for i, j := 0, len(widths)-1; i < j; i, j = i+1, j-1 {
			nws[i], nws[j] = nws[j], nws[i]
		}
	}
	return
}

func Split(widths []int) (quietStart int, startSym []int, data []int, checkSym []int, stopPat []int, quietEnd int) {
	quietStart = widths[0]
	startSym = widths[1:7]
	data = widths[7 : len(widths)-14]
	checkSym = widths[len(widths)-14 : len(widths)-8]
	stopPat = widths[len(widths)-8 : len(widths)-1]
	quietEnd = widths[len(widths)-1]
	return
}

func Decode(img image.Image) (msg string, err error) {
	widths, err := Widths(img)
	if err != nil {
		return msg, err
	}
	fmt.Printf("%+v\n", widths)

	widths, reversed := Reverse(widths)
	if reversed {
		fmt.Println("reading in reverse!")
	}

	qs, sta, d, c, stp, qe := Split(widths)
	fmt.Printf("qs: %d\nsta: %+v\nd: %+v\nc: %+v\nstp: %+v\nqe: %d\n", qs, sta, d, c, stp, qe)

	if len(d)%6 != 0 {
		return msg, errors.New("invalid data segment")
	}

	decodeTable := DecodeTableA
	current, posMul := 5, 1

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
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
			fmt.Printf("(Table %s) Unable to parse sequence:", table)
			for i := -5; current+i < len(d) && i <= 0; i++ {
				fmt.Printf(" %d", d[current+i])
			}
			fmt.Println()
			fmt.Printf("More: %+v\n", d[current+1:])
		}
	}()

	staSym := decodeTable[sta[0]][sta[1]][sta[2]][sta[3]][sta[4]][sta[5]]
	switch staSym {
	case START_A:
		decodeTable = DecodeTableA
	case START_B:
		decodeTable = DecodeTableB
	case START_C:
		decodeTable = DecodeTableC
	default:
		return msg, fmt.Errorf("invalid start symbol: %s -- %+v", ASCIITable[staSym], sta)
	}

	checksum := staSym

	for current < len(d) {
		sym := decodeTable[d[current-5]][d[current-4]][d[current-3]][d[current-2]][d[current-1]][d[current-0]]
		checksum += sym * posMul
		fmt.Printf("Sym: %s (%d%d%d%d%d%d) [%d × %d = %d]\n", ASCIITable[sym], d[current-5], d[current-4], d[current-3], d[current-2], d[current-1], d[current-0], sym, posMul, sym*posMul)

		switch sym {
		case CODE_A:
			decodeTable = DecodeTableA
		case CODE_B:
			decodeTable = DecodeTableB
		case CODE_C:
			decodeTable = DecodeTableC
		default:
			msg += string(ASCIITable[sym])
		}

		posMul++
		current += 6
	}

	checksum = checksum % 103
	cksmVal := DecodeTableA[c[0]][c[1]][c[2]][c[3]][c[4]][c[5]]
	cksmOK := cksmVal == checksum
	fmt.Printf("Checksum: %d (expected: %d, ok: %t)\n", checksum, cksmVal, cksmOK)

	if !cksmOK {
		return msg, fmt.Errorf("invalid checksum: want: %d, got: %d", cksmVal, checksum)
	}

	return msg, nil
}

func Encode(text string) (image.Image, error) {
	return nil, errors.New("Not implemented")
}
