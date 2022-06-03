package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) Clear() {
	for i := 0; i < len(board.FieldsF); i++ {
		board.FieldsF[i] = 0
	}
	for x := 0; x < pos.WidthF; x++ {
		board.FieldsF[x] = piece.Blocked
		board.FieldsF[x+(pos.HeightF-1)*pos.WidthF] = piece.Blocked
	}
	for y := 0; y < pos.HeightF; y++ {
		board.FieldsF[y*pos.WidthF] = piece.Blocked
		board.FieldsF[y*pos.WidthF+pos.WidthF-1] = piece.Blocked
	}

	board.HalfmoveClock = 0
	board.MoveNumber = 1
	board.WhiteKingPosF = 0
	board.BlackKingPosF = 0
	board.EnPassantPosF = 0

	board.WhiteMove = true
	board.WhiteCanCastleKingside = false
	board.WhiteCanCastleQueenside = false
	board.BlackCanCastleKingside = false
	board.BlackCanCastleQueenside = false
}
