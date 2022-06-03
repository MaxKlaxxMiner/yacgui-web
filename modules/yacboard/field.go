package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) GetField(pos Pos) piece.Piece {
	if uint(pos) >= FieldCount {
		return piece.Blocked
	}
	return board.FieldsF[PToFp(pos)]
}

func (board *YacBoard) SetField(pos Pos, p piece.Piece) {
	if uint(pos) >= FieldCount {
		panic("argument out of range")
	}
	board.FieldsF[PToFp(pos)] = p

	if p&piece.King == piece.King {
		if p == piece.WhiteKing {
			board.WhiteKingPos = pos
		} else {
			board.BlackKingPos = pos
		}
	}
}
