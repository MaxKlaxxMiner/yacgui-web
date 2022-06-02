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

func (board *YacBoard) simpleMoveCheck(move Move) bool {
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

func (board *YacBoard) MoveCheck(move Move) bool {
	for _, m := range board.GetMoves() {
		if m == move {
			return true
		}
	}
	return false
}

func getWhiteMoves(b *YacBoard, mv *[256]Move) byte {
	var mi byte = 0
	for pos := len(b.Fields) - 1; pos >= 0; pos-- {
		field := b.Fields[pos]
		if field&piece.Colors != piece.White { // wrong color / no p?
			continue
		}
		posX := pos % Width
		posY := pos / Width

		if field == piece.WhitePawn && pos < Width*2 {
			if b.Fields[pos-Width] == piece.None {
				move := Move{FromPos: byte(pos), ToPos: byte(pos - Width), CapturePiece: b.Fields[pos-Width], PromotionPiece: piece.WhiteQueen}
				if b.simpleMoveCheck(move) {
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
			}
			if posX > 0 && b.Fields[pos-(Width+1)]&piece.Colors == piece.Black { // capture left-top
				move := Move{FromPos: byte(pos), ToPos: byte(pos - (Width + 1)), CapturePiece: b.Fields[pos-(Width+1)], PromotionPiece: piece.WhiteQueen}
				if b.simpleMoveCheck(move) {
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
			}
			if posX < Width-1 && b.Fields[pos-(Width-1)]&piece.Colors == piece.Black { // capture right-top
				move := Move{FromPos: byte(pos), ToPos: byte(pos - (Width - 1)), CapturePiece: b.Fields[pos-(Width-1)], PromotionPiece: piece.WhiteQueen}
				if b.simpleMoveCheck(move) {
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
			}
		} else {
			switch field {
			case piece.WhiteKing:
				var movePos int
				if posX > 0 {
					movePos = pos - (Width + 1) // left-up
					if posY > 0 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos - 1 // left
					if b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width - 1) // left-down
					if posY < Height-1 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
				}
				if posX < Width-1 {
					movePos = pos - (Width - 1) // right-up
					if posY > 0 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + 1 // right
					if b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width + 1) // right-down
					if posY < Height-1 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
				}
				movePos = pos - Width // up
				if posY > 0 && b.Fields[movePos]&piece.White == piece.None {
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
				movePos = pos + Width // down
				if posY < Height-1 && b.Fields[movePos]&piece.White == piece.None {
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
				if pos == 60 {
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
				}

			case piece.WhiteQueen:
				// left
				for i := 1; i < Width; i++ {
					if posX-i < 0 {
						break
					}
					p := pos - i
					f := b.Fields[p]
					if (f & piece.White) != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + i
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - Width*i
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + Width*i
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - (Width*i + i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i - i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - (Width*i - i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i + i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					if f != piece.None {
						break
					}
				}

			case piece.WhiteRook:
				// left
				for i := 1; i < Width; i++ {
					if posX-i < 0 {
						break
					}
					p := pos - i
					f := b.Fields[p]
					if (f & piece.White) != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + i
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - Width*i
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + Width*i
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					if f != piece.None {
						break
					}
				}

			case piece.WhiteBishop:
				// left-up
				for i := 1; i < Width; i++ {
					if posX-i < 0 || posY-i < 0 {
						break
					}
					p := pos - (Width*i + i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i - i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - (Width*i - i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i + i)
					f := b.Fields[p]
					if f&piece.White != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					if f != piece.None {
						break
					}
				}

			case piece.WhiteKnight:
				var movePos int
				if posX > 0 {
					movePos = pos - (Width*2 + 1) // -1, -2
					if posY > 1 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width*2 - 1) // -1, +2
					if posY < Height-2 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					if posX > 1 {
						movePos = pos - (Width + 2) // -2, -1
						if posY > 0 && b.Fields[movePos]&piece.White == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
						movePos = pos + (Width - 2) // -2, +1
						if posY < Height-1 && b.Fields[movePos]&piece.White == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
					}
				}
				if posX < Width-1 {
					movePos = pos - (Width*2 - 1) // +1, -2
					if posY > 1 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width*2 + 1) // +1, +2
					if posY < Height-2 && b.Fields[movePos]&piece.White == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					if posX < Width-2 {
						movePos = pos - (Width - 2) // +2, +1
						if posY > 0 && b.Fields[movePos]&piece.White == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
						movePos = pos + (Width + 2) // +2, -1
						if posY < Height-1 && b.Fields[movePos]&piece.White == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
					}
				}

			case piece.WhitePawn:
				if posY < 1 || posY >= Height-1 { // invalid pos?
					break
				}
				var movePos int
				movePos = pos - Width
				if b.Fields[movePos] == piece.None {
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					movePos = pos - Width*2
					if posY == Height-2 && b.Fields[movePos] == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
				}
				movePos = pos - (Width + 1)
				if posX > 0 && (int(b.EnPassantPos) == movePos || b.Fields[movePos]&piece.Colors == piece.Black) { // capture left-top
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
				movePos = pos - (Width - 1)
				if posX < Width-1 && (int(b.EnPassantPos) == movePos || b.Fields[movePos]&piece.Colors == piece.Black) { // capture right-top
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
			}
		}
	}
	return mi
}

func getBlackMoves(b *YacBoard, mv *[256]Move) byte {
	var mi byte = 0
	for pos := 0; pos < len(b.Fields); pos++ {
		field := b.Fields[pos]
		if field&piece.Colors != piece.Black { // wrong color / no p?
			continue
		}
		posX := pos % Width
		posY := pos / Width

		if field == piece.BlackPawn && pos >= Height*Width-Width*2 {
			if b.Fields[pos+Width] == piece.None {
				move := Move{FromPos: byte(pos), ToPos: byte(pos + Width), CapturePiece: b.Fields[pos+Width], PromotionPiece: piece.BlackQueen}
				if b.simpleMoveCheck(move) {
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
			}
			if posX > 0 && b.Fields[pos+(Width-1)]&piece.Colors == piece.White { // capture left-bottom
				move := Move{FromPos: byte(pos), ToPos: byte(pos + (Width - 1)), CapturePiece: b.Fields[pos+(Width-1)], PromotionPiece: piece.BlackQueen}
				if b.simpleMoveCheck(move) {
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
			}
			if posX < Width-1 && b.Fields[pos+(Width+1)]&piece.Colors == piece.White { // capture right-bottom
				move := Move{FromPos: byte(pos), ToPos: byte(pos + (Width + 1)), CapturePiece: b.Fields[pos+(Width+1)], PromotionPiece: piece.BlackQueen}
				if b.simpleMoveCheck(move) {
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
			}
		} else {
			switch field {
			case piece.BlackKing:
				var movePos int
				if posX > 0 {
					movePos = pos - (Width + 1) // left-up
					if posY > 0 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos - 1 // left
					if b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width - 1) // left-down
					if posY < Height-1 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
				}
				if posX < Width-1 {
					movePos = pos - (Width - 1) // right-up
					if posY > 0 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + 1 // right
					if b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width + 1) // right-down
					if posY < Height-1 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
				}
				movePos = pos - Width // up
				if posY > 0 && b.Fields[movePos]&piece.Black == piece.None {
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
				movePos = pos + Width // down
				if posY < Height-1 && b.Fields[movePos]&piece.Black == piece.None {
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
				if pos == 4 {
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

			case piece.BlackQueen:
				// left
				for i := 1; i < Width; i++ {
					if posX-i < 0 {
						break
					}
					p := pos - i
					f := b.Fields[p]
					if (f & piece.Black) != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + i
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - Width*i
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + Width*i
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - (Width*i + i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i - i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - (Width*i - i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i + i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					if f != piece.None {
						break
					}
				}

			case piece.BlackRook:
				// left
				for i := 1; i < Width; i++ {
					if posX-i < 0 {
						break
					}
					p := pos - i
					f := b.Fields[p]
					if (f & piece.Black) != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + i
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - Width*i
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + Width*i
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					if f != piece.None {
						break
					}
				}

			case piece.BlackBishop:
				// left-up
				for i := 1; i < Width; i++ {
					if posX-i < 0 || posY-i < 0 {
						break
					}
					p := pos - (Width*i + i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i - i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos - (Width*i - i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
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
					p := pos + (Width*i + i)
					f := b.Fields[p]
					if f&piece.Black != piece.None {
						break
					}
					move := Move{FromPos: byte(pos), ToPos: byte(p), CapturePiece: f}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					if f != piece.None {
						break
					}
				}

			case piece.BlackKnight:
				var movePos int
				if posX > 0 {
					movePos = pos - (Width*2 + 1) // -1, -2
					if posY > 1 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width*2 - 1) // -1, +2
					if posY < Height-2 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					if posX > 1 {
						movePos = pos - (Width + 2) // -2, -1
						if posY > 0 && b.Fields[movePos]&piece.Black == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
						movePos = pos + (Width - 2) // -2, +1
						if posY < Height-1 && b.Fields[movePos]&piece.Black == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
					}
				}
				if posX < Width-1 {
					movePos = pos - (Width*2 - 1) // +1, -2
					if posY > 1 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					movePos = pos + (Width*2 + 1) // +1, +2
					if posY < Height-2 && b.Fields[movePos]&piece.Black == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
					if posX < Width-2 {
						movePos = pos - (Width - 2) // +2, +1
						if posY > 0 && b.Fields[movePos]&piece.Black == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
						movePos = pos + (Width + 2) // +2, -1
						if posY < Height-1 && b.Fields[movePos]&piece.Black == piece.None {
							move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
							if b.simpleMoveCheck(move) {
								mv[mi] = move
								mi++
							}
						}
					}
				}

			case piece.BlackPawn:
				if posY < 1 || posY >= Height-1 { // invalid pos?
					break
				}
				var movePos int
				movePos = pos + Width
				if b.Fields[movePos] == piece.None {
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
					movePos = pos + Width*2
					if posY == 1 && b.Fields[movePos] == piece.None {
						move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
						if b.simpleMoveCheck(move) {
							mv[mi] = move
							mi++
						}
					}
				}
				movePos = pos + (Width - 1)
				if posX > 0 && (int(b.EnPassantPos) == movePos || b.Fields[movePos]&piece.Colors == piece.White) { // capture left-top
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
				}
				movePos = pos + (Width + 1)
				if posX < Width-1 && (int(b.EnPassantPos) == movePos || b.Fields[movePos]&piece.Colors == piece.White) { // capture right-top
					move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
					if b.simpleMoveCheck(move) {
						mv[mi] = move
						mi++
					}
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
