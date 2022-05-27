package yacboard

import (
	"testing"
)

func TestString(t *testing.T) {
	board := New()
	const cmp = "    rnbqkbnr\n    pppppppp\n    ........\n    ........\n    ........\n    ........\n    PPPPPPPP\n    RNBQKBNR\n\nFEN: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	if board.String() != cmp {
		t.Errorf("string() changed?")
	}
}
