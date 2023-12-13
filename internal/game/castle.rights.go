package game

func (gs *GameState) whiteQueenSide() bool {
	// 1. Validate the king
	if gs.board[102] == nil || gs.board[102].Type != King || gs.board[102].HasMoved {
		return false
	}

	// 2. validate the rook
	if gs.board[98] == nil || gs.board[98].Type != Rook || gs.board[98].HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[101] != nil || gs.board[100] != nil || gs.board[99] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	// todo

	return true
}

func (gs *GameState) whiteKingSide() bool {
	// 1. Validate the king
	if gs.board[102] == nil || gs.board[102].Type != King || gs.board[102].HasMoved {
		return false
	}

	// 2. validate the rook
	if gs.board[105] == nil || gs.board[105].Type != Rook || gs.board[105].HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[103] != nil || gs.board[104] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	// todo

	return true
}

func (gs *GameState) blackQueenSide() bool {
	// 1. Validate the king
	if gs.board[18] == nil || gs.board[18].Type != King || gs.board[18].HasMoved {
		return false
	}

	// 2. validate the rook
	if gs.board[14] == nil || gs.board[14].Type != Rook || gs.board[14].HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[15] != nil || gs.board[16] != nil || gs.board[17] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	// todo

	return true
}

func (gs *GameState) blackKingSide() bool {
	// 1. Validate the king
	if gs.board[18] == nil || gs.board[18].Type != King || gs.board[18].HasMoved {
		return false
	}

	// 2. validate the rook
	if gs.board[21] == nil || gs.board[21].Type != Rook || gs.board[21].HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[19] != nil || gs.board[20] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	// todo

	return true
}
