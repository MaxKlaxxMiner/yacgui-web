package yacboard

import (
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) invertedMoveColor() piece.Piece {
	if board.WhiteMove {
		return piece.Black
	} else {
		return piece.White
	}
}

func (board *YacBoard) moveCheck(move Move) bool {
	p := board.Fields[move.FromPos]

	board.Fields[move.ToPos] = p
	board.Fields[move.FromPos] = piece.None

	if Pos(move.ToPos) == board.EnPassantPos && p&piece.Pawn != piece.None { // "en passant"?
		if board.WhiteMove {
			board.Fields[move.ToPos+Width] = piece.None
		} else {
			board.Fields[move.ToPos-Width] = piece.None
		}
	}

	if move.PromotionPiece != piece.None {
		board.Fields[move.ToPos] = move.PromotionPiece
	}

	if p&piece.King != piece.None { // kingmove?
		if p == piece.WhiteKing {
			board.WhiteKingPos = Pos(move.ToPos)
		} else {
			board.BlackKingPos = Pos(move.ToPos)
		}
	}

	// --- is the king is in check ---
	{
		var kingPos Pos
		if board.WhiteMove {
			kingPos = board.WhiteKingPos
		} else {
			kingPos = board.BlackKingPos
		}
		if board.isChecked(kingPos, board.invertedMoveColor()) {
			board.Fields[move.ToPos] = move.CapturePiece
			board.Fields[move.FromPos] = p
			if Pos(move.ToPos) == board.EnPassantPos && p&piece.Pawn != piece.None { // "en passant" ?
				if board.WhiteMove {
					board.Fields[move.ToPos+Width] = piece.BlackPawn
				} else {
					board.Fields[move.ToPos-Width] = piece.WhitePawn
				}
			}
			if p&piece.King != piece.None {
				if p == piece.WhiteKing {
					board.WhiteKingPos = Pos(move.FromPos)
				} else {
					board.BlackKingPos = Pos(move.FromPos)
				}
			}
			return false
		}
	}

	board.Fields[move.ToPos] = move.CapturePiece
	board.Fields[move.FromPos] = p
	if Pos(move.ToPos) == board.EnPassantPos && p&piece.Pawn != piece.None {
		if board.WhiteMove {
			board.Fields[move.ToPos+Width] = piece.BlackPawn
		} else {
			board.Fields[move.ToPos-Width] = piece.WhitePawn
		}
	}
	if p&piece.King != piece.None {
		if p == piece.WhiteKing {
			board.WhiteKingPos = Pos(move.FromPos)
		} else {
			board.BlackKingPos = Pos(move.FromPos)
		}
	}
	return true
}

