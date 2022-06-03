package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"testing"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestGetField(t *testing.T) {
	board := New()
	checkFields := "" +
		"rnbqkbnr" +
		"pppppppp" +
		"........" +
		"........" +
		"........" +
		"........" +
		"PPPPPPPP" +
		"RNBQKBNR"
	for i := range checkFields {
		f := board.GetField(pos.Pos(i))
		if f == piece.Blocked {
			t.Errorf("invalid blocked field at pos %d", i)
		}
		if f.String() != checkFields[i:i+1] {
			t.Errorf("invalid field: %s, expected %s", f.String(), checkFields[i:i+1])
		}
	}
	if board.GetField(pos.Pos(-1)) != piece.Blocked {
		t.Errorf("expected blocked field at pos %d", -1)
	}
	if board.GetField(pos.Pos(len(checkFields))) != piece.Blocked {
		t.Errorf("expected blocked field at pos %d", len(checkFields))
	}
}

func TestSetField(t *testing.T) {
	board, _ := NewFromFEN("rnbk1bnr/pppp1ppp/8/8/2B1P3/8/P2K1PPP/qN1Q2NR b - - 1 7")

	board.SetField(pos.FromChars("d8"), piece.BlackQueen)
	if board.BlackKingPos != pos.FromChars("d8") || board.BlackKingPos == pos.FromChars("e8") {
		t.Errorf("invalid old black king pos")
	}
	board.SetField(pos.FromChars("e8"), piece.BlackKing)
	if board.BlackKingPos == pos.FromChars("d8") || board.BlackKingPos != pos.FromChars("e8") {
		t.Errorf("invalid new black king pos")
	}
	board.SetField(pos.FromChars("e7"), piece.BlackPawn)
	board.SetField(pos.FromChars("c4"), piece.None)
	board.SetField(pos.FromChars("e4"), piece.None)
	board.SetField(pos.FromChars("b2"), piece.WhitePawn)
	board.SetField(pos.FromChars("c2"), piece.WhitePawn)
	board.SetField(pos.FromChars("d2"), piece.WhitePawn)
	board.SetField(pos.FromChars("e2"), piece.WhitePawn)
	board.SetField(pos.FromChars("a1"), piece.WhiteRook)
	board.SetField(pos.FromChars("c1"), piece.WhiteBishop)
	if board.WhiteKingPos != pos.FromChars("d2") || board.WhiteKingPos == pos.FromChars("e1") {
		t.Errorf("invalid old white king pos")
	}
	board.SetField(pos.FromChars("e1"), piece.WhiteKing)
	if board.WhiteKingPos == pos.FromChars("d2") || board.WhiteKingPos != pos.FromChars("e1") {
		t.Errorf("invalid new white king pos")
	}
	board.SetField(pos.FromChars("f1"), piece.WhiteBishop)

	board.WhiteCanCastleKingside = true
	board.WhiteCanCastleQueenside = true
	board.BlackCanCastleKingside = true
	board.BlackCanCastleQueenside = true
	board.HalfmoveClock = 0
	board.MoveNumber = 1
	board.WhiteMove = true

	newBoard := New()
	if board.GetFEN() != newBoard.GetFEN() {
		t.Errorf("repair with setfield failed, result-fen: %s, expected fen: %s", board.GetFEN(), newBoard.GetFEN())
	}

	assertPanic(t, func() {
		board.SetField(-1, piece.BlackQueen)
	})

	assertPanic(t, func() {
		board.SetField(pos.Pos(pos.FieldCount), piece.WhiteQueen)
	})
}
