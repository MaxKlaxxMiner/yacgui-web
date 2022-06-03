package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

const maxStringLength = FieldCount +
	(4+1)*Height + // indent + linebreaks
	6 + FenMaxBytes

func (board YacBoard) String() string {
	result := make([]byte, 0, maxStringLength)
	for y := 0; y < Height; y++ {
		result = append(result, ' ', ' ', ' ', ' ')
		for x := 0; x < Width; x++ {
			result = append(result, piece.ToChar(board.GetField(FromXY(x, y))))
		}
		result = append(result, '\n')
	}
	result = append(result, "\nFEN: "...)
	result = append(result, board.GetFEN()...)
	return string(result)
}
