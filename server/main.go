package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
)

func main() {
	board := yacboard.New()

	fmt.Println(board)

	moves := board.GetMoves()
	fmt.Printf("\nMoves [%d]: ", len(moves))
	for i, m := range moves {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(m)
	}
	fmt.Println()
}
