package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"testing"
)

func testGetFENCheck(t *testing.T, board *YacBoard, expectedFEN string) {
	if board.GetFEN() != expectedFEN {
		t.Errorf("invalid FEN: %s, expected: %s", board.GetFEN(), expectedFEN)
	}
}

func TestGetFEN(t *testing.T) {
	board := New()
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	board.WhiteMove = false
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")

	board.WhiteCanCastleKingside = false
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Qkq - 0 1")

	board.WhiteCanCastleQueenside = false
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b kq - 0 1")

	board.BlackCanCastleKingside = false
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b q - 0 1")

	board.BlackCanCastleQueenside = false
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - - 0 1")

	board.HalfmoveClock = 33
	board.MoveNumber = 40
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b - - 33 40")

	board.SetFEN("rnbqkbnr/ppp2ppp/4p3/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3")
	testGetFENCheck(t, &board, "rnbqkbnr/ppp2ppp/4p3/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3")
	if pos.FToPp(board.EnPassantPosF).String() != "d6" {
		t.Errorf("invalid en passant-pos: %s", pos.FToPp(board.EnPassantPosF))
	}

	board.SetFEN("8/8/8/3KB3/3N4/8/8/7k w - - 0 1")
	testGetFENCheck(t, &board, "8/8/8/3KB3/3N4/8/8/7k w - - 0 1")
}
