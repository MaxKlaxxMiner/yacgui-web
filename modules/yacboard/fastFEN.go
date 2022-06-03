package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) GetFastFEN(buf []byte, ofs int) int {
	p := 0
	gap := 0
	for i := 0; i < FieldCount; i++ {
		field := board.FieldsF[PToF(i)]
		if field == piece.None {
			gap++
			continue
		}
		if gap > 0 {
			buf[ofs+p] = byte(uint(gap))
			p++
			gap = 0
		}
		buf[ofs+p] = byte(field)
		p++
	}
	if gap > 0 {
		buf[ofs+p] = byte(uint(gap))
		p++
	}

	binfo := uint32(board.GetBoardInfo())
	buf[ofs+p] = byte(binfo)
	buf[ofs+p+1] = byte(binfo >> 8)
	buf[ofs+p+2] = byte(binfo >> 16)
	buf[ofs+p+3] = byte(binfo >> 24)
	buf[ofs+p+4] = byte(board.MoveNumber)
	buf[ofs+p+5] = byte(board.MoveNumber >> 8)
	p += 6

	return p
}

func (board *YacBoard) SetFastFEN(buf []byte, ofs int) int {
	p := 0
	var b byte
	for i := 0; i < FieldCount; i++ {
		b = buf[ofs+p]
		p++
		if b < 64 { // gap?
			board.FieldsF[PToF(i)] = piece.None
			for b > 1 {
				i++
				board.FieldsF[PToF(i)] = piece.None
				b--
			}
			continue
		}
		board.FieldsF[PToF(i)] = piece.Piece(b)
		if piece.Piece(b)&piece.King != piece.None {
			if piece.Piece(b) == piece.WhiteKing {
				board.WhiteKingPosF = PToFp(Pos(i))
			} else {
				board.BlackKingPosF = PToFp(Pos(i))
			}
		}
	}
	var binfo = uint32(buf[ofs+p]) | uint32(buf[ofs+p+1])<<8 | uint32(buf[ofs+p+2])<<16 | uint32(buf[ofs+p+3])<<24
	board.SetBoardInfo(BoardInfo(binfo))
	board.MoveNumber = int(buf[ofs+p+4]) | int(buf[ofs+p+5])<<8
	p += 6
	return 0
}
