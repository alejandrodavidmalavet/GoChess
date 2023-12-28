package main

import (
	"fmt"

	"github.com/alejandrodavidmalavet/GoChess/internal/game"
)

type Choice int

const (
	RandomMove Choice = iota
	CustomMove
	UndoMove
	BestMove
	AIVsAI
)

func main() {
	gs := game.NewGame()
	for {
		gs.PrettyPrint()

		var c Choice
		fmt.Printf("\n%v's turn\n", gs.CurrentPlayer())
		fmt.Print("\n[", RandomMove, "] Random Move\n",
			"[", CustomMove, "] Custom Move\n",
			"[", UndoMove, "] Undo Move\n",
			"[", BestMove, "] Best Move\n",
			"[", AIVsAI, "] AI v AI\n",
			"Choice: ")
		fmt.Scanln(&c)

		switch c {
		case RandomMove:
			if ok := gs.ExecuteRandomMove(); !ok {
				fmt.Println("Invalid move")
			}
		case CustomMove:
			var from, to int8
			fmt.Print("From: ")
			fmt.Scanln(&from)
			fmt.Print("To: ")
			fmt.Scanln(&to)

			var mt game.MoveType
			fmt.Print("[", game.Neutral, "] Neutral\n",
				"[", game.EnPassantAttack, "] EnPassantAttack\n",
				"[", game.EnPassantPrimer, "] EnPassantPrimer\n",
				"[", game.WhiteKingSideCastle, "] WhiteKingSideCastle\n",
				"[", game.WhiteQueenSideCastle, "] WhiteQueenSideCastle\n",
				"[", game.BlackKingSideCastle, "] BlackKingSideCastle\n",
				"[", game.BlackQueenSideCastle, "] BlackQueenSideCastle\n",
				"[", game.QueenPromotion, "] QueenPromotion\n",
				"[", game.RookPromotion, "] RookPromotion\n",
				"[", game.BishopPromotion, "] BishopPromotion\n",
				"[", game.KnightPromotion, "] KnightPromotion\n",
				"MoveType: ")
			fmt.Scanln(&mt)

			if ok := gs.ExecuteMove(from, to, mt); !ok {
				fmt.Println("Invalid move")
			}
		case UndoMove:
			gs.Undo()
		case BestMove:
			var depth int8
			fmt.Print("Depth: ")
			fmt.Scanln(&depth)
			gs.ExecuteBestMove(depth)
		case AIVsAI:
			var depth int8
			fmt.Print("Depth: ")
			fmt.Scanln(&depth)
			for {
				gs.ExecuteBestMove(depth)
				gs.PrettyPrint()
			}
		default:
			fmt.Println("Invalid choice")
		}
	}

}
