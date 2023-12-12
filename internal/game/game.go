package game

type Player bool

const (
	White Player = true
	Black Player = false
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
	Player   Player
	Value    int
	HasMoved bool
}

type State struct {
	Board           [120]*Piece
	CurrPlayer      Player
	Score           float64
	EnPassantSquare int

	// Prev  *State not sure if we need this
}

// validMoves returns a list of valid moves for the piece at the given square
// THIS DOES NOT HANDLE CHECK
func (state *State) validMoves(square int) []int {
	// You cannot move a piece that doesn't exist
	if state.Board[square] == nil {
		return nil
	}

	currPiece := state.Board[square]

	// You cannot move a piece that isn't yours
	if currPiece.Player != state.CurrPlayer {
		return nil
	}

	validMoves := make([]int, 0)

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

				// 2. Ensure the move does not result in a capture of your own piece
				if state.Board[target] != nil && state.Board[target].Player == state.CurrPlayer {
					// 2.1 If the piece is a knight continue to process moves
					if currPiece.Type == Knight {
						continue
					}
					break
				}

				// 3. Handle king and pawn special cases
				if currPiece.Type == Pawn {

					// 0. Ensure that a pawn does not move backwards
					if currPiece.Player == White && sign == 1 {
						break
					} else if currPiece.Player == Black && sign == -1 {
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
				} else if currPiece.Type == King {
					// todo: handle castling
					// 1. Ensure that there are no pieces between the king and rook
					// 2. Ensure that neither the king nor the rook have moved
					// 3. Ensure that the king is not in check
					// 4. Ensure that a king does not move into OR through check
				}

				validMoves = append(validMoves, target)

				// if the move vector is blocked, we should break
				if state.Board[target] != nil && currPiece.Type != Knight {
					break
				}
			}

		}
	}

	return nil
}

func isValid(target Piece, player Player) bool {
	return false
}

func (state *State) isUnderAttack(square int) bool {

	return false
}

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
