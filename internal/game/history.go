package game

type Action struct {
	From          int8
	To            int8
	hasMoved      bool
	capture       *Piece
	promotionPawn *Piece
}

type HistoryEntry struct {
	Actions         []Action
	enPassantSquare int8
	whiteKingSquare int8
	blackKingSquare int8
}

// Undo the latest move
func (gs *GameState) Undo() {
	// if there is no history, do nothing
	if len(gs.history) == 0 {
		return
	}

	// pop the last entry from the history
	var entry *HistoryEntry
	entry, gs.history = gs.history[len(gs.history)-1], gs.history[:len(gs.history)-1]

	// undo the actions
	for _, change := range entry.Actions {
		if change.promotionPawn != nil {
			gs.board[change.To] = change.promotionPawn
		}
		gs.board[change.From] = gs.board[change.To]
		gs.board[change.To] = change.capture
		gs.board[change.From].HasMoved = change.hasMoved
	}

	// update the state variables
	gs.enPassantSquare = entry.enPassantSquare
	gs.kingSquares[White] = entry.whiteKingSquare
	gs.kingSquares[Black] = entry.blackKingSquare

	gs.currColor *= -1

	gs.update()
}