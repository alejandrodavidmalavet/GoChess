package main

import (
	"fmt"

	"github.com/alejandrodavidmalavet/GoChess/internal/game"
)

func main() {
	gs := game.NewGame()
	for {
		gs.PrettyPrint()

		var choice int
		fmt.Printf("\n%v's turn\n", gs.CurrentPlayer())
		fmt.Print("\n[0] Random Move, [1] Custom Move, [2] Undo Move, [3] Best Move ")
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
		} else if choice == 2 {
			gs.Undo(true)
		} else if choice == 3 {
			var depth int8
			fmt.Print("Depth: ")
			fmt.Scanln(&depth)
			gs.ExecuteBestMove(depth)
		} else {
			fmt.Println("Invalid choice")
		}
	}

}

// func main() {

// 	gs := game.NewGame()
// 	count := 0
// 	defer func(gs *game.GameState) {
// 		if r := recover(); r != nil {
// 			gs.PrettyPrint()
// 			panic(gs)
// 		}
// 	}(gs)
// 	simCount := 0
// 	totalMoves := 0
// 	startTime := time.Now()
// 	for {
// 		// gs.PrettyPrint()
// 		// fmt.Printf("\n%v's turn\n", gs.CurrentPlayer())

// 		if ok := gs.ExecuteRandomMove(); !ok || count > 500 {
// 			simCount++
// 			// fmt.Println("Game Over!")
// 			if simCount%100 == 0 {
// 				gs.PrettyPrint()
// 			}
// 			for gs.Undo(true) {
// 			}
// 			// gs.PrettyPrint()
// 			// fmt.Println("Game Reset!")
// 			totalMoves += count
// 			count = 0
// 			if simCount%100 == 0 {
// 				gs.PrettyPrint()
// 				println("\n# of games simulated & undone:", simCount)
// 				fmt.Printf("Time elapsed: %f\n", time.Since(startTime).Seconds())
// 				fmt.Printf("Moves per second: %f\n", float64(totalMoves)/time.Since(startTime).Seconds())
// 			}

// 		} else {
// 			count++
// 		}
// 	}

// }
