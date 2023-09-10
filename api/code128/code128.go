// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"reflect"
)

// ASCII special characters
const (
	NUL = 0x00 // '\0' Null
	SOH = 0x01 //      Start of Header
	STX = 0x02 //      Start of Text
	ETX = 0x03 //      End of Text
	EOT = 0x04 //      End of Transmission
	ENQ = 0x05 //      Enquiry
	ACK = 0x06 //      Acknowledgement
	BEL = 0x07 // '\a' Bell
	BS  = 0x08 // '\b' Backspace
	HT  = 0x09 // '\t' Horizontal Tab
	LF  = 0x0A // '\n' Line Feed
	VT  = 0x0B // '\v' Vertical Tab
	FF  = 0x0C // '\f' Form Feed
	CR  = 0x0D // '\r' Carriage Return
	SO  = 0x0E //      Shift Out
	SI  = 0x0F //      Shift In
	DLE = 0x10 //      Device Idle
	DC1 = 0x11 //      Device Control 1
	DC2 = 0x12 //      Device Control 2
	DC3 = 0x13 //      Device Control 3
	DC4 = 0x14 //      Device Control 4
	NAK = 0x15 //      Negative Acknoledgement
	SYN = 0x16 //      Synchronize
	ETB = 0x17 //      End of Transmission Block
	CAN = 0x18 //      Cancel
	EM  = 0x19 //      End of Medium
	SUB = 0x1A //      Substitute
	ESC = 0x1B // '\e' Escape
	FS  = 0x1C //      Field Separator
	GS  = 0x1D //      Group Separator
	RS  = 0x1E //      Record Separator
	US  = 0x1F //      Unit Separator
	SP  = 0x20 //      Space
	DEL = 0x7F //      Delete
)

// Code128 special characters
const (
	FNC3    = 96 + 32
	FNC2    = 97 + 32
	SHIFT   = 98 + 32
	CODE_C  = 99 + 32
	CODE_B  = 100 + 32
	FNC4_B  = 100 + 32
	CODE_A  = 101 + 32
	FNC4_A  = 101 + 32
	FNC1    = 102 + 32
	START_A = 103 + 32
	START_B = 104 + 32
	START_C = 105 + 32
	STOP    = 106 + 32
)

