package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

const maxStringLength = boardsize.FieldCount +
	(4+1)*boardsize.Height + // indent + linebreaks
	6 + boardsize.FenMaxBytes

func (board YacBoard) String() string {
	result := make([]byte, 0, maxStringLength)
	for y := 0; y < boardsize.Height; y++ {
		result = append(result, ' ', ' ', ' ', ' ')
		for x := 0; x < boardsize.Width; x++ {
			result = append(result, piece.ToChar(board.GetField(pos.FromXY(x, y))))
		}
		result = append(result, '\n')
	}
	result = append(result, "\nFEN: "...)
	result = append(result, board.GetFEN()...)
	return string(result)
}
