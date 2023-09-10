// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"reflect"
)

const (
	// character set A
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
	OPEN_BRACKET      = 0x3b
	BACKSLASH         = 0x3c
	CLOSE_BRACKET     = 0x3d
	CARET             = 0x3e
	UNDERSCORE        = 0x3f
	NUL               = 0x40
	SOH               = 0x41
	STX               = 0x42
	ETX               = 0x43
	EOT               = 0x44

	ENQ = 0x45
	ACK = 0x46
	BEL = 0x47
	BS  = 0x48
	HT  = 0x49
	LF  = 0x4a
	VT  = 0x4b
	FF  = 0x4c
	CR  = 0x4d
	SO  = 0x4e
	SI  = 0x4f
	DLE = 0x50

	DC1     = 0x51
	DC2     = 0x52
	DC3     = 0x53
	DC4     = 0x54
	NAK     = 0x55
	SYN     = 0x56
	ETB     = 0x57
	CAN     = 0x58
	EM      = 0x59
	SUB     = 0x5a
	ESC     = 0x5b
	FS      = 0x5c
	GS      = 0x5d
	RS      = 0x5e
	US      = 0x5f
	FNC_3   = 0x60
	FNC_2   = 0x61
	SHIFT_B = 0x62
	CODE_C  = 0x63
	CODE_B  = 0x64
	FNC_4   = 0x65
	FNC_1   = 0x66

	/*
		51 	DC1 121142
		52 	DC2 121241
		53 	DC3 114212
		54 	DC4 124112
		55 	NAK 124211
		56 	SYN 411212
		57 	ETB 421112
		58 	CAN 421211
		59 	EM 	212141
		5a 	SUB 214121
		5b 	ESC 412121
		5c 	FS 	111143
		5d 	GS 	111341
		5e 	RS 	131141
		5f 	US 	114113
		60 	FNC 3 114311
		61 	FNC 2 411113
		62 	Shift B 411311
		63 	Code C 113141
		64 	Code B 114131
		65 	FNC 4 311141
		66 	FNC 1 411131
	*/

	CODE_A = 0x65

	// character set A, B, C
	START_A      = 0x67
	START_B      = 0x68
	START_C      = 0x69
	STOP         = 0x6a
	REVERSE_STOP = -1
)

var DecodeTableA = [][][][][][]int{
	1: {
		1: {
			1: {
				2: {
					2: {
						4: UNDERSCORE,
					},
					4: {
						2: DLE,
					},
				},
				3: {
					2: {
						3: A,
					},
				},
				4: {
					2: {
						2: NUL,
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
					1: {
						4: ENQ,
					},
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
				4: {
					1: {
						2: ACK,
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
				1: {
					2: {
						4: SOH,
					},
				},
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
				4: {
					2: {
						1: STX,
					},
				},
			},
			2: {
				1: {
					1: {
						4: BEL,
					},
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
				4: {
					1: {
						1: BS,
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
			4: {
				1: {
					1: {
						1: SI,
					},
				},
			},
		},
		4: {
			1: {
				2: {
					2: {
						1: EOT,
					},
				},
				1: {
					2: {
						2: ETX,
					},
				},
			},
			2: {
				1: {
					1: {
						2: HT,
					},
				},
				2: {
					1: {
						1: LF,
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
					1: {
						4: FF,
					},
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
				4: {
					1: {
						1: CLOSE_BRACKET,
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
		4: {
			1: {
				1: {
					1: {
						2: SO,
					},
				},
				2: {
					1: {
						1: VT,
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
			4: {
				1: {
					1: {
						1: BACKSLASH,
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
			2: {
				1: {
					1: {
						1: OPEN_BRACKET,
					},
				},
			},
		},
	},
	4: {
		1: {
			3: {
				1: {
					1: {
						1: CR,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						1: CARET,
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
	OPEN_BRACKET:      "[",
	BACKSLASH:         `\`,
	CLOSE_BRACKET:     "]",
	CARET:             "^",
	UNDERSCORE:        "_",
	NUL:               "<NUL>",
	SOH:               "<SOH>",
	STX:               "<STX>",
	ETX:               "<ETX>",
	EOT:               "<EOT>",
	ENQ: "<ENQ>",
	ACK: "<ACK>",
	BEL: "<BEL>",
	BS:  "<BS>",
	HT:  "<HT>",
	LF:  "<LF>",
	VT:  "<VT>",
	FF:  "<FF>",
	CR:  "<CR>",
	SO:  "<SO>",
	SI:  "<SI>",
	DLE: "<DLE>",

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
