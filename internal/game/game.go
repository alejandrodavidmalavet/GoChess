package game

import (
	"fmt"
)

type Piece struct {
	Type     Type
	Color    Color
	Value    float64
	HasMoved bool
}

type GameState struct {
	board           [120]*Piece
	enPassantSquare int8
	currColor       Color

	kingSquares map[Color]int8

	allMoves         map[Color]map[int8]map[int8]map[MoveType]struct{}
	attackingSquares map[Color]map[int8]struct{}

	history []*HistoryEntry
}

// NewGame returns a new game state with the board in the starting position
func NewGame() *GameState {
	board := newBoard()
	gs := &GameState{
		board:            board,
		currColor:        White,
		kingSquares:      map[Color]int8{White: 102, Black: 18},
		allMoves:         map[Color]map[int8]map[int8]map[MoveType]struct{}{White: {}, Black: {}},
		attackingSquares: map[Color]map[int8]struct{}{White: {}, Black: {}},
	}

	gs.update()

	return gs
}

// PrettyPrint prints the board in a human readable format to the console
func (gs *GameState) PrettyPrint() {
	for i, piece := range gs.board {
		if i%12 == 0 {
			fmt.Println()
		}
		if piece == nil {
			if _, ok := validSquares[int8(i)]; ok {
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

// CurrentPlayer returns the current player as a string
func (gs *GameState) CurrentPlayer() string {
	if gs.currColor == White {
		return "White"
	}
	return "Black"
}

// updateValidMoves returns a list of valid moves for the piece at the given square, without checking for check
func (gs *GameState) updateValidMoves(square int8) {
	// You cannot move a piece that doesn't exist
	if gs.board[square] == nil {
		return
	}

	currPiece := gs.board[square]

	gs.allMoves[currPiece.Color][square] = make(map[int8]map[MoveType]struct{})

	// Handle the movement vectors
	for _, vector := range moveVectors[currPiece.Type] {
		for _, direction := range []int8{1, -1} {
			for _, offset := range vector {
				target := square + offset*direction

				// 1. Ensure the move is on the board
				if _, ok := validSquares[target]; !ok {
					break
				}

				// 2. Ensure the move does not capture a friendly piece
				if gs.board[target] != nil && gs.board[target].Color == currPiece.Color {
					break
				}

				// 3. Handle pawns
				if currPiece.Type == Pawn {

					// a. Ensure that a pawn does not move backwards
					if int8(currPiece.Color) == direction {
						break
					}

					// b. Ensure the pawn does not move forward if the square is occupied
					if offset%12 == 0 && gs.board[target] != nil {
						break
					}

					// c. The pawn moves forward one square
					if offset == 12 {

						// handle promotions
						if _, ok := promotionSquares[target]; ok {
							gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{
								QueenPromotion:  {},
								RookPromotion:   {},
								BishopPromotion: {},
								KnightPromotion: {},
							}
						} else {
							gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{Neutral: {}}
						}

						// d. The pawn moves forward two squares
					} else if offset == 24 && !currPiece.HasMoved {
						gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{EnPassantTrigger: {}}

						// e. The pawn captures diagonally
					} else if gs.board[target] != nil {
						gs.attackingSquares[currPiece.Color][target] = struct{}{}
						// handle promotions
						if _, ok := promotionSquares[target]; ok {
							gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{
								QueenPromotion:  {},
								RookPromotion:   {},
								BishopPromotion: {},
								KnightPromotion: {},
							}
						} else {
							gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{Neutral: {}}
						}

						// f. The pawn captures en passant
					} else if target == gs.enPassantSquare && gs.currColor == currPiece.Color { // en passant capture
						gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{EnPassantAttack: {}}
						// g. Mark the pawns diagonal as an attacking square
					} else {
						gs.attackingSquares[currPiece.Color][target] = struct{}{}
					}
				} else {
					gs.allMoves[currPiece.Color][square][target] = map[MoveType]struct{}{Neutral: {}}
					gs.attackingSquares[currPiece.Color][target] = struct{}{}
				}

				// if the move vector is blocked, we should break
				if gs.board[target] != nil {
					break
				}
			}
		}
	}
}

// executeMove executes a move on the board, it does NOT validate the move
func (gs *GameState) executeMove(origin, destination int8, moveType MoveType) {

	changeLog := &HistoryEntry{
		Actions:         []Action{{From: origin, To: destination, hasMoved: gs.board[origin].HasMoved, capture: gs.board[destination]}},
		enPassantSquare: gs.enPassantSquare,
		whiteKingSquare: gs.kingSquares[White],
		blackKingSquare: gs.kingSquares[Black],
	}

	// handle a typical move
	gs.board[destination] = gs.board[origin]
	gs.board[origin] = nil
	gs.board[destination].HasMoved = true

	// set the king square
	if king := gs.board[destination]; king.Type == King {
		gs.kingSquares[king.Color] = destination
	}

	// reset the en passant square
	gs.enPassantSquare = 0

	// handle special moves
	switch moveType {
	// castling
	case WhiteKingSideCastle:
		gs.board[103] = gs.board[105]
		gs.board[105] = nil
		gs.board[103].HasMoved = true
		changeLog.Actions = append(changeLog.Actions, Action{From: 105, To: 103, hasMoved: false, capture: nil})
	case WhiteQueenSideCastle:
		gs.board[101] = gs.board[98]
		gs.board[98] = nil
		gs.board[101].HasMoved = true
		changeLog.Actions = append(changeLog.Actions, Action{From: 98, To: 101, hasMoved: false, capture: nil})
	case BlackKingSideCastle:
		gs.board[19] = gs.board[21]
		gs.board[21] = nil
		gs.board[19].HasMoved = true
		changeLog.Actions = append(changeLog.Actions, Action{From: 21, To: 19, hasMoved: false, capture: nil})
	case BlackQueenSideCastle:
		gs.board[17] = gs.board[14]
		gs.board[14] = nil
		gs.board[17].HasMoved = true
		changeLog.Actions = append(changeLog.Actions, Action{From: 14, To: 17, hasMoved: false, capture: nil})

	// en passant
	case EnPassantAttack:
		square := destination + 12*int8(gs.currColor)
		gs.board[square] = nil
		changeLog.Actions = append(changeLog.Actions, Action{From: square, To: square, hasMoved: false, capture: gs.board[square]})

	case EnPassantTrigger:
		gs.enPassantSquare = destination + 12*int8(gs.currColor)

	// promotions
	case QueenPromotion:
		changeLog.Actions[0].promotionPawn = gs.board[destination]
		gs.board[destination] = &Piece{Type: Queen, Color: gs.currColor, Value: 9, HasMoved: true}
	case RookPromotion:
		changeLog.Actions[0].promotionPawn = gs.board[destination]
		gs.board[destination] = &Piece{Type: Rook, Color: gs.currColor, Value: 5, HasMoved: true}
	case BishopPromotion:
		changeLog.Actions[0].promotionPawn = gs.board[destination]
		gs.board[destination] = &Piece{Type: Bishop, Color: gs.currColor, Value: 3, HasMoved: true}
	case KnightPromotion:
		changeLog.Actions[0].promotionPawn = gs.board[destination]
		gs.board[destination] = &Piece{Type: Knight, Color: gs.currColor, Value: 3, HasMoved: true}
	}

	// update the history
	gs.history = append(gs.history, changeLog)

	// update the current player
	gs.currColor *= -1

	gs.update()
}

// update updates the game state to reflect the current board
func (gs *GameState) update() {

	// reset the valid moves state
	gs.allMoves = map[Color]map[int8]map[int8]map[MoveType]struct{}{White: {}, Black: {}}
	gs.attackingSquares = map[Color]map[int8]struct{}{White: {}, Black: {}}

	for i := range gs.board {
		gs.updateValidMoves(int8(i))
	}

	// check if white can castle on king side
	if gs.whiteKingSide() {
		gs.allMoves[White][102][104] = map[MoveType]struct{}{WhiteKingSideCastle: {}}
	}

	// check if white can castle on queen side
	if gs.whiteQueenSide() {
		gs.allMoves[White][102][100] = map[MoveType]struct{}{WhiteQueenSideCastle: {}}
	}

	// check if black can castle on king side
	if gs.blackKingSide() {
		gs.allMoves[Black][18][20] = map[MoveType]struct{}{BlackKingSideCastle: {}}
	}

	// check if black can castle on queen side
	if gs.blackQueenSide() {
		gs.allMoves[Black][18][16] = map[MoveType]struct{}{BlackQueenSideCastle: {}}
	}
}

// isDangerous returns true if the given square is being attacked by the given color
func (gs *GameState) isDangerous(square int8, attacker Color) bool {
	_, ok := gs.attackingSquares[attacker][square]
	return ok
}

func (gs *GameState) ExecuteRandomMove() bool {
	// get a random move
	for origin, moves := range gs.allMoves[gs.currColor] {
		for destination := range moves {
			for moveType := range moves[destination] {
				gs.executeMove(origin, destination, moveType)
				return true
			}
		}
	}

	return false
}

func (gs *GameState) ExecuteMove(origin, destination int8, moveType MoveType) bool {
	if _, ok := gs.allMoves[gs.currColor][origin][destination][moveType]; !ok {
		return false
	}
	gs.executeMove(origin, destination, moveType)
	return true
}