func scanMove(board *YacBoard, pos Pos, foundMove func(pos Pos)) {
	field := board.Fields[pos]
	if field == piece.None {
		return
	}
	color := field & piece.Colors
	posX := int(pos % Width)
	posY := int(pos / Width)
	switch field & piece.BasicMask {
	case piece.King:
		if posX > 0 {
			if posY > 0 && board.Fields[pos-(Width+1)]&color == piece.None {
				foundMove(pos - (Width + 1)) // left-up
			}
			if board.Fields[pos-1]&color == piece.None {
				foundMove(pos - 1) // left
			}
			if posY < Height-1 && board.Fields[pos+(Width-1)]&color == piece.None {
				foundMove(pos + (Width - 1)) // left-down
			}
		}
		if posX < Width-1 {
			if posY > 0 && board.Fields[pos-(Width-1)]&color == piece.None {
				foundMove(pos - (Width - 1)) // right-up
			}
			if board.Fields[pos+1]&color == piece.None {
				foundMove(pos + 1) // right
			}
			if posY < Height-1 && board.Fields[pos+(Width+1)]&color == piece.None {
				foundMove(pos + (Width + 1)) // right-down
			}
		}
		if posY > 0 && board.Fields[pos-Width]&color == piece.None {
			foundMove(pos - Width) // up
		}
		if posY < Height-1 && board.Fields[pos+Width]&color == piece.None {
			foundMove(pos + Width) // down
		}

	case piece.Queen:
		// left
		for i := 1; i < Width; i++ {
			if posX-i < 0 {
				break
			}
			p := pos - Pos(i)
			f := board.Fields[p]
			if (f & color) != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if (f & color) != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
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
			f := board.Fields[p]
			if f&color != piece.None {
				break
			}
			foundMove(p)
			if f != piece.None {
				break
			}
		}

	case piece.Knight:
		if posX > 0 {
			if posY > 1 && board.Fields[pos-(Width*2+1)]&color == piece.None {
				foundMove(pos - (Width*2 + 1)) // -1, -2
			}
			if posY < Height-2 && board.Fields[pos+(Width*2-1)]&color == piece.None {
				foundMove(pos + (Width*2 - 1)) // -1, +2
			}
			if posX > 1 {
				if posY > 0 && board.Fields[pos-(Width+2)]&color == piece.None {
					foundMove(pos - (Width + 2)) // -2, -1
				}
				if posY < Height-1 && board.Fields[pos+(Width-2)]&color == piece.None {
					foundMove(pos + (Width - 2)) // -2, +1
				}
			}
		}
		if posX < Width-1 {
			if posY > 1 && board.Fields[pos-(Width*2-1)]&color == piece.None {
				foundMove(pos - (Width*2 - 1)) // +1, -2
			}
			if posY < Height-2 && board.Fields[pos+(Width*2+1)]&color == piece.None {
				foundMove(pos + (Width*2 + 1)) // +1, +2
			}
			if posX < Width-2 {
				if posY > 0 && board.Fields[pos-(Width-2)]&color == piece.None {
					foundMove(pos - (Width - 2)) // +2, +1
				}
				if posY < Height-1 && board.Fields[pos+(Width+2)]&color == piece.None {
					foundMove(pos + (Width + 2)) // +2, -1
				}
			}
		}

	case piece.Pawn:
		if posY < 1 || posY >= Height-1 { // invalid pos?
			break
		}

		if color == piece.White { // white pawn goes up
			if board.Fields[pos-Width] == piece.None {
				foundMove(pos - Width)
				if posY == Height-2 && board.Fields[pos-Width*2] == piece.None {
					foundMove(pos - Width*2)
				}
			}
			if posX > 0 && (board.EnPassantPos == pos-(Width+1) || board.Fields[pos-(Width+1)]&piece.Colors == piece.Black) { // capture left-top
				foundMove(pos - (Width + 1))
			}
			if posX < Width-1 && (board.EnPassantPos == pos-(Width-1) || board.Fields[pos-(Width-1)]&piece.Colors == piece.Black) { // capture right-top
				foundMove(pos - (Width - 1))
			}
		} else { // black pawn goes down
			if board.Fields[pos+Width] == piece.None {
				foundMove(pos + Width)
				if posY == 1 && board.Fields[pos+Width*2] == piece.None {
					foundMove(pos + Width*2)
				}
			}
			if posX > 0 && (board.EnPassantPos == pos+(Width-1) || board.Fields[pos+(Width-1)]&piece.Colors == piece.White) {
				foundMove(pos + (Width - 1))
			}
			if posX < Width-1 && (board.EnPassantPos == pos+(Width+1) || board.Fields[pos+(Width+1)]&piece.Colors == piece.White) {
				foundMove(pos + (Width + 1))
			}
		}
	}
}

