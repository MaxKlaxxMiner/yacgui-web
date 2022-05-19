package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) GetField(pos pos.Pos) piece.Piece {
	if uint(pos) >= boardsize.FieldCount {
		return piece.Blocked
	}
	return board.Fields[pos]
}

func (board *YacBoard) SetField(pos pos.Pos, p piece.Piece) {
	if uint(pos) >= boardsize.FieldCount {
		panic("argument out of range")
		return
	}
	board.Fields[pos] = p

	if p&piece.King == piece.King {
		if p == piece.WhiteKing {
			board.WhiteKingPos = pos
		} else {
			board.BlackKingPos = pos
		}
	}
}
