package yacboard

import (
	"strings"
	"testing"
)

func fenError(t *testing.T, board YacBoard, fenStr, expectedError string) {
	err := board.SetFEN(fenStr)
	if err == nil {
		t.Errorf("error expected: \"%s\" in FEN: \"%s\"", expectedError, fenStr)
	}
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("wrong error: \"%s\", not contains: \"%s\" in FEN: \"%s\"", err.Error(), expectedError, fenStr)
	}
}

func fenOk(t *testing.T, board YacBoard, fenStr string) {
	err := board.SetFEN(fenStr)
	if err != nil {
		t.Errorf("unexpected error: \"%s\" in FEN: \"%s\"", err.Error(), fenStr)
	}
}

func TestSetFEN(t *testing.T) {
	board := New()

	testFen := board.GetFEN()

	fenError(t, board, testFen+testFen+testFen, "too long")

	fenError(t, board, testFen+" 0", "too many elements")

	fenError(t, board, testFen[:len(testFen)-2], "too few elements")

	fenError(t, board, "8/"+testFen, "ranks")

	fenError(t, board, "x"+testFen[1:], "unknown char")

	fenError(t, board, "2"+testFen, "files")

	fenError(t, board, testFen[:strings.Index(testFen, " w ")]+" x "+testFen[strings.Index(testFen, " w ")+3:], "unknown move color")

	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" - "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" q "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" k "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" kq "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Q "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Qq "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Qk "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Qkq "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" K "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Kq "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Kk "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" Kkq "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" KQ "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" KQq "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenOk(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" KQk "+testFen[strings.Index(testFen, " KQkq ")+6:])
	fenError(t, board, testFen[:strings.Index(testFen, " KQkq ")]+" xD "+testFen[strings.Index(testFen, " KQkq ")+6:], "unknown castling")

	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq a2 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a2 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a3 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq a4 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a4 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq a6 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq a7 0 1", "invalid en passant")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq a7 0 1", "invalid en passant")

	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - -1 1", "halfmove-counter")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 10000 1", "halfmove-counter")
	fenOk(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 9999 1")

	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 -1", "movenumber")
	fenError(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 10000", "movenumber")
	fenOk(t, board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 9999")
}
