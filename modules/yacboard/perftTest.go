package yacboard

import (
	"fmt"
	"strconv"
	"time"
)

func numFormat(n int64) string {
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

func moveCounter(board *YacBoard, level int) int {
	var moves [256]Move
	moveCount := int(board.GetMovesFast(&moves))
	if level <= 1 {
		return moveCount
	}
	level--
	totalCount := 0
	bi := board.GetBoardInfo()
	for m := 0; m < moveCount; m++ {
		board.DoMove(moves[m])
		totalCount += moveCounter(board, level)
		board.DoMoveBackward(moves[m], bi)
	}
	return totalCount
}

func perftTestFEN(fen string, nodeCounter []int64) {
	board, err := NewFromFEN(fen)
	if err != nil {
		panic(err)
	}
	fmt.Println()
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

	for level := 1; level <= len(nodeCounter); level++ {
		fmt.Print("Level: ", level, " / ", len(nodeCounter), " Nodes: ")
		tim := time.Now()
		count := int64(moveCounter(&board, level))
		fmt.Printf("%s (%s ms)", numFormat(count), numFormat(time.Since(tim).Milliseconds()))
		if count == nodeCounter[len(nodeCounter)-level] {
			fmt.Println(" [ok]")
		} else {
			fmt.Printf(" [FAIL] %d != %d\n", count, nodeCounter[len(nodeCounter)-level])
			panic("perft fail")
		}
	}
}

func PerftTest(trim int) {
	// source: https://www.chessprogramming.org/Perft_Results
	perftTestFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", ([]int64{3195901860, 119060324, 4865609, 197281, 8902, 400, 20})[trim:])
	perftTestFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", ([]int64{8031647685, 193690690, 4085603, 97862, 2039, 48})[trim:])
	perftTestFEN("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1", ([]int64{3009794393, 178633661, 11030083, 674624, 43238, 2812, 191, 14})[trim:])
	perftTestFEN("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1", ([]int64{706045033, 15833292, 422333, 9467, 264, 6})[trim:])
	perftTestFEN("r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1", ([]int64{706045033, 15833292, 422333, 9467, 264, 6})[trim:])
	perftTestFEN("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8", ([]int64{3048196529, 89941194, 2103487, 62379, 1486, 44})[trim:])
	perftTestFEN("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10", ([]int64{6923051137, 164075551, 3894594, 89890, 2079, 46})[trim:])
}
