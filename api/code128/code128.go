// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
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
	VT  = 0x0B // '\v' Verical Tab
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

// Decode code128 barcode
func Decode(img image.Image) (string, error) {
	for x := 0; x < img.Bounds().Max.X; x++ {
		c := img.At(x, 0)
		//fmt.Printf("c: %v\n", c)
		r, _, _, _ := c.RGBA()
		if r == 0x0 {
			fmt.Println(0)
		} else if r == 0xFFFF {
			fmt.Println(1)
		}
	}
	return "", errors.New("Not implemented")
}

// Encode text to code128
func Encode(text string) (image.Image, error) {
	return nil, errors.New("Not implemented")
}
