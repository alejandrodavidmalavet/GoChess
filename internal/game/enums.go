package game

type Color int8

const (
	White Color = 1
	Black Color = -1
)

type MoveType int8

const (
	Neutral MoveType = 0

	// en passant moves
	EnPassantAttack  MoveType = 1
	EnPassantTrigger MoveType = 2

	// castling moves
	WhiteKingSideCastle  MoveType = 3
	WhiteQueenSideCastle MoveType = 4
	BlackKingSideCastle  MoveType = 5
	BlackQueenSideCastle MoveType = 6

	// promotion moves
	QueenPromotion  MoveType = 7
	RookPromotion   MoveType = 8
	BishopPromotion MoveType = 9
	KnightPromotion MoveType = 10
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
