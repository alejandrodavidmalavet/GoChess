package main

import (
	"fmt"

	"github.com/alejandrodavidmalavet/GoChess/internal/game"
)

func main() {
	gs := game.NewGame()
	for {
		var choice int
		fmt.Printf("\n%v's turn\n", gs.CurrentPlayer())
		fmt.Print("\n[0] Random Move, [1] Custom Move, [2] Undo Move: ")
		fmt.Scanln(&choice)

		if choice == 0 {
			if ok := gs.ExecuteRandomMove(); !ok {
				fmt.Println("Invalid move")
			}
		} else if choice == 1 {
			var from, to int8
			fmt.Print("From: ")
			fmt.Scanln(&from)
			fmt.Print("To: ")
			fmt.Scanln(&to)
			fmt.Print("[0] Neutral\n[1] EnPassantAttack\n[2] EnPassantTrigger\n[3] WhiteKingSideCastle\n[4] WhiteQueenSideCastle\n[5] BlackKingSideCastle\n[6] BlackQueenSideCastle\n[7] QueenPromotion\n[8] RookPromotion\n[9] BishopPromotion\n[10] KnightPromotion\nMoveType: ")
			fmt.Scanln(&choice)

			if ok := gs.ExecuteMove(from, to, game.MoveType(choice)); !ok {
				fmt.Println("Invalid move")
			}
		} else {
			gs.Undo()
		}

	}

}
