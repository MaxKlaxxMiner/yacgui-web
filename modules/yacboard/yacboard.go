package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

type YacBoard struct {
	FieldsF [FieldCountF]piece.Piece

	HalfmoveClock int
	MoveNumber    int
	WhiteKingPosF Pos
	BlackKingPosF Pos
	EnPassantPosF Pos

	WhiteMove               bool
	WhiteCanCastleKingside  bool
	WhiteCanCastleQueenside bool
	BlackCanCastleKingside  bool
	BlackCanCastleQueenside bool
}

func New() YacBoard {
	board, _ := NewFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return board
}

func NewFromFEN(fen string) (YacBoard, error) {
	var board YacBoard
	err := board.SetFEN(fen)
	return board, err
}