var bitpattern = [][]int{
	{' ', ' ', 0, 2, 1, 2, 2, 2, 2},
	{'!', '!', 1, 2, 2, 2, 1, 2, 2},
	{'"', '"', 2, 2, 2, 2, 2, 2, 1},
	{'#', '#', 3, 1, 2, 1, 2, 2, 3},
	{'$', '$', 4, 1, 2, 1, 3, 2, 2},
	{'%', '%', 5, 1, 3, 1, 2, 2, 2},
	{'&', '&', 6, 1, 2, 2, 2, 1, 3},
	{'\'', '\'', 7, 1, 2, 2, 3, 1, 2},
	{'(', '(', 8, 1, 3, 2, 2, 1, 2},
	{')', ')', 9, 2, 2, 1, 2, 1, 3},
	{'*', '*', 10, 2, 2, 1, 3, 1, 2},
	{'+', '+', 11, 2, 3, 1, 2, 1, 2},
	{',', ',', 12, 1, 1, 2, 2, 3, 2},
	{'-', '-', 13, 1, 2, 2, 1, 3, 2},
	{'.', '.', 14, 1, 2, 2, 2, 3, 1},
	{'/', '/', 15, 1, 1, 3, 2, 2, 2},
	{'0', '0', 16, 1, 2, 3, 1, 2, 2},
	{'1', '1', 17, 1, 2, 3, 2, 2, 1},
	{'2', '2', 18, 2, 2, 3, 2, 1, 1},
	{'3', '3', 19, 2, 2, 1, 1, 3, 2},
	{'4', '4', 20, 2, 2, 1, 2, 3, 1},
	{'5', '5', 21, 2, 1, 3, 2, 1, 2},
	{'6', '6', 22, 2, 2, 3, 1, 1, 2},
	{'7', '7', 23, 3, 1, 2, 1, 3, 1},
	{'8', '8', 24, 3, 1, 1, 2, 2, 2},
	{'9', '9', 25, 3, 2, 1, 1, 2, 2},
	{':', ':', 26, 3, 2, 1, 2, 2, 1},
	{';', ';', 27, 3, 1, 2, 2, 1, 2},
	{'<', '<', 28, 3, 2, 2, 1, 1, 2},
	{'=', '=', 29, 3, 2, 2, 2, 1, 1},
	{'>', '>', 30, 2, 1, 2, 1, 2, 3},
	{'?', '?', 31, 2, 1, 2, 3, 2, 1},
	{'@', '@', 32, 2, 3, 2, 1, 2, 1},
	{'A', 'A', 33, 1, 1, 1, 3, 2, 3},
	{'B', 'B', 34, 1, 3, 1, 1, 2, 3},
	{'C', 'C', 35, 1, 3, 1, 3, 2, 1},
	{'D', 'D', 36, 1, 1, 2, 3, 1, 3},
	{'E', 'E', 37, 1, 3, 2, 1, 1, 3},
	{'F', 'F', 38, 1, 3, 2, 3, 1, 1},
	{'G', 'G', 39, 2, 1, 1, 3, 1, 3},
	{'H', 'H', 40, 2, 3, 1, 1, 1, 3},
	{'I', 'I', 41, 2, 3, 1, 3, 1, 1},
	{'J', 'J', 42, 1, 1, 2, 1, 3, 3},
	{'K', 'K', 43, 1, 1, 2, 3, 3, 1},
	{'L', 'L', 44, 1, 3, 2, 1, 3, 1},
	{'M', 'M', 45, 1, 1, 3, 1, 2, 3},
	{'N', 'N', 46, 1, 1, 3, 3, 2, 1},
	{'O', 'O', 47, 1, 3, 3, 1, 2, 1},
	{'P', 'P', 48, 3, 1, 3, 1, 2, 1},
	{'Q', 'Q', 49, 2, 1, 1, 3, 3, 1},
	{'R', 'R', 50, 2, 3, 1, 1, 3, 1},
	{'S', 'S', 51, 2, 1, 3, 1, 1, 3},
	{'T', 'T', 52, 2, 1, 3, 3, 1, 1},
	{'U', 'U', 53, 2, 1, 3, 1, 3, 1},
	{'V', 'V', 54, 3, 1, 1, 1, 2, 3},
	{'W', 'W', 55, 3, 1, 1, 3, 2, 1},
	{'X', 'X', 56, 3, 3, 1, 1, 2, 1},
	{'Y', 'Y', 57, 3, 1, 2, 1, 1, 3},
	{'Z', 'Z', 58, 3, 1, 2, 3, 1, 1},
	{'[', '[', 59, 3, 3, 2, 1, 1, 1},
	{'\\', '\\', 60, 3, 1, 4, 1, 1, 1},
	{']', ']', 61, 2, 2, 1, 4, 1, 1},
	{'^', '^', 62, 4, 3, 1, 1, 1, 1},
	{'_', '_', 63, 1, 1, 1, 2, 2, 4},
	{NUL, '`', 64, 1, 1, 1, 4, 2, 2},
	{SOH, 'a', 65, 1, 2, 1, 1, 2, 4},
	{STX, 'b', 66, 1, 2, 1, 4, 2, 1},
	{ETX, 'c', 67, 1, 4, 1, 1, 2, 2},
	{EOT, 'd', 68, 1, 4, 1, 2, 2, 1},
	{ENQ, 'e', 69, 1, 1, 2, 2, 1, 4},
	{ACK, 'f', 70, 1, 1, 2, 4, 1, 2},
	{BEL, 'g', 71, 1, 2, 2, 1, 1, 4},
	{BS, 'h', 72, 1, 2, 2, 4, 1, 1},
	{HT, 'i', 73, 1, 4, 2, 1, 1, 2},
	{LF, 'j', 74, 1, 4, 2, 2, 1, 1},
	{VT, 'k', 75, 2, 4, 1, 2, 1, 1},
	{FF, 'l', 76, 2, 2, 1, 1, 1, 4},
	{CR, 'm', 77, 4, 1, 3, 1, 1, 1},
	{SO, 'n', 78, 2, 4, 1, 1, 1, 2},
	{SI, 'o', 79, 1, 3, 4, 1, 1, 1},
	{DLE, 'p', 80, 1, 1, 1, 2, 4, 2},
	{DC1, 'q', 81, 1, 2, 1, 1, 4, 2},
	{DC2, 'r', 82, 1, 2, 1, 2, 4, 1},
	{DC3, 's', 83, 1, 1, 4, 2, 1, 2},
	{DC4, 't', 84, 1, 2, 4, 1, 1, 2},
	{NAK, 'u', 85, 1, 2, 4, 2, 1, 1},
	{SYN, 'v', 86, 4, 1, 1, 2, 1, 2},
	{ETB, 'w', 87, 4, 2, 1, 1, 1, 2},
	{CAN, 'x', 88, 4, 2, 1, 2, 1, 1},
	{EM, 'y', 89, 2, 1, 2, 1, 4, 1},
	{SUB, 'z', 90, 2, 1, 4, 1, 2, 1},
	{ESC, '{', 91, 4, 1, 2, 1, 2, 1},
	{FS, '|', 92, 1, 1, 1, 1, 4, 3},
	{GS, '}', 93, 1, 1, 1, 3, 4, 1},
	{RS, '~', 94, 1, 3, 1, 1, 4, 1},
	{US, DEL, 95, 1, 1, 4, 1, 1, 3},
	{FNC3, FNC3, 96, 1, 1, 4, 3, 1, 1},
	{FNC2, FNC2, 97, 4, 1, 1, 1, 1, 3},
	{SHIFT, SHIFT, 98, 4, 1, 1, 3, 1, 1},
	{CODE_C, CODE_C, 99, 1, 1, 3, 1, 4, 1},
	{CODE_B, FNC4_B, CODE_B, 1, 1, 4, 1, 3, 1},
	{FNC4_A, CODE_A, CODE_A, 3, 1, 1, 1, 4, 1},
	{FNC1, FNC1, FNC1, 4, 1, 1, 1, 3, 1},
	{START_A, START_A, START_A, 2, 1, 1, 4, 1, 2},
	{START_B, START_B, START_B, 2, 1, 1, 2, 1, 4},
	{START_C, START_C, START_C, 2, 1, 1, 2, 3, 2},
}