func getWhiteMoves(b *YacBoard, mv *[256]Move) byte {
	var mi byte = 0
	for pos := len(b.Fields) - 1; pos >= 0; pos-- {
		p := b.Fields[pos]
		if p&piece.Colors != piece.White { // wrong color / no p?
			continue
		}

		if p == piece.WhitePawn && pos < Width*2 {
			// promotion move found?
			scanMove(b, Pos(pos), func(movePos Pos) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos], PromotionPiece: piece.WhiteQueen}
				if b.moveCheck(move) {
					mv[mi] = move
					mi++
					move.PromotionPiece = piece.WhiteRook
					mv[mi] = move
					mi++
					move.PromotionPiece = piece.WhiteBishop
					mv[mi] = move
					mi++
					move.PromotionPiece = piece.WhiteKnight
					mv[mi] = move
					mi++
				}
			})
		} else {
			scanMove(b, Pos(pos), func(movePos Pos) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
				if b.moveCheck(move) {
					mv[mi] = move
					mi++
				}
			})

			if pos == 60 && p == piece.WhiteKing {
				if b.WhiteCanCastleQueenside &&
					b.Fields[57] == piece.None && b.Fields[58] == piece.None && b.Fields[59] == piece.None &&
					!b.isChecked(58, piece.Black) && !b.isChecked(59, piece.Black) && !b.isChecked(60, piece.Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.WhiteCanCastleKingside &&
					b.Fields[61] == piece.None && b.Fields[62] == piece.None &&
					!b.isChecked(60, piece.Black) && !b.isChecked(61, piece.Black) && !b.isChecked(62, piece.Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			} else if pos == 4 && p == piece.BlackKing {
				if b.BlackCanCastleQueenside &&
					b.Fields[1] == piece.None && b.Fields[2] == piece.None && b.Fields[3] == piece.None &&
					!b.isChecked(2, piece.White) && !b.isChecked(3, piece.White) && !b.isChecked(4, piece.White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.BlackCanCastleKingside &&
					b.Fields[5] == piece.None && b.Fields[6] == piece.None &&
					!b.isChecked(4, piece.White) && !b.isChecked(5, piece.White) && !b.isChecked(6, piece.White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			}
		}
	}
	return mi
}

func getBlackMoves(b *YacBoard, mv *[256]Move) byte {
	var mi byte = 0
	for pos := 0; pos < len(b.Fields); pos++ {
		p := b.Fields[pos]
		if p&piece.Colors != piece.Black { // wrong color / no p?
			continue
		}

		if p == piece.BlackPawn && pos >= Height*Width-Width*2 {
			// promotion move found?
			scanMove(b, Pos(pos), func(movePos Pos) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos], PromotionPiece: piece.BlackQueen}
				if b.moveCheck(move) {
					mv[mi] = move
					mi++
					move.PromotionPiece = piece.BlackRook
					mv[mi] = move
					mi++
					move.PromotionPiece = piece.BlackBishop
					mv[mi] = move
					mi++
					move.PromotionPiece = piece.BlackKnight
					mv[mi] = move
					mi++
				}
			})
		} else {
			scanMove(b, Pos(pos), func(movePos Pos) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
				if b.moveCheck(move) {
					mv[mi] = move
					mi++
				}
			})

			if pos == 60 && p == piece.WhiteKing {
				if b.WhiteCanCastleQueenside &&
					b.Fields[57] == piece.None && b.Fields[58] == piece.None && b.Fields[59] == piece.None &&
					!b.isChecked(58, piece.Black) && !b.isChecked(59, piece.Black) && !b.isChecked(60, piece.Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.WhiteCanCastleKingside &&
					b.Fields[61] == piece.None && b.Fields[62] == piece.None &&
					!b.isChecked(60, piece.Black) && !b.isChecked(61, piece.Black) && !b.isChecked(62, piece.Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			} else if pos == 4 && p == piece.BlackKing {
				if b.BlackCanCastleQueenside &&
					b.Fields[1] == piece.None && b.Fields[2] == piece.None && b.Fields[3] == piece.None &&
					!b.isChecked(2, piece.White) && !b.isChecked(3, piece.White) && !b.isChecked(4, piece.White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.BlackCanCastleKingside &&
					b.Fields[5] == piece.None && b.Fields[6] == piece.None &&
					!b.isChecked(4, piece.White) && !b.isChecked(5, piece.White) && !b.isChecked(6, piece.White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			}
		}
	}
	return mi
}

func (board *YacBoard) GetMovesFast(moves *[256]Move) byte {
	if board.WhiteMove {
		return getWhiteMoves(board, moves)
	} else {
		return getBlackMoves(board, moves)
	}
}

func (board *YacBoard) GetMoves() []Move {
	var tmp [256]Move
	count := board.GetMovesFast(&tmp)
	return tmp[:count]
}
