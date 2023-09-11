// Package code128 implements encode and decode for code128.
package code128

import (
	"errors"
	"fmt"
	"image"
	"reflect"
)

const (
	SPACE, NUM_00             = 0x00, 0x00
	EXCLAMATION, NUM_01       = 0x01, 0x01
	DOUBLE_QUOTE, NUM_02      = 0x02, 0x02
	POUND, NUM_03             = 0x03, 0x03
	DOLLAR, NUM_04            = 0x04, 0x04
	PERCENT, NUM_05           = 0x05, 0x05
	AMPERSAND, NUM_06         = 0x06, 0x06
	SINGLE_QUOTE, NUM_07      = 0x07, 0x07
	OPEN_PARENTHESIS, NUM_08  = 0x08, 0x08
	CLOSE_PARENTHESIS, NUM_09 = 0x09, 0x09
	ASTERISK, NUM_10          = 0x0a, 0x0a
	PLUS, NUM_11              = 0x0b, 0x0b
	COMMA, NUM_12             = 0x0c, 0x0c
	HYPHEN, NUM_13            = 0x0d, 0x0d
	PERIOD, NUM_14            = 0x0e, 0x0e
	SLASH, NUM_15             = 0x0f, 0x0f
	ZERO, NUM_16              = 0x10, 0x10
	ONE, NUM_17               = 0x11, 0x11
	TWO, NUM_18               = 0x12, 0x12
	THREE, NUM_19             = 0x13, 0x13
	FOUR, NUM_20              = 0x14, 0x14
	FIVE, NUM_21              = 0x15, 0x15
	SIX, NUM_22               = 0x16, 0x16
	SEVEN, NUM_23             = 0x17, 0x17
	EIGHT, NUM_24             = 0x18, 0x18
	NINE, NUM_25              = 0x19, 0x19
	COLON, NUM_26             = 0x1a, 0x1a
	SEMICOLON, NUM_27         = 0x1b, 0x1b
	LESS_THAN, NUM_28         = 0x1c, 0x1c
	EQUAL, NUM_29             = 0x1d, 0x1d
	GREATER_THAN, NUM_30      = 0x1e, 0x1e
	QUESTION, NUM_31          = 0x1f, 0x1f
	AT, NUM_32                = 0x20, 0x20
	A, NUM_33                 = 0x21, 0x21
	B, NUM_34                 = 0x22, 0x22
	C, NUM_35                 = 0x23, 0x23
	D, NUM_36                 = 0x24, 0x24
	E, NUM_37                 = 0x25, 0x25
	F, NUM_38                 = 0x26, 0x26
	G, NUM_39                 = 0x27, 0x27
	H, NUM_40                 = 0x28, 0x28
	I, NUM_41                 = 0x29, 0x29
	J, NUM_42                 = 0x2a, 0x2a
	K, NUM_43                 = 0x2b, 0x2b
	L, NUM_44                 = 0x2c, 0x2c
	M, NUM_45                 = 0x2d, 0x2d
	N, NUM_46                 = 0x2e, 0x2e
	O, NUM_47                 = 0x2f, 0x2f
	P, NUM_48                 = 0x30, 0x30
	Q, NUM_49                 = 0x31, 0x31
	R, NUM_50                 = 0x32, 0x32
	S, NUM_51                 = 0x33, 0x33
	T, NUM_52                 = 0x34, 0x34
	U, NUM_53                 = 0x35, 0x35
	V, NUM_54                 = 0x36, 0x36
	W, NUM_55                 = 0x37, 0x37
	X, NUM_56                 = 0x38, 0x38
	Y, NUM_57                 = 0x39, 0x39
	Z, NUM_58                 = 0x3a, 0x3a
	OPEN_BRACKET, NUM_59      = 0x3b, 0x3b
	BACKSLASH, NUM_60         = 0x3c, 0x3c
	CLOSE_BRACKET, NUM_61     = 0x3d, 0x3d
	CARET, NUM_62             = 0x3e, 0x3e
	UNDERSCORE, NUM_63        = 0x3f, 0x3f
	NUL, BACKTICK, NUM_64     = 0x40, 0x40, 0x40
	SOH, a, NUM_65            = 0x41, 0x41, 0x41
	STX, b, NUM_66            = 0x42, 0x42, 0x42
	ETX, c, NUM_67            = 0x43, 0x43, 0x43
	EOT, d, NUM_68            = 0x44, 0x44, 0x44
	ENQ, e, NUM_69            = 0x45, 0x45, 0x45
	ACK, f, NUM_70            = 0x46, 0x46, 0x46
	BEL, g, NUM_71            = 0x47, 0x47, 0x47
	BS, h, NUM_72             = 0x48, 0x48, 0x48
	HT, i, NUM_73             = 0x49, 0x49, 0x49
	LF, j, NUM_74             = 0x4a, 0x4a, 0x4a
	VT, k, NUM_75             = 0x4b, 0x4b, 0x4b
	FF, l, NUM_76             = 0x4c, 0x4c, 0x4c
	CR, m, NUM_77             = 0x4d, 0x4d, 0x4d
	SO, n, NUM_78             = 0x4e, 0x4e, 0x4e
	SI, o, NUM_79             = 0x4f, 0x4f, 0x4f
	DLE, p, NUM_80            = 0x50, 0x50, 0x50
	DC1, q, NUM_81            = 0x51, 0x51, 0x51
	DC2, r, NUM_82            = 0x52, 0x52, 0x52
	DC3, s, NUM_83            = 0x53, 0x53, 0x53
	DC4, t, NUM_84            = 0x54, 0x54, 0x54
	NAK, u, NUM_85            = 0x55, 0x55, 0x55
	SYN, v, NUM_86            = 0x56, 0x56, 0x56
	ETB, w, NUM_87            = 0x57, 0x57, 0x57
	CAN, x, NUM_88            = 0x58, 0x58, 0x58
	EM, y, NUM_89             = 0x59, 0x59, 0x59
	SUB, z, NUM_90            = 0x5a, 0x5a, 0x5a
	ESC, OPEN_BRACE, NUM_91   = 0x5b, 0x5b, 0x5b
	FS, PIPE, NUM_92          = 0x5c, 0x5c, 0x5c
	GS, CLOSE_BRACE, NUM_93   = 0x5d, 0x5d, 0x5d
	RS, TILDE, NUM_94         = 0x5e, 0x5e, 0x5e
	US, DEL, NUM_95           = 0x5f, 0x5f, 0x5f
	FNC3, NUM_96              = 0x60, 0x60
	FNC2, NUM_97              = 0x61, 0x61
	SHIFT_B, SHIFT_A, NUM_98  = 0x62, 0x62, 0x62
	CODE_C, NUM_99            = 0x63, 0x63
	CODE_B, FNC4_B            = 0x64, 0x64
	FNC4_A, CODE_A            = 0x65, 0x65
	FNC1                      = 0x66
	START_A                   = 0x67
	START_B                   = 0x68
	START_C                   = 0x69
	STOP                      = 0x6a
	REVERSE_STOP              = -1
)