func Widths(img image.Image) (widths []int, err error) {
	bars := false
	run := 0
	div := 1
	divFound := false
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
				widths = append(widths, run / div)

				bars = true
				run = 1
			}
		} else if r == 0xFFFF {
			if !bars {
				run++
			} else {
				// finish bar run
				widths = append(widths, run / div)

				bars = false
				run = 1
			}
		}
	}
	// don't forget to record last run!
	widths = append(widths, run / div)
	return widths, nil
}

func Analyse(widths []int) (quietStart int, startSym []int, data []int, checkSym []int, stopPat []int, quietEnd int) {
	quietStart = widths[0]
	startSym = widths[1:7]
	data = widths[7:len(widths)-14]
	checkSym = widths[len(widths)-14:len(widths)-8]
	stopPat = widths[len(widths)-8:len(widths)-1]
	quietEnd = widths[len(widths)-1]
	return
}

var DecodeTableA = [][][][][][]string{
	2: {
		1: {
			1: {
				1: {
					3: {
						3: "REVERSE_STOP",
					},
				},
				2: {
					1: {
						4: "START_B",
					},
					3: {
						2: "START_C",
					},
				},
				4: {
					1: {
						2: "START_A",
					},
				},
			},
		},
		3: {
			3: {
				1: {
					1: {
						1: "STOP",
					},
				},
			},
		},
	},
	3: {
		1: {
			1: {
				3: {
					2: {
						1: "W",
					},
				},
			},
		},
	},
}

var DecodeTableB = [][][][][][]string{
	1: {
		1: {
			1: {
				3: {
					2: {
						3: "A",
					},
				},
			},
			2: {
				1: {
					3: {
						3: "J",
					},
				},
				3: {
					1: {
						3: "D",
					},
				},
			},
			3: {
				1: {
					4: {
						1: "CODE_C",
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						4: "a",
					},
				},
				4: {
					2: {
						1: "b",
					},
				},
			},
			2: {
				1: {
					3: {
						2: "-",
					},
				},
			},
			3: {
				2: {
					2: {
						1: "1",
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: "B",
					},
				},
				3: {
					2: {
						1: "C",
					},
				},
			},
		},
		4: {
			1: {
				1: {
					2: {
						2: "c",
					},
				},
				2: {
					2: {
						1: "d",
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
						3: "REVERSE_STOP",
					},
				},
				2: {
					1: {
						4: "START_B",
					},
					3: {
						2: "START_C",
					},
				},
				4: {
					1: {
						2: "START_A",
					},
				},
			},
		},
		2: {
			1: {
				1: {
					3: {
						2: "3",
					},
				},
				4: {
					1: {
						1: "]",
					},
				},
			},
			3: {
				2: {
					1: {
						1: "2",
					},
				},
			},
		},
		3: {
			3: {
				1: {
					1: {
						1: "STOP",
					},
				},
			},
		},
	},
	3: {
		1: {
			1: {
				3: {
					2: {
						1: "W",
					},
				},
			},
			3: {
				1: {
					2: {
						1: "P",
					},
				},
			},
		},
	},
}

