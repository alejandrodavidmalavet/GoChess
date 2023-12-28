package game

type Color int8

const (
	White Color = 1
	Black Color = -1
)

type MoveType int8

const (
	Neutral MoveType = iota

	// en passant moves
	EnPassantAttack
	EnPassantTrigger

	// castling moves
	WhiteKingSideCastle
	WhiteQueenSideCastle
	BlackKingSideCastle
	BlackQueenSideCastle

	// promotion moves
	QueenPromotion
	RookPromotion
	BishopPromotion
	KnightPromotion
)

type Type int

const (
	King Type = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)
