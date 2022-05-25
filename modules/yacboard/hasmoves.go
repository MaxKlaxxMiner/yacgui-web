package yacboard

import (
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func scanMoveStop(b *YacBoard, pos Pos, foundMove func(pos Pos) bool) bool {
	field := b.Fields[pos]
	if field == piece.None {
		return false
	}
	color := field & piece.Colors
	posX := int(pos % Width)
	posY := int(pos / Width)
	switch field & piece.BasicMask {
	case piece.King:
		if posX > 0 {
			if posY > 0 && b.Fields[pos-(Width+1)]&color == piece.None {
				if foundMove(pos - (Width + 1)) {
					return true
				}
			}
			if b.Fields[pos-1]&color == piece.None {
				if foundMove(pos - 1) {
					return true
				}
			}
			if posY < Height-1 && b.Fields[pos+(Width-1)]&color == piece.None {
				if foundMove(pos + (Width - 1)) {
					return true
				}
			}
		}
		if posX < Width-1 {
			if posY > 0 && b.Fields[pos-(Width-1)]&color == piece.None {
				if foundMove(pos - (Width - 1)) {
					return true
				}
			}
			if b.Fields[pos+1]&color == piece.None {
				if foundMove(pos + 1) {
					return true
				}
			}
			if posY < Height-1 && b.Fields[pos+(Width+1)]&color == piece.None {
				if foundMove(pos + (Width + 1)) {
					return true
				}
			}
		}
		if posY > 0 && b.Fields[pos-Width]&color == piece.None {
			if foundMove(pos - Width) {
				return true
			}
		}
		if posY < Height-1 && b.Fields[pos+Width]&color == piece.None {
			if foundMove(pos + Width) {
				return true
			}
		}

	case piece.Queen:
		// left
		for i := 1; i < Width; i++ {
			if posX-i < 0 {
				break
			}
			p := pos - Pos(i)
			f := b.Fields[p]
			if (f & color) != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// right
		for i := 1; i < Width; i++ {
			if posX+i >= Width {
				break
			}
			p := pos + Pos(i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// up
		for i := 1; i < Height; i++ {
			if posY-i < 0 {
				break
			}
			p := pos - Pos(Width*i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// down
		for i := 1; i < Height; i++ {
			if posY+i >= Height {
				break
			}
			p := pos + Pos(Width*i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// left-up
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY-i < 0 {
				break
			}
			p := pos - Pos(Width*i+i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// left-down
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY+i >= Height {
				break
			}
			p := pos + Pos(Width*i-i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// right-up
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY-i < 0 {
				break
			}
			p := pos - Pos(Width*i-i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// right-down
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY+i >= Height {
				break
			}
			p := pos + Pos(Width*i+i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}

	case piece.Rook:
		// left
		for i := 1; i < Width; i++ {
			if posX-i < 0 {
				break
			}
			p := pos - Pos(i)
			f := b.Fields[p]
			if (f & color) != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// right
		for i := 1; i < Width; i++ {
			if posX+i >= Width {
				break
			}
			p := pos + Pos(i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// up
		for i := 1; i < Height; i++ {
			if posY-i < 0 {
				break
			}
			p := pos - Pos(Width*i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// down
		for i := 1; i < Height; i++ {
			if posY+i >= Height {
				break
			}
			p := pos + Pos(Width*i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}

	case piece.Bishop:
		// left-up
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY-i < 0 {
				break
			}
			p := pos - Pos(Width*i+i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// left-down
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY+i >= Height {
				break
			}
			p := pos + Pos(Width*i-i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// right-up
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY-i < 0 {
				break
			}
			p := pos - Pos(Width*i-i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}
		// right-down
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY+i >= Height {
				break
			}
			p := pos + Pos(Width*i+i)
			f := b.Fields[p]
			if f&color != piece.None {
				break
			}
			if foundMove(p) {
				return true
			}
			if f != piece.None {
				break
			}
		}

	case piece.Knight:
		if posX > 0 {
			if posY > 1 && b.Fields[pos-(Width*2+1)]&color == piece.None {
				if foundMove(pos - (Width*2 + 1)) {
					return true
				}
			}
			if posY < Height-2 && b.Fields[pos+(Width*2-1)]&color == piece.None {
				if foundMove(pos + (Width*2 - 1)) {
					return true
				}
			}
			if posX > 1 {
				if posY > 0 && b.Fields[pos-(Width+2)]&color == piece.None {
					if foundMove(pos - (Width + 2)) {
						return true
					}
				}
				if posY < Height-1 && b.Fields[pos+(Width-2)]&color == piece.None {
					if foundMove(pos + (Width - 2)) {
						return true
					}
				}
			}
		}
		if posX < Width-1 {
			if posY > 1 && b.Fields[pos-(Width*2-1)]&color == piece.None {
				if foundMove(pos - (Width*2 - 1)) {
					return true
				}
			}
			if posY < Height-2 && b.Fields[pos+(Width*2+1)]&color == piece.None {
				if foundMove(pos + (Width*2 + 1)) {
					return true
				}
			}
			if posX < Width-2 {
				if posY > 0 && b.Fields[pos-(Width-2)]&color == piece.None {
					if foundMove(pos - (Width - 2)) {
						return true
					}
				}
				if posY < Height-1 && b.Fields[pos+(Width+2)]&color == piece.None {
					if foundMove(pos + (Width + 2)) {
						return true
					}
				}
			}
		}

	case piece.Pawn:
		if posY < 1 || posY >= Height-1 { // invalid pos?
			break
		}

		if color == piece.White { // white pawn goes up
			if b.Fields[pos-Width] == piece.None {
				if foundMove(pos - Width) {
					return true
				}
				if posY == Height-2 && b.Fields[pos-Width*2] == piece.None {
					if foundMove(pos - Width*2) {
						return true
					}
				}
			}
			if posX > 0 && (b.EnPassantPos == pos-(Width+1) || b.Fields[pos-(Width+1)]&piece.Colors == piece.Black) { // capture left-top
				if foundMove(pos - (Width + 1)) {
					return true
				}
			}
			if posX < Width-1 && (b.EnPassantPos == pos-(Width-1) || b.Fields[pos-(Width-1)]&piece.Colors == piece.Black) { // capture right-top
				if foundMove(pos - (Width - 1)) {
					return true
				}
			}
		} else { // black pawn goes down
			if b.Fields[pos+Width] == piece.None {
				if foundMove(pos + Width) {
					return true
				}
				if posY == 1 && b.Fields[pos+Width*2] == piece.None {
					if foundMove(pos + Width*2) {
						return true
					}
				}
			}
			if posX > 0 && (b.EnPassantPos == pos+(Width-1) || b.Fields[pos+(Width-1)]&piece.Colors == piece.White) {
				if foundMove(pos + (Width - 1)) {
					return true
				}
			}
			if posX < Width-1 && (b.EnPassantPos == pos+(Width+1) || b.Fields[pos+(Width+1)]&piece.Colors == piece.White) {
				if foundMove(pos + (Width + 1)) {
					return true
				}
			}
		}
	}
	return false
}

func hasWhiteMoves(b *YacBoard) bool {
	for pos := len(b.Fields) - 1; pos >= 0; pos-- {
		p := b.Fields[pos]
		if p&piece.Colors != piece.White { // wrong color / no p?
			continue
		}

		if p == piece.WhitePawn && pos < Width*2 {
			// promotion move found?
			if scanMoveStop(b, Pos(pos), func(movePos Pos) bool {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos], PromotionPiece: piece.WhiteQueen}
				return b.simpleMoveCheck(move)
			}) {
				return true
			}
		} else {
			if scanMoveStop(b, Pos(pos), func(movePos Pos) bool {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
				return b.simpleMoveCheck(move)
			}) {
				return true
			}

			if pos == 60 && p == piece.WhiteKing {
				if b.WhiteCanCastleQueenside &&
					b.Fields[57] == piece.None && b.Fields[58] == piece.None && b.Fields[59] == piece.None &&
					!b.isChecked(58, piece.Black) && !b.isChecked(59, piece.Black) && !b.isChecked(60, piece.Black) {
					return true
				}
				if b.WhiteCanCastleKingside &&
					b.Fields[61] == piece.None && b.Fields[62] == piece.None &&
					!b.isChecked(60, piece.Black) && !b.isChecked(61, piece.Black) && !b.isChecked(62, piece.Black) {
					return true
				}
			} else if pos == 4 && p == piece.BlackKing {
				if b.BlackCanCastleQueenside &&
					b.Fields[1] == piece.None && b.Fields[2] == piece.None && b.Fields[3] == piece.None &&
					!b.isChecked(2, piece.White) && !b.isChecked(3, piece.White) && !b.isChecked(4, piece.White) {
					return true
				}
				if b.BlackCanCastleKingside &&
					b.Fields[5] == piece.None && b.Fields[6] == piece.None &&
					!b.isChecked(4, piece.White) && !b.isChecked(5, piece.White) && !b.isChecked(6, piece.White) {
					return true
				}
			}
		}
	}
	return false
}

func hasBlackMoves(b *YacBoard) bool {
	for pos := 0; pos < len(b.Fields); pos++ {
		p := b.Fields[pos]
		if p&piece.Colors != piece.Black { // wrong color / no p?
			continue
		}

		if p == piece.BlackPawn && pos >= Height*Width-Width*2 {
			// promotion move found?
			if scanMoveStop(b, Pos(pos), func(movePos Pos) bool {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos], PromotionPiece: piece.BlackQueen}
				return b.simpleMoveCheck(move)
			}) {
				return true
			}
		} else {
			if scanMoveStop(b, Pos(pos), func(movePos Pos) bool {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
				return b.simpleMoveCheck(move)
			}) {
				return true
			}

			if pos == 60 && p == piece.WhiteKing {
				if b.WhiteCanCastleQueenside &&
					b.Fields[57] == piece.None && b.Fields[58] == piece.None && b.Fields[59] == piece.None &&
					!b.isChecked(58, piece.Black) && !b.isChecked(59, piece.Black) && !b.isChecked(60, piece.Black) {
					return true
				}
				if b.WhiteCanCastleKingside &&
					b.Fields[61] == piece.None && b.Fields[62] == piece.None &&
					!b.isChecked(60, piece.Black) && !b.isChecked(61, piece.Black) && !b.isChecked(62, piece.Black) {
					return true
				}
			} else if pos == 4 && p == piece.BlackKing {
				if b.BlackCanCastleQueenside &&
					b.Fields[1] == piece.None && b.Fields[2] == piece.None && b.Fields[3] == piece.None &&
					!b.isChecked(2, piece.White) && !b.isChecked(3, piece.White) && !b.isChecked(4, piece.White) {
					return true
				}
				if b.BlackCanCastleKingside &&
					b.Fields[5] == piece.None && b.Fields[6] == piece.None &&
					!b.isChecked(4, piece.White) && !b.isChecked(5, piece.White) && !b.isChecked(6, piece.White) {
					return true
				}
			}
		}
	}
	return false
}

func (board *YacBoard) HasMoves() bool {
	if board.WhiteMove {
		kp := board.WhiteKingPos
		posX := kp % Width
		board.Fields[kp] = piece.None

		if posX > 0 && board.Fields[kp-1]&piece.White == piece.None && !board.isChecked(kp-1, piece.Black) {
			board.Fields[kp] = piece.WhiteKing
			return true
		}

		if posX < 7 && board.Fields[kp+1]&piece.White == piece.None && !board.isChecked(kp+1, piece.Black) {
			board.Fields[kp] = piece.WhiteKing
			return true
		}

		board.Fields[kp] = piece.WhiteKing
		return hasWhiteMoves(board)
	} else {
		kp := board.BlackKingPos
		posX := kp % Width
		board.Fields[kp] = piece.None

		if posX > 0 && board.Fields[kp-1]&piece.Black == piece.None && !board.isChecked(kp-1, piece.White) {
			board.Fields[kp] = piece.BlackKing
			return true
		}

		if posX < 7 && board.Fields[kp+1]&piece.Black == piece.None && !board.isChecked(kp+1, piece.White) {
			board.Fields[kp] = piece.BlackKing
			return true
		}
		board.Fields[kp] = piece.BlackKing
		return hasBlackMoves(board)
	}
}
