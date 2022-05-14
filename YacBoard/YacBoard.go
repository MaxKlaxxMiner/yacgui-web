package YacBoard

import (
	"strconv"
	"unsafe"
)

const (
	Width  = 8
	Height = 8
)

type YacBoard struct {
	Fields [Width * Height]Piece

	HalfmoveClock int
	MoveNumber    int
	WhiteKingPos  int
	BlackKingPos  int
	EnPassantPos  int

	WhiteMove               bool
	WhiteCanCastleKingside  bool
	WhiteCanCastleQueenside bool
	BlackCanCastleKingside  bool
	BlackCanCastleQueenside bool
}

type byte64 struct {
	u0 uint64
	u1 uint64
	u2 uint64
	u3 uint64
	u4 uint64
	u5 uint64
	u6 uint64
	u7 uint64
}

func (board *YacBoard) Clear() {
	if strconv.IntSize == 64 && len(board.Fields) == 64 {
		ptr := (*byte64)(unsafe.Pointer(&board.Fields))
		ptr.u0 = 0
		ptr.u1 = 0
		ptr.u2 = 0
		ptr.u3 = 0
		ptr.u4 = 0
		ptr.u5 = 0
		ptr.u6 = 0
		ptr.u7 = 0
	} else {
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

func (board *YacBoard) SetField(pos int, piece Piece) {
	if uint(pos) >= Width*Height {
		panic("argument out of range")
		return
	}
	board.Fields[pos] = piece

	if piece&King == King {
		if piece == WhiteKing {
			board.WhiteKingPos = pos
		} else {
			board.BlackKingPos = pos
		}
	}
}

func (board *YacBoard) GetField(pos int) Piece {
	if uint(pos) >= Width*Height {
		return Blocked
	}
	return board.Fields[pos]
}

func (board *YacBoard) GetFieldXY(x, y int) Piece {
	if uint(x) >= Width || uint(y) >= Height {
		return Blocked
	}
	return board.GetField(x + y*Width)
}

func (board *YacBoard) GetBoardInfo() BoardInfo {
	result := BoardInfo(uint8(int8(board.EnPassantPos))) | BoardInfo(uint(board.HalfmoveClock)<<16)

	if board.WhiteCanCastleKingside {
		result |= WhiteCanCastleKingside
	}
	if board.WhiteCanCastleQueenside {
		result |= WhiteCanCastleQueenside
	}
	if board.BlackCanCastleKingside {
		result |= BlackCanCastleKingside
	}
	if board.BlackCanCastleQueenside {
		result |= BlackCanCastleQueenside
	}

	return result
}

func (board *YacBoard) SetBoardInfo(boardInfo BoardInfo) {
	board.EnPassantPos = int(int8(uint8(boardInfo & EnPassantMask)))
	board.WhiteCanCastleKingside = (boardInfo & WhiteCanCastleKingside) != BoardInfoNone
	board.WhiteCanCastleQueenside = (boardInfo & WhiteCanCastleQueenside) != BoardInfoNone
	board.BlackCanCastleKingside = (boardInfo & BlackCanCastleKingside) != BoardInfoNone
	board.BlackCanCastleQueenside = (boardInfo & BlackCanCastleQueenside) != BoardInfoNone
	board.HalfmoveClock = (int)(boardInfo&HalfmoveClockMask) >> 16
}

func (board YacBoard) String() string {
	result := make([]byte, 0, Width*Height+(4+1)*Height+6+FenMaxLength)
	for y := 0; y < Height; y++ {
		result = append(result, ' ', ' ', ' ', ' ')
		for x := 0; x < Width; x++ {
			result = append(result, PieceToChar(board.GetFieldXY(x, y)))
		}
		result = append(result, '\n')
	}
	result = append(result, "\nFEN: "...)
	result = append(result, board.GetFEN()...)
	return string(result)
}
