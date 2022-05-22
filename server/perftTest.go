package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"strconv"
	"time"
)

func NumFormat(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func MoveCounter(board *yacboard.YacBoard, level int) int {
	var moves [256]yacboard.Move
	moveCount := int(board.GetMovesFast(&moves))
	if level <= 1 {
		return moveCount
	}
	level--
	totalCount := 0
	bi := board.GetBoardInfo()
	for m := 0; m < moveCount; m++ {
		board.DoMove(moves[m])
		totalCount += MoveCounter(board, level)
		board.DoMoveBackward(moves[m], bi)
	}
	return totalCount
}

func perftTest() {
	board := yacboard.New()
	//board, _ := yacboard.NewFromFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
	//board, _ := yacboard.NewFromFEN("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1")
	//board, _ := yacboard.NewFromFEN("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1")
	//board, _ := yacboard.NewFromFEN("r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1")
	//board, _ := yacboard.NewFromFEN("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8")
	//board, _ := yacboard.NewFromFEN("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10")

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

	for level := 1; level < 7; level++ {
		fmt.Print("Level: ", level, " Count: ")
		tim := time.Now()
		fmt.Printf("%s (%s ms)\n", NumFormat(int64(MoveCounter(&board, level))), NumFormat(time.Since(tim).Milliseconds()))
	}
}
