package game

import (
	"fmt"
	"math"
)

type Color bool

const (
	White Color = true
	Black Color = false
)

type MoveType int

const (
	BlackQueenSideCastle MoveType = 10
	BlackKingSideCastle  MoveType = 20
	WhiteQueenSideCastle MoveType = 30
	WhiteKingSideCastle  MoveType = 40

	Aggressive MoveType = 1
	Neutral    MoveType = 0
	Passive    MoveType = -1
	EnPassant  MoveType = -2
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

type GameState struct {
	board           [120]*Piece
	enPassantSquare int
	currColor       Color

	blackValidMoves map[int]map[int]MoveType
	whiteValidMoves map[int]map[int]MoveType
}

func (gs *GameState) PrettyPrint() {
	for i, piece := range gs.board {
		if i%12 == 0 {
			fmt.Println()
		}
		if piece == nil {
			if _, ok := validSquares[i]; ok {
				fmt.Print(" - ")
			} else {
				fmt.Print("   ")
			}
			continue
		}

		switch piece.Type {
		case King:
			if piece.Color == White {
				fmt.Print(" ♚ ")
			} else {
				fmt.Print(" ♔ ")
			}
		case Queen:
			if piece.Color == White {
				fmt.Print(" ♛ ")
			} else {
				fmt.Print(" ♕ ")
			}
		case Rook:
			if piece.Color == White {
				fmt.Print(" ♜ ")
			} else {
				fmt.Print(" ♖ ")
			}
		case Bishop:
			if piece.Color == White {
				fmt.Print(" ♝ ")
			} else {
				fmt.Print(" ♗ ")
			}
		case Knight:
			if piece.Color == White {
				fmt.Print(" ♞ ")
			} else {
				fmt.Print(" ♘ ")
			}
		case Pawn:
			if piece.Color == White {
				fmt.Print(" ♟ ")
			} else {
				fmt.Print(" ♙ ")
			}
		}
	}
}

// getMovesPreCheck returns a list of valid moves for the piece at the given square
// THIS DOES NOT HANDLE CHECK
func (gs *GameState) getMovesPreCheck(square int) map[int]MoveType {
	// You cannot move a piece that doesn't exist
	if gs.board[square] == nil {
		return nil
	}

	currPiece := gs.board[square]

	validMoves := make(map[int]MoveType)

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
				if gs.board[target] != nil && gs.board[target].Color == currPiece.Color {
					// 2.1 If the piece is a knight continue to process moves
					if currPiece.Type == Knight {
						continue
					}
					break
				}

				// 3. Handle special pawn movement & en passant
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

					// 2. Handle forward movement
					if offset%12 == 0 {
						// 2.1 Ensure that a pawn does not move verically onto or through another piece
						if gs.board[target] != nil {
							break
						}
						validMoves[target] = Passive
						// 2.2 Only allow diagonal movement if there is a piece to capture
					} else if gs.board[target] != nil {
						validMoves[target] = Aggressive
						// 2.3 or the square is the en passant square & the pawn belongs to the current player
					} else if target == gs.enPassantSquare && gs.currColor == currPiece.Color {
						validMoves[target] = EnPassant
					}
					continue
				}

				validMoves[target] = Neutral

				// if the move vector is blocked, we should break
				if gs.board[target] != nil && currPiece.Type != Knight {
					break
				}
			}
		}
	}

	return validMoves
}

func (gs *GameState) executeMove(origin, destination int, moveType MoveType) {
	defer func() {
		gs.Update()
	}()

	// handle castling
	if moveType == WhiteKingSideCastle {
		// move king
		gs.board[104] = gs.board[102]
		gs.board[102] = nil
		gs.board[104].HasMoved = true

		// move rook
		gs.board[103] = gs.board[105]
		gs.board[105] = nil
		gs.board[103].HasMoved = true
		return
	} else if moveType == WhiteQueenSideCastle {
		// move king
		gs.board[100] = gs.board[102]
		gs.board[102] = nil
		gs.board[100].HasMoved = true

		// move rook
		gs.board[101] = gs.board[98]
		gs.board[98] = nil
		gs.board[101].HasMoved = true
		return
	} else if moveType == BlackKingSideCastle {
		// move king
		gs.board[20] = gs.board[18]
		gs.board[18] = nil
		gs.board[20].HasMoved = true

		// move rook
		gs.board[19] = gs.board[21]
		gs.board[21] = nil
		gs.board[19].HasMoved = true
		return
	} else if moveType == BlackQueenSideCastle {
		// move king
		gs.board[16] = gs.board[18]
		gs.board[18] = nil
		gs.board[16].HasMoved = true

		// move rook
		gs.board[17] = gs.board[14]
		gs.board[14] = nil
		gs.board[17].HasMoved = true
		return
	}

	// handle a generic move
	gs.board[destination] = gs.board[origin]
	gs.board[origin] = nil
	gs.board[destination].HasMoved = true

	// handle en passant attack
	if moveType == EnPassant {
		if gs.currColor == White {
			gs.board[destination+12] = nil
		} else {
			gs.board[destination-12] = nil
		}
	}

	// handle engaging a new en passant square
	if gs.board[destination].Type == Pawn && math.Abs(float64(destination-origin)) == 24 {
		gs.enPassantSquare = (destination + origin) / 2
	} else {
		gs.enPassantSquare = 0
	}
}

func NewGame() *GameState {
	gs := &GameState{
		board: newBoard(),
	}

	gs.Update()

	return gs
}

func (gs *GameState) Update() {

	// switch teams
	if gs.currColor == White {
		gs.currColor = Black
	} else {
		gs.currColor = White
	}

	gs.blackValidMoves = make(map[int]map[int]MoveType)
	gs.whiteValidMoves = make(map[int]map[int]MoveType)

	for i, piece := range gs.board {
		if piece == nil {
			continue
		}

		switch piece.Color {
		case White:
			gs.whiteValidMoves[i] = gs.getMovesPreCheck(i)
		case Black:
			gs.blackValidMoves[i] = gs.getMovesPreCheck(i)
		}
	}

	// check if white can castle on king side
	if gs.whiteKingSide() {
		gs.whiteValidMoves[102][104] = WhiteKingSideCastle
	}

	// check if white can castle on queen side
	if gs.whiteQueenSide() {
		gs.whiteValidMoves[102][100] = WhiteQueenSideCastle
	}

	// check if black can castle on king side
	if gs.blackKingSide() {
		gs.blackValidMoves[18][20] = BlackKingSideCastle
	}

	// check if black can castle on queen side
	if gs.blackQueenSide() {
		gs.blackValidMoves[18][16] = BlackQueenSideCastle
	}

	gs.PrettyPrint()
	return

}

// scrappy random game generator for naive testing
func RandomGame() {
	gs := NewGame()
	for i := 0; i < 100; i++ {

		if gs.currColor == White {
			for start, ends := range gs.whiteValidMoves {
				if len(ends) == 0 {
					continue
				}
				for end, mt := range ends {
					gs.executeMove(start, end, mt)
					break
				}
				break
			}
		} else {
			for start, ends := range gs.blackValidMoves {
				if len(ends) == 0 {
					continue
				}
				for end, mt := range ends {
					gs.executeMove(start, end, mt)
					break
				}
				break
			}
		}
	}
	fmt.Println("Game Over")
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
