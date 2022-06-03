package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"testing"
)

func TestNew(t *testing.T) {
	board := New()
	if board.GetFEN() != "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" {
		t.Errorf("invalid start-FEN: %s", board.GetFEN())
	}

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
		f := board.FieldsF[pos.PToF(i)]
		if f == piece.Blocked {
			t.Errorf("invalid blocked field at pos %d", i)
		}
		if f.String() != checkFields[i:i+1] {
			t.Errorf("invalid field: %s, expected %s", f.String(), checkFields[i:i+1])
		}
	}
}
