package yacboard

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"unicode"
)

func (board *YacBoard) GetFEN() string {
	result := make([]byte, 0, 64)

	for y := 0; y < boardsize.Height; y++ {
		result = append(result, '/')
		for x := 0; x < boardsize.Width; x++ {
			c := piece.ToChar(board.GetField(pos.FromXY(x, y)))
			if c == '.' {
				if unicode.IsDigit(rune(result[len(result)-1])) {
					result[len(result)-1]++
				} else {
					result = append(result, '1')
				}
			} else {
				result = append(result, c)
			}
		}
	}

	if board.WhiteMove {
		result = append(result, ' ', 'w', ' ')
	} else {
		result = append(result, ' ', 'b', ' ')
	}

	if board.WhiteCanCastleKingside {
		result = append(result, 'K')
	}
	if board.WhiteCanCastleQueenside {
		result = append(result, 'Q')
	}
	if board.BlackCanCastleKingside {
		result = append(result, 'k')
	}
	if board.BlackCanCastleQueenside {
		result = append(result, 'q')
	}
	if !board.WhiteCanCastleKingside && !board.WhiteCanCastleQueenside && !board.BlackCanCastleKingside && !board.BlackCanCastleQueenside {
		result = append(result, '-')
	}

	result = append(result, fmt.Sprintf(" %s %d %d", board.EnPassantPos, board.HalfmoveClock, board.MoveNumber)...)

	return string(result[1:])
}