var DecodeTableC = [][][][][][]string{
	1: {
		1: {
			2: {
				2: {
					3: {
						2: "12",
					},
				},
			},
			4: {
				1: {
					3: {
						1: "CODE_B",
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: "34",
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
						3: "REVERSE_STOP",
					},
				},
				2: {
					1: {
						4: "START_B",
					},
					3: {
						2: "START_C",
					},
				},
				4: {
					1: {
						2: "START_A",
					},
				},
			},
		},
		3: {
			3: {
				1: {
					1: {
						1: "STOP",
					},
				},
			},
		},
	},
}

var ValueTable = map[string]int{
	"12": 12,
	"34": 13,
	"-": 13,
	"1": 17,
	"2": 18,
	"3": 19,
	"A": 33,
	"B": 34,
	"C": 35,
	"D": 36,
	"P": 48,
	"J": 42,
	"W": 55,
	"a": 65,
	"b": 66,
	"c": 67,
	"d": 68,
	"CODE_C": 99,
	"CODE_B": 100,
	"START_B": 104,
}

func Decode(img image.Image) (msg string, err error) {
	widths, err := Widths(img)
	if err != nil {
		return "", err
	}
	fmt.Printf("%+v\n", widths)
	qs, sta, d, c, stp, qe := Analyse(widths)
	fmt.Printf("qs: %d\nsta: %+v\nd: %+v\nc: %+v\nstp: %+v\nqe: %d\n", qs, sta, d, c, stp, qe)

	decodeTable := DecodeTableA

	staSym := decodeTable[sta[0]][sta[1]][sta[2]][sta[3]][sta[4]][sta[5]]
	switch staSym {
	case "START_A":
		decodeTable = DecodeTableA
	case "START_B":
		decodeTable = DecodeTableB
	case "START_C":
		decodeTable = DecodeTableC
	default:
		// @todo: handle case where we're reading barcode in reverse (right-to-left)
		return msg, fmt.Errorf("invalid start symbol: %s -- %+v", staSym, sta)
	}

	checksum := ValueTable[staSym]
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
			fmt.Printf("(Table %s) Unable to parse sequence:", table)
			for i := -5; current+i < len(d) && i <= 0; i++ {
				fmt.Printf(" %d", d[current+i])
			}
			fmt.Println()
			fmt.Printf("More: %+v\n", d[current+1:])
		}
	}()

	for current < len(d) {
		sym := decodeTable[d[current-5]][d[current-4]][d[current-3]][d[current-2]][d[current-1]][d[current-0]]
		val := ValueTable[sym]
		checksum += val*posMul
		fmt.Printf("Sym: %s (%d%d%d%d%d%d) [%d × %d = %d]\n", sym, d[current-5], d[current-4], d[current-3], d[current-2], d[current-1], d[current-0], val, posMul, val*posMul)

		switch sym {
		case "CODE_A":
			decodeTable = DecodeTableA
		case "CODE_B":
			decodeTable = DecodeTableB
		case "CODE_C":
			decodeTable = DecodeTableC
		default:
			msg += sym
		}

		posMul++
		current += 6
	}

	checksum = checksum % 103
	cksmVal := ValueTable[DecodeTableA[c[0]][c[1]][c[2]][c[3]][c[4]][c[5]]]
	cksmOK := cksmVal == checksum
	fmt.Printf("Checksum: %d (expected: %d, ok: %t)\n", checksum, cksmVal, cksmOK)

	if !cksmOK {
		return msg, fmt.Errorf("invalid checksum: want: %d, got: %d", cksmVal, checksum)
	}

	return msg, nil
}

/*
func Decode(img image.Image) (string, error) {
	set := false
	count := 0
	fmt.Printf("Width: %d\n", img.Bounds().Dx())
	for x := 0; x < img.Bounds().Dx(); x++ {
		c := img.At(x, 0)

		r, _, _, _ := c.RGBA()
		if r == 0x0 {
			if set {
				count++
			} else {
				fmt.Printf("0 × %d\n", count)
				count = 1
				set = true
			}
		} else if r == 0xFFFF {
			if !set {
				count++
			} else {
				fmt.Printf("1 × %d\n", count)
				count = 1
				set = false
			}
		} else {
			panic("i don't know how to handle this")
		}
	}
	if set {
		fmt.Printf("1 × %d\n", count)
	} else {
		fmt.Printf("0 × %d\n", count)
	}
	return "", errors.New("Not implemented")
}
*/

func Encode(text string) (image.Image, error) {
	return nil, errors.New("Not implemented")
}
