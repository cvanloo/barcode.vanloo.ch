package code128

/*
 * Encoding Tables
 */

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

// SpecialOffset is added to the value of Code-128 special symbols, to offset
// them from the ASCII table.
const SpecialOffset = 32

// Code128 special characters
// The special characters are offset from the ASCII table by SpecialOffset.
const (
	FNC3    = 96 + SpecialOffset
	FNC2    = 97 + SpecialOffset
	SHIFT   = 98 + SpecialOffset
	CODE_C  = 99 + SpecialOffset
	CODE_B  = 100 + SpecialOffset
	FNC4_B  = 100 + SpecialOffset
	CODE_A  = 101 + SpecialOffset
	FNC4_A  = 101 + SpecialOffset
	FNC1    = 102 + SpecialOffset
	START_A = 103 + SpecialOffset
	START_B = 104 + SpecialOffset
	START_C = 105 + SpecialOffset
	STOP    = 106 + SpecialOffset
)

// Bitpattern of the Code-128 symbols
// A bitpattern is laid out as the array {A, B, C, M1-M6}, where A, B, C
// designate the symbol according to its corresponding character set, and M1
// through M6 are the module widths.
// The index into the array coincides with the symbol's Code-128 value
// (necessary for calculating the checksum).
// For special, non-ascii symbols, such as START, CODE, and FNC, the value is
// offset by 32 from the index, as to not overlap with ASCII.
var Bitpattern = [][]int{
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

type TableIndex int
const (
	LookupNone TableIndex = -1
	LookupA TableIndex = 0
	LookupB TableIndex = 1
	LookupC TableIndex = 2
)

/*
 * Decoding Tables
 */

const (
	SYM_SPACE, NUM_00                = 0x00, 0x00
	SYM_EXCLAMATION, NUM_01          = 0x01, 0x01
	SYM_DOUBLE_QUOTE, NUM_02         = 0x02, 0x02
	SYM_POUND, NUM_03                = 0x03, 0x03
	SYM_DOLLAR, NUM_04               = 0x04, 0x04
	SYM_PERCENT, NUM_05              = 0x05, 0x05
	SYM_AMPERSAND, NUM_06            = 0x06, 0x06
	SYM_SINGLE_QUOTE, NUM_07         = 0x07, 0x07
	SYM_OPEN_PARENTHESIS, NUM_08     = 0x08, 0x08
	SYM_CLOSE_PARENTHESIS, NUM_09    = 0x09, 0x09
	SYM_ASTERISK, NUM_10             = 0x0a, 0x0a
	SYM_PLUS, NUM_11                 = 0x0b, 0x0b
	SYM_COMMA, NUM_12                = 0x0c, 0x0c
	SYM_HYPHEN, NUM_13               = 0x0d, 0x0d
	SYM_PERIOD, NUM_14               = 0x0e, 0x0e
	SYM_SLASH, NUM_15                = 0x0f, 0x0f
	SYM_ZERO, NUM_16                 = 0x10, 0x10
	SYM_ONE, NUM_17                  = 0x11, 0x11
	SYM_TWO, NUM_18                  = 0x12, 0x12
	SYM_THREE, NUM_19                = 0x13, 0x13
	SYM_FOUR, NUM_20                 = 0x14, 0x14
	SYM_FIVE, NUM_21                 = 0x15, 0x15
	SYM_SIX, NUM_22                  = 0x16, 0x16
	SYM_SEVEN, NUM_23                = 0x17, 0x17
	SYM_EIGHT, NUM_24                = 0x18, 0x18
	SYM_NINE, NUM_25                 = 0x19, 0x19
	SYM_COLON, NUM_26                = 0x1a, 0x1a
	SYM_SEMICOLON, NUM_27            = 0x1b, 0x1b
	SYM_LESS_THAN, NUM_28            = 0x1c, 0x1c
	SYM_EQUAL, NUM_29                = 0x1d, 0x1d
	SYM_GREATER_THAN, NUM_30         = 0x1e, 0x1e
	SYM_QUESTION, NUM_31             = 0x1f, 0x1f
	SYM_AT, NUM_32                   = 0x20, 0x20
	SYM_A, NUM_33                    = 0x21, 0x21
	SYM_B, NUM_34                    = 0x22, 0x22
	SYM_C, NUM_35                    = 0x23, 0x23
	SYM_D, NUM_36                    = 0x24, 0x24
	SYM_E, NUM_37                    = 0x25, 0x25
	SYM_F, NUM_38                    = 0x26, 0x26
	SYM_G, NUM_39                    = 0x27, 0x27
	SYM_H, NUM_40                    = 0x28, 0x28
	SYM_I, NUM_41                    = 0x29, 0x29
	SYM_J, NUM_42                    = 0x2a, 0x2a
	SYM_K, NUM_43                    = 0x2b, 0x2b
	SYM_L, NUM_44                    = 0x2c, 0x2c
	SYM_M, NUM_45                    = 0x2d, 0x2d
	SYM_N, NUM_46                    = 0x2e, 0x2e
	SYM_O, NUM_47                    = 0x2f, 0x2f
	SYM_P, NUM_48                    = 0x30, 0x30
	SYM_Q, NUM_49                    = 0x31, 0x31
	SYM_R, NUM_50                    = 0x32, 0x32
	SYM_S, NUM_51                    = 0x33, 0x33
	SYM_T, NUM_52                    = 0x34, 0x34
	SYM_U, NUM_53                    = 0x35, 0x35
	SYM_V, NUM_54                    = 0x36, 0x36
	SYM_W, NUM_55                    = 0x37, 0x37
	SYM_X, NUM_56                    = 0x38, 0x38
	SYM_Y, NUM_57                    = 0x39, 0x39
	SYM_Z, NUM_58                    = 0x3a, 0x3a
	SYM_OPEN_BRACKET, NUM_59         = 0x3b, 0x3b
	SYM_BACKSLASH, NUM_60            = 0x3c, 0x3c
	SYM_CLOSE_BRACKET, NUM_61        = 0x3d, 0x3d
	SYM_CARET, NUM_62                = 0x3e, 0x3e
	SYM_UNDERSCORE, NUM_63           = 0x3f, 0x3f
	SYM_NUL, SYM_BACKTICK, NUM_64    = 0x40, 0x40, 0x40
	SYM_SOH, SYM_a, NUM_65           = 0x41, 0x41, 0x41
	SYM_STX, SYM_b, NUM_66           = 0x42, 0x42, 0x42
	SYM_ETX, SYM_c, NUM_67           = 0x43, 0x43, 0x43
	SYM_EOT, SYM_d, NUM_68           = 0x44, 0x44, 0x44
	SYM_ENQ, SYM_e, NUM_69           = 0x45, 0x45, 0x45
	SYM_ACK, SYM_f, NUM_70           = 0x46, 0x46, 0x46
	SYM_BEL, SYM_g, NUM_71           = 0x47, 0x47, 0x47
	SYM_BS, SYM_h, NUM_72            = 0x48, 0x48, 0x48
	SYM_HT, SYM_i, NUM_73            = 0x49, 0x49, 0x49
	SYM_LF, SYM_j, NUM_74            = 0x4a, 0x4a, 0x4a
	SYM_VT, SYM_k, NUM_75            = 0x4b, 0x4b, 0x4b
	SYM_FF, SYM_l, NUM_76            = 0x4c, 0x4c, 0x4c
	SYM_CR, SYM_m, NUM_77            = 0x4d, 0x4d, 0x4d
	SYM_SO, SYM_n, NUM_78            = 0x4e, 0x4e, 0x4e
	SYM_SI, SYM_o, NUM_79            = 0x4f, 0x4f, 0x4f
	SYM_DLE, SYM_p, NUM_80           = 0x50, 0x50, 0x50
	SYM_DC1, SYM_q, NUM_81           = 0x51, 0x51, 0x51
	SYM_DC2, SYM_r, NUM_82           = 0x52, 0x52, 0x52
	SYM_DC3, SYM_s, NUM_83           = 0x53, 0x53, 0x53
	SYM_DC4, SYM_t, NUM_84           = 0x54, 0x54, 0x54
	SYM_NAK, SYM_u, NUM_85           = 0x55, 0x55, 0x55
	SYM_SYN, SYM_v, NUM_86           = 0x56, 0x56, 0x56
	SYM_ETB, SYM_w, NUM_87           = 0x57, 0x57, 0x57
	SYM_CAN, SYM_x, NUM_88           = 0x58, 0x58, 0x58
	SYM_EM, SYM_y, NUM_89            = 0x59, 0x59, 0x59
	SYM_SUB, SYM_z, NUM_90           = 0x5a, 0x5a, 0x5a
	SYM_ESC, SYM_OPEN_BRACE, NUM_91  = 0x5b, 0x5b, 0x5b
	SYM_FS, SYM_PIPE, NUM_92         = 0x5c, 0x5c, 0x5c
	SYM_GS, SYM_CLOSE_BRACE, NUM_93  = 0x5d, 0x5d, 0x5d
	SYM_RS, SYM_TILDE, NUM_94        = 0x5e, 0x5e, 0x5e
	SYM_US, SYM_DEL, NUM_95          = 0x5f, 0x5f, 0x5f
	SYM_FNC3, NUM_96                 = 0x60, 0x60
	SYM_FNC2, NUM_97                 = 0x61, 0x61
	SYM_SHIFT_B, SYM_SHIFT_A, NUM_98 = 0x62, 0x62, 0x62
	SYM_CODE_C, NUM_99               = 0x63, 0x63
	SYM_CODE_B, SYM_FNC4_B           = 0x64, 0x64
	SYM_FNC4_A, SYM_CODE_A           = 0x65, 0x65
	SYM_FNC1                         = 0x66
	SYM_START_A                      = 0x67
	SYM_START_B                      = 0x68
	SYM_START_C                      = 0x69
	SYM_STOP                         = 0x6a
	SYM_REVERSE_STOP                 = -1
)

var DecodeTableA = [][][][][][]int{
	1: {
		1: {
			1: {
				1: {
					4: {
						3: SYM_FS,
					},
				},
				2: {
					2: {
						4: SYM_UNDERSCORE,
					},
					4: {
						2: SYM_DLE,
					},
				},
				3: {
					2: {
						3: SYM_A,
					},
					4: {
						1: SYM_GS,
					},
				},
				4: {
					2: {
						2: SYM_NUL,
					},
				},
			},
			2: {
				1: {
					3: {
						3: SYM_J,
					},
				},
				2: {
					1: {
						4: SYM_ENQ,
					},
					3: {
						2: SYM_COMMA,
					},
				},
				3: {
					1: {
						3: SYM_D,
					},
					3: {
						1: SYM_K,
					},
				},
				4: {
					1: {
						2: SYM_ACK,
					},
				},
			},
			3: {
				1: {
					2: {
						3: SYM_M,
					},
					4: {
						1: SYM_CODE_C,
					},
				},
				2: {
					2: {
						2: SYM_SLASH,
					},
				},
				3: {
					2: {
						1: SYM_N,
					},
				},
			},
			4: {
				1: {
					1: {
						3: SYM_US,
					},
					3: {
						1: SYM_CODE_B,
					},
				},
				2: {
					1: {
						2: SYM_DC3,
					},
				},
				3: {
					1: {
						1: SYM_FNC3,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						4: SYM_SOH,
					},
					4: {
						2: SYM_DC1,
					},
				},
				2: {
					2: {
						3: SYM_POUND,
					},
					4: {
						1: SYM_DC2,
					},
				},
				3: {
					2: {
						2: SYM_DOLLAR,
					},
				},
				4: {
					2: {
						1: SYM_STX,
					},
				},
			},
			2: {
				1: {
					1: {
						4: SYM_BEL,
					},
					3: {
						2: SYM_HYPHEN,
					},
				},
				2: {
					1: {
						3: SYM_AMPERSAND,
					},
					3: {
						1: SYM_PERIOD,
					},
				},
				3: {
					1: {
						2: SYM_SINGLE_QUOTE,
					},
				},
				4: {
					1: {
						1: SYM_BS,
					},
				},
			},
			3: {
				1: {
					2: {
						2: SYM_ZERO,
					},
				},
				2: {
					2: {
						1: SYM_ONE,
					},
				},
			},
			4: {
				1: {
					1: {
						2: SYM_DC4,
					},
				},
				2: {
					1: {
						1: SYM_NAK,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: SYM_B,
					},
					4: {
						1: SYM_RS,
					},
				},
				2: {
					2: {
						2: SYM_PERCENT,
					},
				},
				3: {
					2: {
						1: SYM_C,
					},
				},
			},
			2: {
				1: {
					1: {
						3: SYM_E,
					},
					3: {
						1: SYM_L,
					},
				},
				2: {
					1: {
						2: SYM_OPEN_PARENTHESIS,
					},
				},
				3: {
					1: {
						1: SYM_F,
					},
				},
			},
			3: {
				1: {
					2: {
						1: SYM_O,
					},
				},
			},
			4: {
				1: {
					1: {
						1: SYM_SI,
					},
				},
			},
		},
		4: {
			1: {
				2: {
					2: {
						1: SYM_EOT,
					},
				},
				1: {
					2: {
						2: SYM_ETX,
					},
				},
			},
			2: {
				1: {
					1: {
						2: SYM_HT,
					},
				},
				2: {
					1: {
						1: SYM_LF,
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
						3: SYM_REVERSE_STOP,
					},
				},
				2: {
					1: {
						4: SYM_START_B,
					},
					3: {
						2: SYM_START_C,
					},
				},
				3: {
					1: {
						3: SYM_G,
					},
					3: {
						1: SYM_Q,
					},
				},
				4: {
					1: {
						2: SYM_START_A,
					},
				},
			},
			2: {
				1: {
					2: {
						3: SYM_GREATER_THAN,
					},
					4: {
						1: SYM_EM,
					},
				},
				2: {
					2: {
						2: SYM_SPACE,
					},
				},
				3: {
					2: {
						1: SYM_QUESTION,
					},
				},
			},
			3: {
				1: {
					1: {
						3: SYM_S,
					},
					3: {
						1: SYM_U,
					},
				},
				2: {
					1: {
						2: SYM_FIVE,
					},
				},
				3: {
					1: {
						1: SYM_T,
					},
				},
			},
			4: {
				1: {
					2: {
						1: SYM_SUB,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						4: SYM_FF,
					},
					3: {
						2: SYM_THREE,
					},
				},
				2: {
					1: {
						3: SYM_CLOSE_PARENTHESIS,
					},
					3: {
						1: SYM_FOUR,
					},
				},
				3: {
					1: {
						2: SYM_ASTERISK,
					},
				},
				4: {
					1: {
						1: SYM_CLOSE_BRACKET,
					},
				},
			},
			2: {
				1: {
					2: {
						2: SYM_EXCLAMATION,
					},
				},
				2: {
					2: {
						1: SYM_DOUBLE_QUOTE,
					},
				},
			},
			3: {
				1: {
					1: {
						2: SYM_SIX,
					},
				},
				2: {
					1: {
						1: SYM_TWO,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						3: SYM_H,
					},
					3: {
						1: SYM_R,
					},
				},
				2: {
					1: {
						2: SYM_PLUS,
					},
				},
				3: {
					1: {
						1: SYM_I,
					},
				},
			},
			2: {
				1: {
					2: {
						1: SYM_AT,
					},
				},
			},
			3: {
				1: {
					1: {
						1: SYM_STOP,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					1: {
						2: SYM_SO,
					},
				},
				2: {
					1: {
						1: SYM_VT,
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
						3: SYM_V,
					},
					4: {
						1: SYM_FNC4_A,
					},
				},
				2: {
					2: {
						2: SYM_EIGHT,
					},
				},
				3: {
					2: {
						1: SYM_W,
					},
				},
			},
			2: {
				1: {
					1: {
						3: SYM_Y,
					},
					3: {
						1: SYM_SEVEN,
					},
				},
				2: {
					1: {
						2: SYM_SEMICOLON,
					},
				},
				3: {
					1: {
						1: SYM_Z,
					},
				},
			},
			3: {
				1: {
					2: {
						1: SYM_P,
					},
				},
			},
			4: {
				1: {
					1: {
						1: SYM_BACKSLASH,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						2: SYM_NINE,
					},
				},
				2: {
					2: {
						1: SYM_COLON,
					},
				},
			},
			2: {
				1: {
					1: {
						2: SYM_LESS_THAN,
					},
				},
				2: {
					1: {
						1: SYM_EQUAL,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						1: SYM_X,
					},
				},
			},
			2: {
				1: {
					1: {
						1: SYM_OPEN_BRACKET,
					},
				},
			},
		},
	},
	4: {
		1: {
			1: {
				1: {
					1: {
						3: SYM_FNC2,
					},
					3: {
						1: SYM_FNC1,
					},
				},
				2: {
					1: {
						2: SYM_SYN,
					},
				},
				3: {
					1: {
						1: SYM_SHIFT_B,
					},
				},
			},
			2: {
				1: {
					2: {
						1: SYM_ESC,
					},
				},
			},
			3: {
				1: {
					1: {
						1: SYM_CR,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						2: SYM_ETB,
					},
				},
				2: {
					1: {
						1: SYM_CAN,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						1: SYM_CARET,
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
				1: {
					4: {
						3: SYM_PIPE,
					},
				},
				2: {
					2: {
						4: SYM_UNDERSCORE,
					},
					4: {
						2: SYM_p,
					},
				},
				3: {
					2: {
						3: SYM_A,
					},
					4: {
						1: SYM_CLOSE_BRACE,
					},
				},
				4: {
					2: {
						2: SYM_BACKTICK,
					},
				},
			},
			2: {
				1: {
					3: {
						3: SYM_J,
					},
				},
				2: {
					1: {
						4: SYM_e,
					},
					3: {
						2: SYM_COMMA,
					},
				},
				3: {
					1: {
						3: SYM_D,
					},
					3: {
						1: SYM_K,
					},
				},
				4: {
					1: {
						2: SYM_f,
					},
				},
			},
			3: {
				1: {
					2: {
						3: SYM_M,
					},
					4: {
						1: SYM_CODE_C,
					},
				},
				2: {
					2: {
						2: SYM_SLASH,
					},
				},
				3: {
					2: {
						1: SYM_N,
					},
				},
			},
			4: {
				1: {
					1: {
						3: SYM_DEL,
					},
					3: {
						1: SYM_FNC4_B,
					},
				},
				2: {
					1: {
						2: SYM_s,
					},
				},
				3: {
					1: {
						1: SYM_FNC3,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						4: SYM_a,
					},
					4: {
						2: SYM_q,
					},
				},
				2: {
					2: {
						3: SYM_POUND,
					},
					4: {
						1: SYM_r,
					},
				},
				3: {
					2: {
						2: SYM_DOLLAR,
					},
				},
				4: {
					2: {
						1: SYM_b,
					},
				},
			},
			2: {
				1: {
					1: {
						4: SYM_g,
					},
					3: {
						2: SYM_HYPHEN,
					},
				},
				2: {
					1: {
						3: SYM_AMPERSAND,
					},
					3: {
						1: SYM_PERIOD,
					},
				},
				3: {
					1: {
						2: SYM_SINGLE_QUOTE,
					},
				},
				4: {
					1: {
						1: SYM_h,
					},
				},
			},
			3: {
				1: {
					2: {
						2: SYM_ZERO,
					},
				},
				2: {
					2: {
						1: SYM_ONE,
					},
				},
			},
			4: {
				1: {
					1: {
						2: SYM_t,
					},
				},
				2: {
					1: {
						1: SYM_u,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: SYM_B,
					},
					4: {
						1: SYM_TILDE,
					},
				},
				2: {
					2: {
						2: SYM_PERCENT,
					},
				},
				3: {
					2: {
						1: SYM_C,
					},
				},
			},
			2: {
				1: {
					1: {
						3: SYM_E,
					},
					3: {
						1: SYM_L,
					},
				},
				2: {
					1: {
						2: SYM_OPEN_PARENTHESIS,
					},
				},
				3: {
					1: {
						1: SYM_F,
					},
				},
			},
			3: {
				1: {
					2: {
						1: SYM_O,
					},
				},
			},
			4: {
				1: {
					1: {
						1: SYM_o,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					2: {
						2: SYM_c,
					},
				},
				2: {
					2: {
						1: SYM_d,
					},
				},
			},
			2: {
				1: {
					1: {
						2: SYM_i,
					},
				},
				2: {
					1: {
						1: SYM_j,
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
						3: SYM_REVERSE_STOP,
					},
				},
				2: {
					1: {
						4: SYM_START_B,
					},
					3: {
						2: SYM_START_C,
					},
				},
				3: {
					1: {
						3: SYM_G,
					},
					3: {
						1: SYM_Q,
					},
				},
				4: {
					1: {
						2: SYM_START_A,
					},
				},
			},
			2: {
				1: {
					2: {
						3: SYM_GREATER_THAN,
					},
					4: {
						1: SYM_y,
					},
				},
				2: {
					2: {
						2: SYM_SPACE,
					},
				},
				3: {
					2: {
						1: SYM_QUESTION,
					},
				},
			},
			3: {
				1: {
					1: {
						3: SYM_S,
					},
					3: {
						1: SYM_U,
					},
				},
				2: {
					1: {
						2: SYM_FIVE,
					},
				},
				3: {
					1: {
						1: SYM_T,
					},
				},
			},
			4: {
				1: {
					2: {
						1: SYM_z,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						4: SYM_l,
					},
					3: {
						2: SYM_THREE,
					},
				},
				2: {
					1: {
						3: SYM_CLOSE_PARENTHESIS,
					},
					3: {
						1: SYM_FOUR,
					},
				},
				3: {
					1: {
						2: SYM_ASTERISK,
					},
				},
				4: {
					1: {
						1: SYM_CLOSE_BRACKET,
					},
				},
			},
			2: {
				1: {
					2: {
						2: SYM_EXCLAMATION,
					},
				},
				2: {
					2: {
						1: SYM_DOUBLE_QUOTE,
					},
				},
			},
			3: {
				1: {
					1: {
						2: SYM_SIX,
					},
				},
				2: {
					1: {
						1: SYM_TWO,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						3: SYM_H,
					},
					3: {
						1: SYM_R,
					},
				},
				2: {
					1: {
						2: SYM_PLUS,
					},
				},
				3: {
					1: {
						1: SYM_I,
					},
				},
			},
			2: {
				1: {
					2: {
						1: SYM_AT,
					},
				},
			},
			3: {
				1: {
					1: {
						1: SYM_STOP,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					1: {
						2: SYM_n,
					},
				},
				2: {
					1: {
						1: SYM_k,
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
						3: SYM_V,
					},
					4: {
						1: SYM_CODE_A,
					},
				},
				2: {
					2: {
						2: SYM_EIGHT,
					},
				},
				3: {
					2: {
						1: SYM_W,
					},
				},
			},
			2: {
				1: {
					1: {
						3: SYM_Y,
					},
					3: {
						1: SYM_SEVEN,
					},
				},
				2: {
					1: {
						2: SYM_SEMICOLON,
					},
				},
				3: {
					1: {
						1: SYM_Z,
					},
				},
			},
			3: {
				1: {
					2: {
						1: SYM_P,
					},
				},
			},
			4: {
				1: {
					1: {
						1: SYM_BACKSLASH,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						2: SYM_NINE,
					},
				},
				2: {
					2: {
						1: SYM_COLON,
					},
				},
			},
			2: {
				1: {
					1: {
						2: SYM_LESS_THAN,
					},
				},
				2: {
					1: {
						1: SYM_EQUAL,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						1: SYM_X,
					},
				},
			},
			2: {
				1: {
					1: {
						1: SYM_OPEN_BRACKET,
					},
				},
			},
		},
	},
	4: {
		1: {
			1: {
				1: {
					1: {
						3: SYM_FNC2,
					},
					3: {
						1: SYM_FNC1,
					},
				},
				2: {
					1: {
						2: SYM_v,
					},
				},
				3: {
					1: {
						1: SYM_SHIFT_A,
					},
				},
			},
			2: {
				1: {
					2: {
						1: SYM_OPEN_BRACE,
					},
				},
			},
			3: {
				1: {
					1: {
						1: SYM_m,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						2: SYM_w,
					},
				},
				2: {
					1: {
						1: SYM_x,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						1: SYM_CARET,
					},
				},
			},
		},
	},
}

var DecodeTableC = [][][][][][]int{
	1: {
		1: {
			1: {
				1: {
					4: {
						3: NUM_92,
					},
				},
				2: {
					2: {
						4: NUM_63,
					},
					4: {
						2: NUM_80,
					},
				},
				3: {
					2: {
						3: NUM_33,
					},
					4: {
						1: NUM_93,
					},
				},
				4: {
					2: {
						2: NUM_64,
					},
				},
			},
			2: {
				1: {
					3: {
						3: NUM_42,
					},
				},
				2: {
					1: {
						4: NUM_69,
					},
					3: {
						2: NUM_12,
					},
				},
				3: {
					1: {
						3: NUM_36,
					},
					3: {
						1: NUM_43,
					},
				},
				4: {
					1: {
						2: NUM_70,
					},
				},
			},
			3: {
				1: {
					2: {
						3: NUM_45,
					},
					4: {
						1: NUM_99,
					},
				},
				2: {
					2: {
						2: NUM_15,
					},
				},
				3: {
					2: {
						1: NUM_46,
					},
				},
			},
			4: {
				1: {
					1: {
						3: NUM_95,
					},
					3: {
						1: SYM_CODE_B,
					},
				},
				2: {
					1: {
						2: NUM_83,
					},
				},
				3: {
					1: {
						1: NUM_96,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						4: NUM_65,
					},
					4: {
						2: NUM_81,
					},
				},
				2: {
					2: {
						3: NUM_03,
					},
					4: {
						1: NUM_82,
					},
				},
				3: {
					2: {
						2: NUM_04,
					},
				},
				4: {
					2: {
						1: NUM_66,
					},
				},
			},
			2: {
				1: {
					1: {
						4: NUM_71,
					},
					3: {
						2: NUM_13,
					},
				},
				2: {
					1: {
						3: NUM_06,
					},
					3: {
						1: NUM_14,
					},
				},
				3: {
					1: {
						2: NUM_07,
					},
				},
				4: {
					1: {
						1: NUM_72,
					},
				},
			},
			3: {
				1: {
					2: {
						2: NUM_16,
					},
				},
				2: {
					2: {
						1: NUM_17,
					},
				},
			},
			4: {
				1: {
					1: {
						2: NUM_84,
					},
				},
				2: {
					1: {
						1: NUM_85,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						3: NUM_34,
					},
					4: {
						1: NUM_94,
					},
				},
				2: {
					2: {
						2: NUM_05,
					},
				},
				3: {
					2: {
						1: NUM_35,
					},
				},
			},
			2: {
				1: {
					1: {
						3: NUM_37,
					},
					3: {
						1: NUM_44,
					},
				},
				2: {
					1: {
						2: NUM_08,
					},
				},
				3: {
					1: {
						1: NUM_38,
					},
				},
			},
			3: {
				1: {
					2: {
						1: NUM_47,
					},
				},
			},
			4: {
				1: {
					1: {
						1: NUM_79,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					2: {
						2: NUM_67,
					},
				},
				2: {
					2: {
						1: NUM_68,
					},
				},
			},
			2: {
				1: {
					1: {
						2: NUM_73,
					},
				},
				2: {
					1: {
						1: NUM_74,
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
						3: SYM_REVERSE_STOP,
					},
				},
				2: {
					1: {
						4: SYM_START_B,
					},
					3: {
						2: SYM_START_C,
					},
				},
				3: {
					1: {
						3: NUM_39,
					},
					3: {
						1: NUM_49,
					},
				},
				4: {
					1: {
						2: SYM_START_A,
					},
				},
			},
			2: {
				1: {
					2: {
						3: NUM_30,
					},
					4: {
						1: NUM_89,
					},
				},
				2: {
					2: {
						2: NUM_00,
					},
				},
				3: {
					2: {
						1: NUM_31,
					},
				},
			},
			3: {
				1: {
					1: {
						3: NUM_51,
					},
					3: {
						1: NUM_53,
					},
				},
				2: {
					1: {
						2: NUM_21,
					},
				},
				3: {
					1: {
						1: NUM_52,
					},
				},
			},
			4: {
				1: {
					2: {
						1: NUM_90,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						4: NUM_76,
					},
					3: {
						2: NUM_19,
					},
				},
				2: {
					1: {
						3: NUM_09,
					},
					3: {
						1: NUM_20,
					},
				},
				3: {
					1: {
						2: NUM_10,
					},
				},
				4: {
					1: {
						1: NUM_61,
					},
				},
			},
			2: {
				1: {
					2: {
						2: NUM_01,
					},
				},
				2: {
					2: {
						1: NUM_02,
					},
				},
			},
			3: {
				1: {
					1: {
						2: NUM_22,
					},
				},
				2: {
					1: {
						1: NUM_18,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						3: NUM_40,
					},
					3: {
						1: NUM_50,
					},
				},
				2: {
					1: {
						2: NUM_11,
					},
				},
				3: {
					1: {
						1: NUM_41,
					},
				},
			},
			2: {
				1: {
					2: {
						1: NUM_32,
					},
				},
			},
			3: {
				1: {
					1: {
						1: SYM_STOP,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					1: {
						2: NUM_78,
					},
				},
				2: {
					1: {
						1: NUM_75,
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
						3: NUM_54,
					},
					4: {
						1: SYM_CODE_A,
					},
				},
				2: {
					2: {
						2: NUM_24,
					},
				},
				3: {
					2: {
						1: NUM_55,
					},
				},
			},
			2: {
				1: {
					1: {
						3: NUM_57,
					},
					3: {
						1: NUM_23,
					},
				},
				2: {
					1: {
						2: NUM_27,
					},
				},
				3: {
					1: {
						1: NUM_58,
					},
				},
			},
			3: {
				1: {
					2: {
						1: NUM_48,
					},
				},
			},
			4: {
				1: {
					1: {
						1: NUM_60,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						2: NUM_25,
					},
				},
				2: {
					2: {
						1: NUM_26,
					},
				},
			},
			2: {
				1: {
					1: {
						2: NUM_28,
					},
				},
				2: {
					1: {
						1: NUM_29,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					2: {
						1: NUM_56,
					},
				},
			},
			2: {
				1: {
					1: {
						1: NUM_59,
					},
				},
			},
		},
	},
	4: {
		1: {
			1: {
				1: {
					1: {
						3: NUM_97,
					},
					3: {
						1: SYM_FNC1,
					},
				},
				2: {
					1: {
						2: NUM_86,
					},
				},
				3: {
					1: {
						1: NUM_98,
					},
				},
			},
			2: {
				1: {
					2: {
						1: NUM_91,
					},
				},
			},
			3: {
				1: {
					1: {
						1: NUM_77,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						2: NUM_87,
					},
				},
				2: {
					1: {
						1: NUM_88,
					},
				},
			},
		},
		3: {
			1: {
				1: {
					1: {
						1: NUM_62,
					},
				},
			},
		},
	},
}

var CharTableA = map[int]string{
	SYM_SPACE:             " ",
	SYM_EXCLAMATION:       "!",
	SYM_DOUBLE_QUOTE:      `"`,
	SYM_POUND:             "#",
	SYM_DOLLAR:            "$",
	SYM_PERCENT:           "%",
	SYM_AMPERSAND:         "&",
	SYM_SINGLE_QUOTE:      "'",
	SYM_OPEN_PARENTHESIS:  "(",
	SYM_CLOSE_PARENTHESIS: ")",
	SYM_ASTERISK:          "*",
	SYM_PLUS:              "+",
	SYM_COMMA:             ",",
	SYM_HYPHEN:            "-",
	SYM_PERIOD:            ".",
	SYM_SLASH:             "/",
	SYM_ZERO:              "0",
	SYM_ONE:               "1",
	SYM_TWO:               "2",
	SYM_THREE:             "3",
	SYM_FOUR:              "4",
	SYM_FIVE:              "5",
	SYM_SIX:               "6",
	SYM_SEVEN:             "7",
	SYM_EIGHT:             "8",
	SYM_NINE:              "9",
	SYM_COLON:             ":",
	SYM_SEMICOLON:         ";",
	SYM_LESS_THAN:         "<",
	SYM_EQUAL:             "=",
	SYM_GREATER_THAN:      ">",
	SYM_QUESTION:          "?",
	SYM_AT:                "@",
	SYM_A:                 "A",
	SYM_B:                 "B",
	SYM_C:                 "C",
	SYM_D:                 "D",
	SYM_E:                 "E",
	SYM_F:                 "F",
	SYM_G:                 "G",
	SYM_H:                 "H",
	SYM_I:                 "I",
	SYM_J:                 "J",
	SYM_K:                 "K",
	SYM_L:                 "L",
	SYM_M:                 "M",
	SYM_N:                 "N",
	SYM_O:                 "O",
	SYM_P:                 "P",
	SYM_Q:                 "Q",
	SYM_R:                 "R",
	SYM_S:                 "S",
	SYM_T:                 "T",
	SYM_U:                 "U",
	SYM_V:                 "V",
	SYM_W:                 "W",
	SYM_X:                 "X",
	SYM_Y:                 "Y",
	SYM_Z:                 "Z",
	SYM_OPEN_BRACKET:      "[",
	SYM_BACKSLASH:         `\`,
	SYM_CLOSE_BRACKET:     "]",
	SYM_CARET:             "^",
	SYM_UNDERSCORE:        "_",
	SYM_NUL:               "\000",
	SYM_SOH:               "\001",
	SYM_STX:               "\002",
	SYM_ETX:               "\003",
	SYM_EOT:               "\004",
	SYM_ENQ:               "\005",
	SYM_ACK:               "\006",
	SYM_BEL:               "\a",
	SYM_BS:                "\b",
	SYM_HT:                "\t",
	SYM_LF:                "\n",
	SYM_VT:                "\v",
	SYM_FF:                "\f",
	SYM_CR:                "\r",
	SYM_SO:                "\016",
	SYM_SI:                "\017",
	SYM_DLE:               "\020",
	SYM_DC1:               "\021",
	SYM_DC2:               "\022",
	SYM_DC3:               "\023",
	SYM_DC4:               "\024",
	SYM_NAK:               "\025",
	SYM_SYN:               "\026",
	SYM_ETB:               "\027",
	SYM_CAN:               "\030",
	SYM_EM:                "\031",
	SYM_SUB:               "\032",
	SYM_ESC:               "\033",
	SYM_FS:                "\034",
	SYM_GS:                "\035",
	SYM_RS:                "\036",
	SYM_US:                "\037",
	SYM_FNC3:              "<FNC3>",
	SYM_FNC2:              "<FNC2>",
	SYM_SHIFT_B:           "<SHIFT_B>",
	SYM_CODE_C:            "<CODE_C>",
	SYM_CODE_B:            "<CODE_B>",
	SYM_FNC4_A:            "<FNC4>",
	SYM_FNC1:              "<FNC1>",
	SYM_START_A:           "<START_A>",
	SYM_START_B:           "<START_B>",
	SYM_START_C:           "<START_C>",
	SYM_STOP:              "<STOP>",
	SYM_REVERSE_STOP:      "<REVERSE_STOP>",
}

var CharTableB = map[int]string{
	SYM_SPACE:             " ",
	SYM_EXCLAMATION:       "!",
	SYM_DOUBLE_QUOTE:      `"`,
	SYM_POUND:             "#",
	SYM_DOLLAR:            "$",
	SYM_PERCENT:           "%",
	SYM_AMPERSAND:         "&",
	SYM_SINGLE_QUOTE:      "'",
	SYM_OPEN_PARENTHESIS:  "(",
	SYM_CLOSE_PARENTHESIS: ")",
	SYM_ASTERISK:          "*",
	SYM_PLUS:              "+",
	SYM_COMMA:             ",",
	SYM_HYPHEN:            "-",
	SYM_PERIOD:            ".",
	SYM_SLASH:             "/",
	SYM_ZERO:              "0",
	SYM_ONE:               "1",
	SYM_TWO:               "2",
	SYM_THREE:             "3",
	SYM_FOUR:              "4",
	SYM_FIVE:              "5",
	SYM_SIX:               "6",
	SYM_SEVEN:             "7",
	SYM_EIGHT:             "8",
	SYM_NINE:              "9",
	SYM_COLON:             ":",
	SYM_SEMICOLON:         ";",
	SYM_LESS_THAN:         "<",
	SYM_EQUAL:             "=",
	SYM_GREATER_THAN:      ">",
	SYM_QUESTION:          "?",
	SYM_AT:                "@",
	SYM_A:                 "A",
	SYM_B:                 "B",
	SYM_C:                 "C",
	SYM_D:                 "D",
	SYM_E:                 "E",
	SYM_F:                 "F",
	SYM_G:                 "G",
	SYM_H:                 "H",
	SYM_I:                 "I",
	SYM_J:                 "J",
	SYM_K:                 "K",
	SYM_L:                 "L",
	SYM_M:                 "M",
	SYM_N:                 "N",
	SYM_O:                 "O",
	SYM_P:                 "P",
	SYM_Q:                 "Q",
	SYM_R:                 "R",
	SYM_S:                 "S",
	SYM_T:                 "T",
	SYM_U:                 "U",
	SYM_V:                 "V",
	SYM_W:                 "W",
	SYM_X:                 "X",
	SYM_Y:                 "Y",
	SYM_Z:                 "Z",
	SYM_OPEN_BRACKET:      "[",
	SYM_BACKSLASH:         `\`,
	SYM_CLOSE_BRACKET:     "]",
	SYM_CARET:             "^",
	SYM_UNDERSCORE:        "_",
	SYM_BACKTICK:          "`",
	SYM_a:                 "a",
	SYM_b:                 "b",
	SYM_c:                 "c",
	SYM_d:                 "d",
	SYM_e:                 "e",
	SYM_f:                 "f",
	SYM_g:                 "g",
	SYM_h:                 "h",
	SYM_i:                 "i",
	SYM_j:                 "j",
	SYM_k:                 "k",
	SYM_l:                 "l",
	SYM_m:                 "m",
	SYM_n:                 "n",
	SYM_o:                 "o",
	SYM_p:                 "p",
	SYM_q:                 "q",
	SYM_r:                 "r",
	SYM_s:                 "s",
	SYM_t:                 "t",
	SYM_u:                 "u",
	SYM_v:                 "v",
	SYM_w:                 "w",
	SYM_x:                 "x",
	SYM_y:                 "y",
	SYM_z:                 "z",
	SYM_OPEN_BRACE:        "{",
	SYM_PIPE:              "|",
	SYM_CLOSE_BRACE:       "}",
	SYM_TILDE:             "~",
	SYM_DEL:               "\177",
	SYM_FNC3:              "<FNC3>",
	SYM_FNC2:              "<FNC2>",
	SYM_SHIFT_A:           "<SHIFT_A>",
	SYM_CODE_C:            "<CODE_C>",
	SYM_FNC4_B:            "<FNC4>",
	SYM_CODE_A:            "<CODE_A>",
	SYM_FNC1:              "<FNC1>",
	SYM_START_A:           "<START_A>",
	SYM_START_B:           "<START_B>",
	SYM_START_C:           "<START_C>",
	SYM_STOP:              "<STOP>",
	SYM_REVERSE_STOP:      "<REVERSE_STOP>",
}

var CharTableC = map[int]string{
	NUM_00:           "00",
	NUM_01:           "01",
	NUM_02:           "02",
	NUM_03:           "03",
	NUM_04:           "04",
	NUM_05:           "05",
	NUM_06:           "06",
	NUM_07:           "07",
	NUM_08:           "08",
	NUM_09:           "09",
	NUM_10:           "10",
	NUM_11:           "11",
	NUM_12:           "12",
	NUM_13:           "13",
	NUM_14:           "14",
	NUM_15:           "15",
	NUM_16:           "16",
	NUM_17:           "17",
	NUM_18:           "18",
	NUM_19:           "19",
	NUM_20:           "20",
	NUM_21:           "21",
	NUM_22:           "22",
	NUM_23:           "23",
	NUM_24:           "24",
	NUM_25:           "25",
	NUM_26:           "26",
	NUM_27:           "27",
	NUM_28:           "28",
	NUM_29:           "29",
	NUM_30:           "30",
	NUM_31:           "31",
	NUM_32:           "32",
	NUM_33:           "33",
	NUM_34:           "34",
	NUM_35:           "35",
	NUM_36:           "36",
	NUM_37:           "37",
	NUM_38:           "38",
	NUM_39:           "39",
	NUM_40:           "40",
	NUM_41:           "41",
	NUM_42:           "42",
	NUM_43:           "43",
	NUM_44:           "44",
	NUM_45:           "45",
	NUM_46:           "46",
	NUM_47:           "47",
	NUM_48:           "48",
	NUM_49:           "49",
	NUM_50:           "50",
	NUM_51:           "51",
	NUM_52:           "52",
	NUM_53:           "53",
	NUM_54:           "54",
	NUM_55:           "55",
	NUM_56:           "56",
	NUM_57:           "57",
	NUM_58:           "58",
	NUM_59:           "59",
	NUM_60:           "60",
	NUM_61:           "61",
	NUM_62:           "62",
	NUM_63:           "63",
	NUM_64:           "64",
	NUM_65:           "65",
	NUM_66:           "66",
	NUM_67:           "67",
	NUM_68:           "68",
	NUM_69:           "69",
	NUM_70:           "70",
	NUM_71:           "71",
	NUM_72:           "72",
	NUM_73:           "73",
	NUM_74:           "74",
	NUM_75:           "75",
	NUM_76:           "76",
	NUM_77:           "77",
	NUM_78:           "78",
	NUM_79:           "79",
	NUM_80:           "80",
	NUM_81:           "81",
	NUM_82:           "82",
	NUM_83:           "83",
	NUM_84:           "84",
	NUM_85:           "85",
	NUM_86:           "86",
	NUM_87:           "87",
	NUM_88:           "88",
	NUM_89:           "89",
	NUM_90:           "90",
	NUM_91:           "91",
	NUM_92:           "92",
	NUM_93:           "93",
	NUM_94:           "94",
	NUM_95:           "95",
	NUM_96:           "96",
	NUM_97:           "97",
	NUM_98:           "98",
	NUM_99:           "99",
	SYM_CODE_B:       "<CODE_B>",
	SYM_CODE_A:       "<CODE_A>",
	SYM_FNC1:         "<FNC1>",
	SYM_START_A:      "<START_A>",
	SYM_START_B:      "<START_B>",
	SYM_START_C:      "<START_C>",
	SYM_STOP:         "<STOP>",
	SYM_REVERSE_STOP: "<REVERSE_STOP>",
}