var DecodeTableA = [][][][][][]int{
	1: {
		1: {
			1: {
				1: {
					4: {
						3: FS,
					},
				},
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
					4: {
						1: GS,
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
					4: {
						1: CODE_C,
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
			4: {
				1: {
					1: {
						3: US,
					},
					3: {
						1: CODE_B,
					},
				},
				2: {
					1: {
						2: DC3,
					},
				},
				3: {
					1: {
						1: FNC3,
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
					4: {
						2: DC1,
					},
				},
				2: {
					2: {
						3: POUND,
					},
					4: {
						1: DC2,
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
			4: {
				1: {
					1: {
						2: DC4,
					},
				},
				2: {
					1: {
						1: NAK,
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
					4: {
						1: RS,
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
					4: {
						1: EM,
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
			4: {
				1: {
					2: {
						1: SUB,
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
			3: {
				1: {
					1: {
						1: STOP,
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
					4: {
						1: FNC4_A,
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
			1: {
				1: {
					1: {
						3: FNC2,
					},
					3: {
						1: FNC1,
					},
				},
				2: {
					1: {
						2: SYN,
					},
				},
				3: {
					1: {
						1: SHIFT_B,
					},
				},
			},
			2: {
				1: {
					2: {
						1: ESC,
					},
				},
			},
			3: {
				1: {
					1: {
						1: CR,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						2: ETB,
					},
				},
				2: {
					1: {
						1: CAN,
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
				1: {
					4: {
						3: PIPE,
					},
				},
				2: {
					2: {
						4: UNDERSCORE,
					},
					4: {
						2: p,
					},
				},
				3: {
					2: {
						3: A,
					},
					4: {
						1: CLOSE_BRACE,
					},
				},
				4: {
					2: {
						2: BACKTICK,
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
						4: e,
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
						2: f,
					},
				},
			},
			3: {
				1: {
					2: {
						3: M,
					},
					4: {
						1: CODE_C,
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
			4: {
				1: {
					1: {
						3: DEL,
					},
					3: {
						1: FNC4_B,
					},
				},
				2: {
					1: {
						2: s,
					},
				},
				3: {
					1: {
						1: FNC3,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					2: {
						4: a,
					},
					4: {
						2: q,
					},
				},
				2: {
					2: {
						3: POUND,
					},
					4: {
						1: r,
					},
				},
				3: {
					2: {
						2: DOLLAR,
					},
				},
				4: {
					2: {
						1: b,
					},
				},
			},
			2: {
				1: {
					1: {
						4: g,
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
						1: h,
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
			4: {
				1: {
					1: {
						2: t,
					},
				},
				2: {
					1: {
						1: u,
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
					4: {
						1: TILDE,
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
						1: o,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					2: {
						2: c,
					},
				},
				2: {
					2: {
						1: d,
					},
				},
			},
			2: {
				1: {
					1: {
						2: i,
					},
				},
				2: {
					1: {
						1: j,
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
					4: {
						1: y,
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
			4: {
				1: {
					2: {
						1: z,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						4: l,
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
			3: {
				1: {
					1: {
						1: STOP,
					},
				},
			},
		},
		4: {
			1: {
				1: {
					1: {
						2: n,
					},
				},
				2: {
					1: {
						1: k,
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
					4: {
						1: CODE_A,
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
			1: {
				1: {
					1: {
						3: FNC2,
					},
					3: {
						1: FNC1,
					},
				},
				2: {
					1: {
						2: v,
					},
				},
				3: {
					1: {
						1: SHIFT_A,
					},
				},
			},
			2: {
				1: {
					2: {
						1: OPEN_BRACE,
					},
				},
			},
			3: {
				1: {
					1: {
						1: m,
					},
				},
			},
		},
		2: {
			1: {
				1: {
					1: {
						2: w,
					},
				},
				2: {
					1: {
						1: x,
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
						1: CODE_B,
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
						3: NUM_39,
					},
					3: {
						1: NUM_49,
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
						1: STOP,
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
						1: CODE_A,
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
						1: FNC1,
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
	ENQ:               "<ENQ>",
	ACK:               "<ACK>",
	BEL:               "<BEL>",
	BS:                "<BS>",
	HT:                "<HT>",
	LF:                "<LF>",
	VT:                "<VT>",
	FF:                "<FF>",
	CR:                "<CR>",
	SO:                "<SO>",
	SI:                "<SI>",
	DLE:               "<DLE>",
	DC1:               "<DC1>",
	DC2:               "<DC2>",
	DC3:               "<DC3>",
	DC4:               "<DC4>",
	NAK:               "<NAK>",
	SYN:               "<SYN>",
	ETB:               "<ETB>",
	CAN:               "<CAN>",
	EM:                "<EM>",
	SUB:               "<SUB>",
	ESC:               "<ESC>",
	FS:                "<FS>",
	GS:                "<GS>",
	RS:                "<RS>",
	US:                "<US>",
	FNC3:              "<FNC3>",
	FNC2:              "<FNC2>",
	SHIFT_B:           "<SHIFT_B>",
	CODE_C:            "<CODE_C>",
	CODE_B:            "<CODE_B>",
	FNC4_A:            "<FNC4>",
	FNC1:              "<FNC1>",
	START_A:           "<START_A>",
	START_B:           "<START_B>",
	START_C:           "<START_C>",
	STOP:              "<STOP>",
	REVERSE_STOP:      "<REVERSE_STOP>",
}

var CharTableB = map[int]string{
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
	BACKTICK:          "`",
	a:                 "a",
	b:                 "b",
	c:                 "c",
	d:                 "d",
	e:                 "e",
	f:                 "f",
	g:                 "g",
	h:                 "h",
	i:                 "i",
	j:                 "j",
	k:                 "k",
	l:                 "l",
	m:                 "m",
	n:                 "n",
	o:                 "o",
	p:                 "p",
	q:                 "q",
	r:                 "r",
	s:                 "s",
	t:                 "t",
	u:                 "u",
	v:                 "v",
	w:                 "w",
	x:                 "x",
	y:                 "y",
	z:                 "z",
	OPEN_BRACE:        "{",
	PIPE:              "|",
	CLOSE_BRACE:       "}",
	TILDE:             "~",
	DEL:               "<DEL>",
	FNC3:              "<FNC3>",
	FNC2:              "<FNC2>",
	SHIFT_A:           "<SHIFT_A>",
	CODE_C:            "<CODE_C>",
	FNC4_B:            "<FNC4>",
	CODE_A:            "<CODE_A>",
	FNC1:              "<FNC1>",
	START_A:           "<START_A>",
	START_B:           "<START_B>",
	START_C:           "<START_C>",
	STOP:              "<STOP>",
	REVERSE_STOP:      "<REVERSE_STOP>",
}

var CharTableC = map[int]string{
	NUM_00:       "00",
	NUM_01:       "01",
	NUM_02:       "02",
	NUM_03:       "03",
	NUM_04:       "04",
	NUM_05:       "05",
	NUM_06:       "06",
	NUM_07:       "07",
	NUM_08:       "08",
	NUM_09:       "09",
	NUM_10:       "10",
	NUM_11:       "11",
	NUM_12:       "12",
	NUM_13:       "13",
	NUM_14:       "14",
	NUM_15:       "15",
	NUM_16:       "16",
	NUM_17:       "17",
	NUM_18:       "18",
	NUM_19:       "19",
	NUM_20:       "20",
	NUM_21:       "21",
	NUM_22:       "22",
	NUM_23:       "23",
	NUM_24:       "24",
	NUM_25:       "25",
	NUM_26:       "26",
	NUM_27:       "27",
	NUM_28:       "28",
	NUM_29:       "29",
	NUM_30:       "30",
	NUM_31:       "31",
	NUM_32:       "32",
	NUM_33:       "33",
	NUM_34:       "34",
	NUM_35:       "35",
	NUM_36:       "36",
	NUM_37:       "37",
	NUM_38:       "38",
	NUM_39:       "39",
	NUM_40:       "40",
	NUM_41:       "41",
	NUM_42:       "42",
	NUM_43:       "43",
	NUM_44:       "44",
	NUM_45:       "45",
	NUM_46:       "46",
	NUM_47:       "47",
	NUM_48:       "48",
	NUM_49:       "49",
	NUM_50:       "50",
	NUM_51:       "51",
	NUM_52:       "52",
	NUM_53:       "53",
	NUM_54:       "54",
	NUM_55:       "55",
	NUM_56:       "56",
	NUM_57:       "57",
	NUM_58:       "58",
	NUM_59:       "59",
	NUM_60:       "60",
	NUM_61:       "61",
	NUM_62:       "62",
	NUM_63:       "63",
	NUM_64:       "64",
	NUM_65:       "65",
	NUM_66:       "66",
	NUM_67:       "67",
	NUM_68:       "68",
	NUM_69:       "69",
	NUM_70:       "70",
	NUM_71:       "71",
	NUM_72:       "72",
	NUM_73:       "73",
	NUM_74:       "74",
	NUM_75:       "75",
	NUM_76:       "76",
	NUM_77:       "77",
	NUM_78:       "78",
	NUM_79:       "79",
	NUM_80:       "80",
	NUM_81:       "81",
	NUM_82:       "82",
	NUM_83:       "83",
	NUM_84:       "84",
	NUM_85:       "85",
	NUM_86:       "86",
	NUM_87:       "87",
	NUM_88:       "88",
	NUM_89:       "89",
	NUM_90:       "90",
	NUM_91:       "91",
	NUM_92:       "92",
	NUM_93:       "93",
	NUM_94:       "94",
	NUM_95:       "95",
	NUM_96:       "96",
	NUM_97:       "97",
	NUM_98:       "98",
	NUM_99:       "99",
	CODE_B:       "<CODE_B>",
	CODE_A:       "<CODE_A>",
	FNC1:         "<FNC1>",
	START_A:      "<START_A>",
	START_B:      "<START_B>",
	START_C:      "<START_C>",
	STOP:         "<STOP>",
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
	charTable := CharTableA
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
		charTable = CharTableA
	case START_B:
		decodeTable = DecodeTableB
		charTable = CharTableB
	case START_C:
		decodeTable = DecodeTableC
		charTable = CharTableC
	default:
		return msg, fmt.Errorf("invalid start symbol: %s -- %+v", charTable[staSym], sta)
	}

	checksum := staSym

	for current < len(d) {
		sym := decodeTable[d[current-5]][d[current-4]][d[current-3]][d[current-2]][d[current-1]][d[current-0]]
		checksum += sym * posMul
		fmt.Printf("Sym: %s (%d%d%d%d%d%d) [%d × %d = %d]\n", charTable[sym], d[current-5], d[current-4], d[current-3], d[current-2], d[current-1], d[current-0], sym, posMul, sym*posMul)

		switch sym {
		case CODE_A:
			decodeTable = DecodeTableA
		case CODE_B:
			decodeTable = DecodeTableB
		case CODE_C:
			decodeTable = DecodeTableC
		default:
			msg += string(charTable[sym])
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
