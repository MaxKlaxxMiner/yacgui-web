package yacboard

import (
	"strconv"
	"unsafe"
)

func (board *YacBoard) Clear() {
	if //goland:noinspection GoBoolExpressions
	strconv.IntSize == 64 && len(board.Fields) == 64 {
		ptr := (*struct{ u0, u1, u2, u3, u4, u5, u6, u7 uint64 })(unsafe.Pointer(&board.Fields))
		ptr.u0 = 0
		ptr.u1 = 0
		ptr.u2 = 0
		ptr.u3 = 0
		ptr.u4 = 0
		ptr.u5 = 0
		ptr.u6 = 0
		ptr.u7 = 0
	} else {
		// slower fallback version
		for i := 0; i < len(board.Fields); i++ {
			board.Fields[i] = 0
		}
	}

	board.HalfmoveClock = 0
	board.MoveNumber = 1
	board.WhiteKingPos = -1
	board.BlackKingPos = -1
	board.EnPassantPos = -1

	board.WhiteMove = true
	board.WhiteCanCastleKingside = false
	board.WhiteCanCastleQueenside = false
	board.BlackCanCastleKingside = false
	board.BlackCanCastleQueenside = false
}
