package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) IsChecked() bool {
	if board.WhiteMove {
		return board.isChecked(board.WhiteKingPosF, piece.Black)
	} else {
		return board.isChecked(board.BlackKingPosF, piece.WhiteKing)
	}
}

func (board *YacBoard) isChecked(posF PosF, checkerColor piece.Piece) bool {
	posX := int(posF%WidthF - 1)
	posY := int(posF/WidthF - 1)
	pos := Pos(FToP(posF))

	// --- check pawn and king ---
	if checkerColor == piece.White {
		if posX > 0 {
			if posY > 0 && pos-(Width+1) == Pos(FToP(board.WhiteKingPosF)) {
				return true
			}
			if pos-1 == Pos(FToP(board.WhiteKingPosF)) {
				return true
			}
			if posY < Height-1 && (pos+(Width-1) == Pos(FToP(board.WhiteKingPosF)) || board.FieldsF[PToFp(pos+(Width-1))] == piece.WhitePawn) {
				return true
			}
		}
		if posX < Width-1 {
			if posY > 0 && pos-(Width-1) == Pos(FToP(board.WhiteKingPosF)) {
				return true
			}
			if pos+1 == Pos(FToP(board.WhiteKingPosF)) {
				return true
			}
			if posY < Height-1 && (pos+(Width+1) == Pos(FToP(board.WhiteKingPosF)) || board.FieldsF[PToFp(pos+(Width+1))] == piece.WhitePawn) {
				return true
			}
		}
		if posY > 0 && pos-Width == Pos(FToP(board.WhiteKingPosF)) {
			return true
		}
		if posY < Height-1 && pos+Width == Pos(FToP(board.WhiteKingPosF)) {
			return true
		}
	} else {
		if posX > 0 {
			if posY > 0 && (pos-(Width+1) == Pos(FToP(board.BlackKingPosF)) || board.FieldsF[PToFp(pos-(Width+1))] == piece.BlackPawn) {
				return true
			}
			if pos-1 == Pos(FToP(board.BlackKingPosF)) {
				return true
			}
			if posY < Height-1 && pos+(Width-1) == Pos(FToP(board.BlackKingPosF)) {
				return true
			}
		}
		if posX < Width-1 {
			if posY > 0 && (pos-(Width-1) == Pos(FToP(board.BlackKingPosF)) || board.FieldsF[PToFp(pos-(Width-1))] == piece.BlackPawn) {
				return true
			}
			if pos+1 == Pos(FToP(board.BlackKingPosF)) {
				return true
			}
			if posY < Height-1 && pos+(Width+1) == Pos(FToP(board.BlackKingPosF)) {
				return true
			}
		}
		if posY > 0 && pos-Width == Pos(FToP(board.BlackKingPosF)) {
			return true
		}
		if posY < Height-1 && pos+Width == Pos(FToP(board.BlackKingPosF)) {
			return true
		}
	}

	// --- check knight ---
	{
		knight := checkerColor | piece.Knight
		if posX > 0 {
			if posY > 1 && board.FieldsF[PToFp(pos-(Width*2+1))] == knight { // -1, -2
				return true
			}
			if posY < Height-2 && board.FieldsF[PToFp(pos+(Width*2-1))] == knight { // -1, +2
				return true
			}
			if posX > 1 {
				if posY > 0 && board.FieldsF[PToFp(pos-(Width+2))] == knight { // -2, -1
					return true
				}
				if posY < Height-1 && board.FieldsF[PToFp(pos+(Width-2))] == knight { // -2, +1
					return true
				}
			}
		}
		if posX < Width-1 {
			if posY > 1 && board.FieldsF[PToFp(pos-(Width*2-1))] == knight { // +1, -2
				return true
			}
			if posY < Height-2 && board.FieldsF[PToFp(pos+(Width*2+1))] == knight { // +1, +2
				return true
			}
			if posX < Width-2 {
				if posY > 0 && board.FieldsF[PToFp(pos-(Width-2))] == knight { // +2, +1
					return true
				}
				if posY < Height-1 && board.FieldsF[PToFp(pos+(Width+2))] == knight { // +2, -1
					return true
				}
			}
		}
	}

	// --- check vertical and horizontal lines ---
	{
		for i := 1; i < Width; i++ {
			if posX-i < 0 {
				break
			}
			f := board.FieldsF[PToFp(pos-Pos(i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Rook|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX+i >= Width {
				break
			}
			f := board.FieldsF[PToFp(pos+Pos(i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Rook|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
		for i := 1; i < Height; i++ {
			if posY-i < 0 {
				break
			}
			f := board.FieldsF[PToFp(pos-Pos(Width*i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Rook|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
		for i := 1; i < Height; i++ {
			if posY+i >= Height {
				break
			}
			f := board.FieldsF[PToFp(pos+Pos(Width*i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Rook|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
	}

	// --- check diagonal lines ---
	{
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY-i < 0 {
				break
			}
			f := board.FieldsF[PToFp(pos-Pos(Width*i+i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Bishop|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY+i >= Height {
				break
			}
			f := board.FieldsF[PToFp(pos+Pos(Width*i-i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Bishop|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY-i < 0 {
				break
			}
			f := board.FieldsF[PToFp(pos-Pos(Width*i-i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Bishop|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY+i >= Height {
				break
			}
			f := board.FieldsF[PToFp(pos+Pos(Width*i+i))]
			if f == piece.None {
				continue
			}
			if f&(piece.Bishop|piece.Queen) != piece.None && f&checkerColor != piece.None {
				return true
			}
			break
		}
	}

	return false
}
