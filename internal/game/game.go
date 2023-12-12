package game

import "math"

type Color bool

const (
	White Color = true
	Black Color = false
)

type Type int

const (
	King   Type = 0
	Queen  Type = 1
	Rook   Type = 2
	Bishop Type = 3
	Knight Type = 4
	Pawn   Type = 5
)

type Piece struct {
	Type     Type
	Color    Color
	Value    float64
	HasMoved bool
}

type State struct {
	Board           [120]*Piece
	CurrColor       Color
	Score           float64
	EnPassantSquare int

	// Prev  *State not sure if we need this
}

// getMovesPreCheck returns a list of valid moves for the piece at the given square
// THIS DOES NOT HANDLE CHECK
func (state *State) getMovesPreCheck(square int, withCastle bool) map[int]struct{} {
	// You cannot move a piece that doesn't exist
	if state.Board[square] == nil {
		return nil
	}

	currPiece := state.Board[square]

	validMoves := make(map[int]struct{})

	// Handle the movement vectors
	for _, vector := range moveVectors[currPiece.Type] {
		for _, sign := range []int{1, -1} {
			for _, offset := range vector {
				target := square + offset*sign

				// 1. Ensure the move is on the board
				if _, ok := validSquares[target]; !ok {
					// 1.1 If the piece is a knight continue to process moves
					if currPiece.Type == Knight {
						continue
					}
					break
				}

				// 2. Ensure the move does not capture a friendly piece
				if state.Board[target] != nil && state.Board[target].Color == currPiece.Color {
					// 2.1 If the piece is a knight continue to process moves
					if currPiece.Type == Knight {
						continue
					}
					break
				}

				// 3. Handle king and pawn special cases
				if currPiece.Type == Pawn {

					// 0. Ensure that a pawn does not move backwards
					if currPiece.Color == White && sign != -1 {
						break
					} else if currPiece.Color == Black && sign != 1 {
						break
					}

					// 1. Ensure that a pawn does not move two squares if it has already moved
					if currPiece.HasMoved && offset == 24 {
						continue
					}

					// 2. Ensure that a pawn does not move verically onto or through another piece
					if offset%12 == 0 && state.Board[target] != nil {
						break
					}

					// 3. Ensure that a pawn does not move diagonally onto an empty square / non-en-passant square
					if offset%12 != 0 && state.Board[target] == nil && target != state.EnPassantSquare {
						continue
					}
				} else if currPiece.Type == King && offset == 2 && withCastle {

					// 1. Ensure the king has not moved
					if currPiece.HasMoved {
						continue
					}

					// 2. Ensure that there are no pieces between the king and rook
					// 3. Ensure that the rook has not moved TODO
					switch target {
					case 104: // white king side
						if state.Board[103] != nil || state.Board[104] != nil {
							continue
						}
					case 100: // white queen side
						if state.Board[99] != nil || state.Board[100] != nil || state.Board[101] != nil {
							continue
						}
					case 19: // black king side
						if state.Board[19] != nil || state.Board[20] != nil {
							continue
						}
					case 16: // black queen side
						if state.Board[15] != nil || state.Board[16] != nil || state.Board[17] != nil {
							continue
						}
					}

					// 3. Ensure that the king is not in check
					if state.isUnderAttack(square, !currPiece.Color) {
						continue
					}

					// todo: handle castling
					// 1. Ensure that there are no pieces between the king and rook
					// 2. Ensure that neither the king nor the rook have moved

					// 4. Ensure that a king does not move into OR through check

				}

				validMoves[target] = struct{}{}

				// if the move vector is blocked, we should break
				if state.Board[target] != nil && currPiece.Type != Knight {
					break
				}
			}

		}
	}

	return validMoves
}

func (state *State) isUnderAttack(square int, by Color) bool {
	for i, piece := range state.Board {
		if piece.Color != by {
			continue
		}
		validMoves := state.getMovesPreCheck(i, false)
		if _, ok := validMoves[square]; ok {
			return true
		}
	}
	return false
}

func newBoard() [120]*Piece {
	return [120]*Piece{
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), BR(), Bk(), BB(), BQ(), BK(), BB(), Bk(), BR(), __(), __(),
		__(), __(), BP(), BP(), BP(), BP(), BP(), BP(), BP(), BP(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
		__(), __(), WP(), WP(), WP(), WP(), WP(), WP(), WP(), WP(), __(), __(),
		__(), __(), WR(), Wk(), WB(), WQ(), WK(), WB(), Wk(), WR(), __(), __(),
		__(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(), __(),
	}
}

func BR() *Piece { return &Piece{Type: Rook, Color: Black, Value: 5} }
func WR() *Piece { return &Piece{Type: Rook, Color: White, Value: 5} }
func BK() *Piece { return &Piece{Type: King, Color: Black, Value: math.Inf(1)} }
func WK() *Piece { return &Piece{Type: King, Color: White, Value: math.Inf(1)} }
func BQ() *Piece { return &Piece{Type: Queen, Color: Black, Value: 9} }
func WQ() *Piece { return &Piece{Type: Queen, Color: White, Value: 9} }
func BB() *Piece { return &Piece{Type: Bishop, Color: Black, Value: 3} }
func WB() *Piece { return &Piece{Type: Bishop, Color: White, Value: 3} }
func Bk() *Piece { return &Piece{Type: Knight, Color: Black, Value: 3} }
func Wk() *Piece { return &Piece{Type: Knight, Color: White, Value: 3} }
func BP() *Piece { return &Piece{Type: Pawn, Color: Black, Value: 1} }
func WP() *Piece { return &Piece{Type: Pawn, Color: White, Value: 1} }
func __() *Piece { return nil }

// generic valid moves for each piece, unsigned

var moveVectors = map[Type][][]int{
	King: {
		{1, 2}, // horizontal + castle
		{12},   // vertical
		{11},   // diagonal 1
		{13},   // diagonal 2
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
	Knight: {{
		10,
		14,
		23,
		25,
	}},
	Pawn: {
		{12, 24}, // vertical
		{11},     // diagonal 1
		{13},     // diagonal 2
	},
}

var validSquares = map[int]struct{}{
	14: {}, 15: {}, 16: {}, 17: {}, 18: {}, 19: {}, 20: {}, 21: {},
	26: {}, 27: {}, 28: {}, 29: {}, 30: {}, 31: {}, 32: {}, 33: {},
	38: {}, 39: {}, 40: {}, 41: {}, 42: {}, 43: {}, 44: {}, 45: {},
	50: {}, 51: {}, 52: {}, 53: {}, 54: {}, 55: {}, 56: {}, 57: {},
	62: {}, 63: {}, 64: {}, 65: {}, 66: {}, 67: {}, 68: {}, 69: {},
	74: {}, 75: {}, 76: {}, 77: {}, 78: {}, 79: {}, 80: {}, 81: {},
	86: {}, 87: {}, 88: {}, 89: {}, 90: {}, 91: {}, 92: {}, 93: {},
	98: {}, 99: {}, 100: {}, 101: {}, 102: {}, 103: {}, 104: {}, 105: {},
}
