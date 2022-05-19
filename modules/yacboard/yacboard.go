package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

type YacBoard struct {
	Fields [boardsize.FieldCount]piece.Piece

	HalfmoveClock int
	MoveNumber    int
	WhiteKingPos  pos.Pos
	BlackKingPos  pos.Pos
	EnPassantPos  pos.Pos

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
