package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"testing"
)

func TestClear(t *testing.T) {
	board, _ := NewFromFEN("rnbk1bnr/pppp1ppp/8/8/2B1P3/8/P2K1PPP/qN1Q2NR b - - 1 7")

	board.Clear()

	clearFen := board.GetFEN()
	if clearFen != "8/8/8/8/8/8/8/8 w - - 0 1" {
		t.Errorf("invalid clear-FEN: %s, expected-FEN: %s", clearFen, "8/8/8/8/8/8/8/8 w - - 0 1")
	}

	cleanBoard := YacBoard{MoveNumber: 1, WhiteKingPos: -1, BlackKingPos: -1, EnPassantPos: -1, WhiteMove: true}
	if cleanBoard != board {
		t.Errorf("invalid clean board")
	}

	board.SetField(pos.FromChars("e4"), piece.WhitePawn)
	if cleanBoard == board {
		t.Errorf("unexpected comparing")
	}

	board.Fields[pos.FromChars("e4")] = piece.None
	if cleanBoard != board {
		t.Errorf("invalid clean board")
	}
}
