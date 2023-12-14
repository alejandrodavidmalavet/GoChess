package game

func (gs *GameState) whiteQueenSide() bool {
	// 1. Validate the king
	if piece := gs.board[102]; piece == nil || piece.Type != King || piece.HasMoved {
		return false
	}

	// 2. validate the rook
	if piece := gs.board[98]; piece == nil || piece.Type != Rook || piece.HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[101] != nil || gs.board[100] != nil || gs.board[99] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	if gs.isDangerous(102, Black) || gs.isDangerous(101, Black) || gs.isDangerous(100, Black) {
		return false
	}

	return true
}

func (gs *GameState) whiteKingSide() bool {
	// 1. Validate the king
	if piece := gs.board[102]; piece == nil || piece.Type != King || piece.HasMoved {
		return false
	}

	// 2. validate the rook
	if piece := gs.board[105]; piece == nil || piece.Type != Rook || piece.HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[103] != nil || gs.board[104] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	if gs.isDangerous(102, Black) || gs.isDangerous(103, Black) || gs.isDangerous(104, Black) {
		return false
	}

	return true
}

func (gs *GameState) blackQueenSide() bool {
	// 1. Validate the king
	if piece := gs.board[18]; piece == nil || piece.Type != King || piece.HasMoved {
		return false
	}

	// 2. validate the rook
	if piece := gs.board[14]; piece == nil || piece.Type != Rook || piece.HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[17] != nil || gs.board[16] != nil || gs.board[15] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	if gs.isDangerous(18, White) || gs.isDangerous(17, White) || gs.isDangerous(16, White) {
		return false
	}

	return true

}

func (gs *GameState) blackKingSide() bool {
	// 1. Validate the king
	if piece := gs.board[18]; piece == nil || piece.Type != King || piece.HasMoved {
		return false
	}

	// 2. validate the rook
	if piece := gs.board[21]; piece == nil || piece.Type != Rook || piece.HasMoved {
		return false
	}

	// 3. validate the squares between the king and rook are empty
	if gs.board[19] != nil || gs.board[20] != nil {
		return false
	}

	// 4. validate that the king will not move through or into check
	if gs.isDangerous(18, White) || gs.isDangerous(19, White) || gs.isDangerous(20, White) {
		return false
	}

	return true
}
