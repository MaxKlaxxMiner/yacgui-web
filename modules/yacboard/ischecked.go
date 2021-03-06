package yacboard

import (
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) IsChecked() bool {
	if board.WhiteMove {
		return board.isChecked(board.WhiteKingPos, piece.Black)
	} else {
		return board.isChecked(board.BlackKingPos, piece.WhiteKing)
	}
}

func (board *YacBoard) isChecked(pos Pos, checkerColor piece.Piece) bool {
	posX := int(pos % Width)
	posY := int(pos / Width)

	// --- check pawn and king ---
	if checkerColor == piece.White {
		if posX > 0 {
			if posY > 0 && pos-(Width+1) == board.WhiteKingPos {
				return true
			}
			if pos-1 == board.WhiteKingPos {
				return true
			}
			if posY < Height-1 && (pos+(Width-1) == board.WhiteKingPos || board.Fields[pos+(Width-1)] == piece.WhitePawn) {
				return true
			}
		}
		if posX < Width-1 {
			if posY > 0 && pos-(Width-1) == board.WhiteKingPos {
				return true
			}
			if pos+1 == board.WhiteKingPos {
				return true
			}
			if posY < Height-1 && (pos+(Width+1) == board.WhiteKingPos || board.Fields[pos+(Width+1)] == piece.WhitePawn) {
				return true
			}
		}
		if posY > 0 && pos-Width == board.WhiteKingPos {
			return true
		}
		if posY < Height-1 && pos+Width == board.WhiteKingPos {
			return true
		}
	} else {
		if posX > 0 {
			if posY > 0 && (pos-(Width+1) == board.BlackKingPos || board.Fields[pos-(Width+1)] == piece.BlackPawn) {
				return true
			}
			if pos-1 == board.BlackKingPos {
				return true
			}
			if posY < Height-1 && pos+(Width-1) == board.BlackKingPos {
				return true
			}
		}
		if posX < Width-1 {
			if posY > 0 && (pos-(Width-1) == board.BlackKingPos || board.Fields[pos-(Width-1)] == piece.BlackPawn) {
				return true
			}
			if pos+1 == board.BlackKingPos {
				return true
			}
			if posY < Height-1 && pos+(Width+1) == board.BlackKingPos {
				return true
			}
		}
		if posY > 0 && pos-Width == board.BlackKingPos {
			return true
		}
		if posY < Height-1 && pos+Width == board.BlackKingPos {
			return true
		}
	}

	// --- check knight ---
	{
		knight := checkerColor | piece.Knight
		if posX > 0 {
			if posY > 1 && board.Fields[pos-(Width*2+1)] == knight { // -1, -2
				return true
			}
			if posY < Height-2 && board.Fields[pos+(Width*2-1)] == knight { // -1, +2
				return true
			}
			if posX > 1 {
				if posY > 0 && board.Fields[pos-(Width+2)] == knight { // -2, -1
					return true
				}
				if posY < Height-1 && board.Fields[pos+(Width-2)] == knight { // -2, +1
					return true
				}
			}
		}
		if posX < Width-1 {
			if posY > 1 && board.Fields[pos-(Width*2-1)] == knight { // +1, -2
				return true
			}
			if posY < Height-2 && board.Fields[pos+(Width*2+1)] == knight { // +1, +2
				return true
			}
			if posX < Width-2 {
				if posY > 0 && board.Fields[pos-(Width-2)] == knight { // +2, +1
					return true
				}
				if posY < Height-1 && board.Fields[pos+(Width+2)] == knight { // +2, -1
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
			f := board.Fields[pos-Pos(i)]
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
			f := board.Fields[pos+Pos(i)]
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
			f := board.Fields[pos-Pos(Width*i)]
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
			f := board.Fields[pos+Pos(Width*i)]
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
			f := board.Fields[pos-Pos(Width*i+i)]
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
			f := board.Fields[pos+Pos(Width*i-i)]
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
			f := board.Fields[pos-Pos(Width*i-i)]
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
			f := board.Fields[pos+Pos(Width*i+i)]
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
