package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"testing"
)

func testDoMoveCheck(t *testing.T, board *YacBoard, move Move, moveStr string) {
	if !move.IsValid(nil) {
		t.Errorf("move without boardcheck should valid: %s", move)
	}
	if !move.IsValid(board) {
		t.Errorf("move should valid: %s at pos: %s", move, board.GetFEN())
	}
	board.DoMove(move)
	if move.IsValid(board) {
		t.Errorf("invalid move marked as valid: %s at pos: %s", move, board.GetFEN())
	}
	str := move.String()
	if str != moveStr {
		t.Errorf("invalid move-string: \"%s\", expected: \"%s\"", str, moveStr)
	}
}

func TestMove(t *testing.T) {
	board := New()

	m0 := Move{}
	if m0.IsValid(nil) || m0.IsValid(&board) {
		t.Errorf("invalid move marked as valid: %s", m0)
	}
	str := m0.String()
	if str != "-" {
		t.Errorf("invalid move-string: \"%s\", expected: \"%s\"", str, "-")
	}

	m2 := Move{FromPos: byte(pos.FromChars("e7")), ToPos: byte(pos.FromChars("e5"))}
	if !m2.IsValid(nil) {
		t.Errorf("move should valid: %s", m2)
	}
	if m2.IsValid(&board) {
		t.Errorf("invalid move marked as valid: %s", m2)
	}

	m1 := Move{FromPos: byte(pos.FromChars("e2")), ToPos: byte(pos.FromChars("e4"))}
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	testDoMoveCheck(t, &board, m1, "e2-e4")
	testGetFENCheck(t, &board, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1")
	testDoMoveCheck(t, &board, m2, "e7-e5")
	testGetFENCheck(t, &board, "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2")

	table := []struct {
		str     string
		move    Move
		nextFEN string
	}{
		{"d2-d4", Move{FromPos: byte(pos.FromChars("d2")), ToPos: byte(pos.FromChars("d4"))}, "rnbqkbnr/pppp1ppp/8/4p3/3PP3/8/PPP2PPP/RNBQKBNR b KQkq - 0 2"},
		{"e5xd4 (xP)", Move{FromPos: byte(pos.FromChars("e5")), ToPos: byte(pos.FromChars("d4")), CapturePiece: piece.WhitePawn}, "rnbqkbnr/pppp1ppp/8/8/3pP3/8/PPP2PPP/RNBQKBNR w KQkq - 0 3"},
		{"c2-c4", Move{FromPos: byte(pos.FromChars("c2")), ToPos: byte(pos.FromChars("c4"))}, "rnbqkbnr/pppp1ppp/8/8/2PpP3/8/PP3PPP/RNBQKBNR b KQkq c3 0 3"},
		{"d4-c3", Move{FromPos: byte(pos.FromChars("d4")), ToPos: byte(pos.FromChars("c3"))}, "rnbqkbnr/pppp1ppp/8/8/4P3/2p5/PP3PPP/RNBQKBNR w KQkq - 0 4"}, // no "CapturePiece" by en passant-moves
		{"f1-c4", Move{FromPos: byte(pos.FromChars("f1")), ToPos: byte(pos.FromChars("c4"))}, "rnbqkbnr/pppp1ppp/8/8/2B1P3/2p5/PP3PPP/RNBQK1NR b KQkq - 1 4"},
		{"c3xb2 (xP)", Move{FromPos: byte(pos.FromChars("c3")), ToPos: byte(pos.FromChars("b2")), CapturePiece: piece.WhitePawn}, "rnbqkbnr/pppp1ppp/8/8/2B1P3/8/Pp3PPP/RNBQK1NR w KQkq - 0 5"},
		{"c1-g5", Move{FromPos: byte(pos.FromChars("c1")), ToPos: byte(pos.FromChars("g5"))}, "rnbqkbnr/pppp1ppp/8/6B1/2B1P3/8/Pp3PPP/RN1QK1NR b KQkq - 1 5"},
		{"b2xa1->q (xR)", Move{FromPos: byte(pos.FromChars("b2")), ToPos: byte(pos.FromChars("a1")), CapturePiece: piece.WhiteRook, PromotionPiece: piece.BlackQueen}, "rnbqkbnr/pppp1ppp/8/6B1/2B1P3/8/P4PPP/qN1QK1NR w Kkq - 0 6"},
		{"g5xd8 (xq)", Move{FromPos: byte(pos.FromChars("g5")), ToPos: byte(pos.FromChars("d8")), CapturePiece: piece.BlackQueen}, "rnbBkbnr/pppp1ppp/8/8/2B1P3/8/P4PPP/qN1QK1NR b Kkq - 0 6"},
		{"e8xd8 (xB)", Move{FromPos: byte(pos.FromChars("e8")), ToPos: byte(pos.FromChars("d8")), CapturePiece: piece.WhiteBishop}, "rnbk1bnr/pppp1ppp/8/8/2B1P3/8/P4PPP/qN1QK1NR w K - 0 7"},
	}

	for _, m := range table {
		testDoMoveCheck(t, &board, m.move, m.str)
		testGetFENCheck(t, &board, m.nextFEN)
	}
}
