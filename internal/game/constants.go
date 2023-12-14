package game

func newBoard() [120]*Piece {
	return [120]*Piece{
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), bR(), bN(), bB(), bQ(), bK(), bB(), bN(), bR(), __(), __(),
		__(), __(), bP(), bP(), bP(), bP(), bP(), bP(), bP(), bP(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), wP(), wP(), wP(), wP(), wP(), wP(), wP(), wP(), __(), __(),
		__(), __(), wR(), wN(), wB(), wQ(), wK(), wB(), wN(), wR(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
	}
}

func bR() *Piece { return &Piece{Type: Rook, Color: Black, Value: 5} }
func wR() *Piece { return &Piece{Type: Rook, Color: White, Value: 5} }
func bK() *Piece { return &Piece{Type: King, Color: Black, Value: 10000} }
func wK() *Piece { return &Piece{Type: King, Color: White, Value: 10000} }
func bQ() *Piece { return &Piece{Type: Queen, Color: Black, Value: 9} }
func wQ() *Piece { return &Piece{Type: Queen, Color: White, Value: 9} }
func bB() *Piece { return &Piece{Type: Bishop, Color: Black, Value: 3} }
func wB() *Piece { return &Piece{Type: Bishop, Color: White, Value: 3} }
func bN() *Piece { return &Piece{Type: Knight, Color: Black, Value: 3} }
func wN() *Piece { return &Piece{Type: Knight, Color: White, Value: 3} }
func bP() *Piece { return &Piece{Type: Pawn, Color: Black, Value: 1} }
func wP() *Piece { return &Piece{Type: Pawn, Color: White, Value: 1} }
func __() *Piece { return nil }

var moveVectors = map[Type][][]int8{
	King: {
		{1},  // horizontal
		{12}, // vertical
		{11}, // diagonal 1
		{13}, // diagonal 2
	},
	Queen: {
		{1, 2, 3, 4, 5, 6, 7},        // horizontal
		{12, 24, 36, 48, 60, 72, 84}, // vertical
		{11, 22, 33, 44, 55, 66, 77}, // diagonal 1
		{13, 26, 39, 52, 65, 78, 91}, // diagonal 2
	},
	Rook: {
		{1, 2, 3, 4, 5, 6, 7},        // horizontal
		{12, 24, 36, 48, 60, 72, 84}, // vertical
	},
	Bishop: {
		{11, 22, 33, 44, 55, 66, 77}, // diagonal 1
		{13, 26, 39, 52, 65, 78, 91}, // diagonal 2
	},
	Knight: {
		{10},
		{14},
		{23},
		{25},
	},
	Pawn: {
		{12, 24}, // vertical
		{11},     // diagonal 1
		{13},     // diagonal 2
	},
}

var validSquares = map[int8]struct{}{
	14: {}, 15: {}, 16: {}, 17: {}, 18: {}, 19: {}, 20: {}, 21: {},
	26: {}, 27: {}, 28: {}, 29: {}, 30: {}, 31: {}, 32: {}, 33: {},
	38: {}, 39: {}, 40: {}, 41: {}, 42: {}, 43: {}, 44: {}, 45: {},
	50: {}, 51: {}, 52: {}, 53: {}, 54: {}, 55: {}, 56: {}, 57: {},
	62: {}, 63: {}, 64: {}, 65: {}, 66: {}, 67: {}, 68: {}, 69: {},
	74: {}, 75: {}, 76: {}, 77: {}, 78: {}, 79: {}, 80: {}, 81: {},
	86: {}, 87: {}, 88: {}, 89: {}, 90: {}, 91: {}, 92: {}, 93: {},
	98: {}, 99: {}, 100: {}, 101: {}, 102: {}, 103: {}, 104: {}, 105: {},
}

var promotionSquares = map[int8]struct{}{
	14: {}, 15: {}, 16: {}, 17: {}, 18: {}, 19: {}, 20: {}, 21: {},
	98: {}, 99: {}, 100: {}, 101: {}, 102: {}, 103: {}, 104: {}, 105: {},
}
