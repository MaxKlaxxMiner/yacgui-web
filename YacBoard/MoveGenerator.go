package YacBoard

func scanMove(b *YacBoard, pos int, foundMove func(pos int)) {
	piece := b.Fields[pos]
	if piece == PieceNone {
		return
	}
	color := piece & Colors
	posX := pos % Width
	posY := pos / Width
	switch piece & BasicPieces {
	case King:
		if posX > 0 {
			if posY > 0 && b.Fields[pos-(Width+1)]&color == PieceNone {
				foundMove(pos - (Width + 1)) // left-top
			}
			if b.Fields[pos-1]&color == PieceNone {
				foundMove(pos - 1) // left
			}
			if posY < Height-1 && b.Fields[pos+(Width-1)]&color == PieceNone {
				foundMove(pos + (Width - 1)) // left-bottom
			}
		}
		if posX < Width-1 {
			if posY > 0 && b.Fields[pos-(Width-1)]&color == PieceNone {
				foundMove(pos - (Width - 1)) // right-top
			}
			if b.Fields[pos+1]&color == PieceNone {
				foundMove(pos + 1) // right
			}
			if posY < Height-1 && b.Fields[pos+(Width+1)]&color == PieceNone {
				foundMove(pos + (Width + 1)) // right-bottom
			}
		}
		if posY > 0 && b.Fields[pos-Width]&color == PieceNone {
			foundMove(pos - Width) // top
		}
		if posY < Height-1 && b.Fields[pos+Width]&color == PieceNone {
			foundMove(pos + Width) // bottom
		}

	case Queen:
		// left
		for i := 1; i < Width; i++ {
			if posX-i < 0 {
				break
			}
			p := pos - i
			f := b.Fields[p]
			if (f & color) != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
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
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// top
		for i := 1; i < Height; i++ {
			if posY-i < 0 {
				break
			}
			p := pos - Width*i
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// bottom
		for i := 1; i < Height; i++ {
			if posY+i >= Height {
				break
			}
			p := pos + Width*i
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// left-top
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY-i < 0 {
				break
			}
			p := pos - (Width*i + i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// left-bottom
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY+i >= Height {
				break
			}
			p := pos + (Width*i - i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// right-top
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY-i < 0 {
				break
			}
			p := pos - (Width*i - i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// right-bottom
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY+i >= Height {
				break
			}
			p := pos + (Width*i + i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}

	case Rook:
		// left
		for i := 1; i < Width; i++ {
			if posX-i < 0 {
				break
			}
			p := pos - i
			f := b.Fields[p]
			if (f & color) != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
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
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// top
		for i := 1; i < Height; i++ {
			if posY-i < 0 {
				break
			}
			p := pos - Width*i
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// bottom
		for i := 1; i < Height; i++ {
			if posY+i >= Height {
				break
			}
			p := pos + Width*i
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}

	case Bishop:
		// left-top
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY-i < 0 {
				break
			}
			p := pos - (Width*i + i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// left-bottom
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY+i >= Height {
				break
			}
			p := pos + (Width*i - i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// right-top
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY-i < 0 {
				break
			}
			p := pos - (Width*i - i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}
		// right-bottom
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY+i >= Height {
				break
			}
			p := pos + (Width*i + i)
			f := b.Fields[p]
			if f&color != PieceNone {
				break
			}
			foundMove(p)
			if f != PieceNone {
				break
			}
		}

	case Knight:
		if posX > 0 {
			if posY > 1 && b.Fields[pos-(Width*2+1)]&color == PieceNone {
				foundMove(pos - (Width*2 + 1)) // -1, -2
			}
			if posY < Height-2 && b.Fields[pos+(Width*2-1)]&color == PieceNone {
				foundMove(pos + (Width*2 - 1)) // -1, +2
			}
			if posX > 1 {
				if posY > 0 && b.Fields[pos-(Width+2)]&color == PieceNone {
					foundMove(pos - (Width + 2)) // -2, -1
				}
				if posY < Height-1 && b.Fields[pos+(Width-2)]&color == PieceNone {
					foundMove(pos + (Width - 2)) // -2, +1
				}
			}
		}
		if posX < Width-1 {
			if posY > 1 && b.Fields[pos-(Width*2-1)]&color == PieceNone {
				foundMove(pos - (Width*2 - 1)) // +1, -2
			}
			if posY < Height-2 && b.Fields[pos+(Width*2+1)]&color == PieceNone {
				foundMove(pos + (Width*2 + 1)) // +1, +2
			}
			if posX < Width-2 {
				if posY > 0 && b.Fields[pos-(Width-2)]&color == PieceNone {
					foundMove(pos - (Width - 2)) // +2, +1
				}
				if posY < Height-1 && b.Fields[pos+(Width+2)]&color == PieceNone {
					foundMove(pos + (Width + 2)) // +2, -1
				}
			}
		}

	case Pawn:
		if posY < 1 || posY >= Height-1 { // invalid pos?
			break
		}

		if color == White { // white pawn goes up
			if b.Fields[pos-Width] == PieceNone {
				foundMove(pos - Width)
				if posY == Height-2 && b.Fields[pos-Width*2] == PieceNone {
					foundMove(pos - Width*2)
				}
			}
			if posX > 0 && (b.EnPassantPos == pos-(Width+1) || b.Fields[pos-(Width+1)]&Colors == Black) { // capture left-top
				foundMove(pos - (Width + 1))
			}
			if posX < Width-1 && (b.EnPassantPos == pos-(Width-1) || b.Fields[pos-(Width-1)]&Colors == Black) { // capture right-top
				foundMove(pos - (Width - 1))
			}
		} else { // black pawn goes down
			if b.Fields[pos+Width] == PieceNone {
				foundMove(pos + Width)
				if posY == 1 && b.Fields[pos+Width*2] == PieceNone {
					foundMove(pos + Width*2)
				}
			}
			if posX > 0 && (b.EnPassantPos == pos+(Width-1) || b.Fields[pos+(Width-1)]&Colors == White) {
				foundMove(pos + (Width - 1))
			}
			if posX < Width-1 && (b.EnPassantPos == pos+(Width+1) || b.Fields[pos+(Width+1)]&Colors == White) {
				foundMove(pos + (Width + 1))
			}
		}
	}
}

func getWhiteMoves(b *YacBoard, mv *[256]Move) byte {
	var mi byte = 0
	for pos := len(b.Fields) - 1; pos >= 0; pos-- {
		piece := b.Fields[pos]
		if piece&Colors != White { // wrong color / no piece?
			continue
		}

		if piece == WhitePawn && pos < Width*2 {
			// promotion move found?
			scanMove(b, pos, func(movePos int) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos], PromotionPiece: WhiteQueen}
				if b.MoveCheck(move) {
					mv[mi] = move
					mi++
					move.PromotionPiece = WhiteRook
					mv[mi] = move
					mi++
					move.PromotionPiece = WhiteBishop
					mv[mi] = move
					mi++
					move.PromotionPiece = WhiteKnight
					mv[mi] = move
					mi++
				}
			})
		} else {
			scanMove(b, pos, func(movePos int) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
				if b.MoveCheck(move) {
					mv[mi] = move
					mi++
				}
			})

			if pos == 60 && piece == WhiteKing {
				if b.WhiteCanCastleQueenside &&
					b.Fields[57] == PieceNone && b.Fields[58] == PieceNone && b.Fields[59] == PieceNone &&
					!b.IsChecked(58, Black) && !b.IsChecked(59, Black) && !b.IsChecked(60, Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.WhiteCanCastleKingside &&
					b.Fields[61] == PieceNone && b.Fields[62] == PieceNone &&
					!b.IsChecked(60, Black) && !b.IsChecked(61, Black) && !b.IsChecked(62, Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			} else if pos == 4 && piece == BlackKing {
				if b.BlackCanCastleQueenside &&
					b.Fields[1] == PieceNone && b.Fields[2] == PieceNone && b.Fields[3] == PieceNone &&
					!b.IsChecked(2, White) && !b.IsChecked(3, White) && !b.IsChecked(4, White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.BlackCanCastleKingside &&
					b.Fields[5] == PieceNone && b.Fields[6] == PieceNone &&
					!b.IsChecked(4, White) && !b.IsChecked(5, White) && !b.IsChecked(6, White) {
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
		piece := b.Fields[pos]
		if piece&Colors != Black { // wrong color / no piece?
			continue
		}

		if piece == BlackPawn && pos >= Height*Width-Width*2 {
			// promotion move found?
			scanMove(b, pos, func(movePos int) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos], PromotionPiece: BlackQueen}
				if b.MoveCheck(move) {
					mv[mi] = move
					mi++
					move.PromotionPiece = BlackRook
					mv[mi] = move
					mi++
					move.PromotionPiece = BlackBishop
					mv[mi] = move
					mi++
					move.PromotionPiece = BlackKnight
					mv[mi] = move
					mi++
				}
			})
		} else {
			scanMove(b, pos, func(movePos int) {
				move := Move{FromPos: byte(pos), ToPos: byte(movePos), CapturePiece: b.Fields[movePos]}
				if b.MoveCheck(move) {
					mv[mi] = move
					mi++
				}
			})

			if pos == 60 && piece == WhiteKing {
				if b.WhiteCanCastleQueenside &&
					b.Fields[57] == PieceNone && b.Fields[58] == PieceNone && b.Fields[59] == PieceNone &&
					!b.IsChecked(58, Black) && !b.IsChecked(59, Black) && !b.IsChecked(60, Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.WhiteCanCastleKingside &&
					b.Fields[61] == PieceNone && b.Fields[62] == PieceNone &&
					!b.IsChecked(60, Black) && !b.IsChecked(61, Black) && !b.IsChecked(62, Black) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			} else if pos == 4 && piece == BlackKing {
				if b.BlackCanCastleQueenside &&
					b.Fields[1] == PieceNone && b.Fields[2] == PieceNone && b.Fields[3] == PieceNone &&
					!b.IsChecked(2, White) && !b.IsChecked(3, White) && !b.IsChecked(4, White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) - 2}
					mi++
				}
				if b.BlackCanCastleKingside &&
					b.Fields[5] == PieceNone && b.Fields[6] == PieceNone &&
					!b.IsChecked(4, White) && !b.IsChecked(5, White) && !b.IsChecked(6, White) {
					mv[mi] = Move{FromPos: byte(pos), ToPos: byte(pos) + 2}
					mi++
				}
			}
		}
	}
	return mi
}

func (board *YacBoard) GetMoves(moves *[256]Move) byte {
	if board.WhiteMove {
		return getWhiteMoves(board, moves)
	} else {
		return getBlackMoves(board, moves)
	}
}

func (board *YacBoard) IsChecked(pos int, checkerColor Piece) bool {
	posX := pos % Width
	posY := pos / Width

	// --- check pawn and king ---
	if checkerColor == White {
		if posX > 0 {
			if posY > 0 && pos-(Width+1) == board.WhiteKingPos {
				return true
			}
			if pos-1 == board.WhiteKingPos {
				return true
			}
			if posY < Height-1 && (pos+(Width-1) == board.WhiteKingPos || board.Fields[pos+(Width-1)] == WhitePawn) {
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
			if posY < Height-1 && (pos+(Width+1) == board.WhiteKingPos || board.Fields[pos+(Width+1)] == WhitePawn) {
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
			if posY > 0 && (pos-(Width+1) == board.BlackKingPos || board.Fields[pos-(Width+1)] == BlackPawn) {
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
			if posY > 0 && (pos-(Width-1) == board.BlackKingPos || board.Fields[pos-(Width-1)] == BlackPawn) {
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
		knight := checkerColor | Knight
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
			f := board.Fields[pos-i]
			if f == PieceNone {
				continue
			}
			if f&(Rook|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX+i >= Width {
				break
			}
			f := board.Fields[pos+i]
			if f == PieceNone {
				continue
			}
			if f&(Rook|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
		for i := 1; i < Height; i++ {
			if posY-i < 0 {
				break
			}
			f := board.Fields[pos-Width*i]
			if f == PieceNone {
				continue
			}
			if f&(Rook|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
		for i := 1; i < Height; i++ {
			if posY+i >= Height {
				break
			}
			f := board.Fields[pos+Width*i]
			if f == PieceNone {
				continue
			}
			if f&(Rook|Queen) != PieceNone && f&checkerColor != PieceNone {
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
			f := board.Fields[pos-(Width*i+i)]
			if f == PieceNone {
				continue
			}
			if f&(Bishop|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX-i < 0 || posY+i >= Height {
				break
			}
			f := board.Fields[pos+(Width*i-i)]
			if f == PieceNone {
				continue
			}
			if f&(Bishop|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY-i < 0 {
				break
			}
			f := board.Fields[pos-(Width*i-i)]
			if f == PieceNone {
				continue
			}
			if f&(Bishop|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
		for i := 1; i < Width; i++ {
			if posX+i >= Width || posY+i >= Height {
				break
			}
			f := board.Fields[pos+(Width*i+i)]
			if f == PieceNone {
				continue
			}
			if f&(Bishop|Queen) != PieceNone && f&checkerColor != PieceNone {
				return true
			}
			break
		}
	}

	return false
}

func (board *YacBoard) InvertedMoveColor() Piece {
	if board.WhiteMove {
		return Black
	} else {
		return White
	}
}

func (board *YacBoard) DoMove(move Move) bool {
	piece := board.Fields[move.FromPos]

	board.Fields[move.ToPos] = piece
	board.Fields[move.FromPos] = PieceNone

	if int(move.ToPos) == board.EnPassantPos && piece&Pawn != PieceNone { // "en passant"?
		if board.WhiteMove {
			board.Fields[move.ToPos+Width] = PieceNone
		} else {
			board.Fields[move.ToPos-Width] = PieceNone
		}
	}

	if move.PromotionPiece != PieceNone {
		board.Fields[move.ToPos] = move.PromotionPiece
	}

	if piece&King != PieceNone { // kingmove?
		if piece == WhiteKing {
			board.WhiteKingPos = int(move.ToPos)
		} else {
			board.BlackKingPos = int(move.ToPos)
		}
	}

	// --- is the king is in check ---
	{
		var kingPos int
		if board.WhiteMove {
			kingPos = board.WhiteKingPos
		} else {
			kingPos = board.BlackKingPos
		}
		if kingPos == int(move.ToPos) && (int(move.ToPos)-int(move.FromPos) == 2 || int(move.ToPos)-int(move.FromPos) == -2) {
			switch kingPos {
			case 2:
				board.Fields[0] = PieceNone
				board.Fields[3] = BlackRook
			case 6:
				board.Fields[7] = PieceNone
				board.Fields[5] = BlackRook
			case 58:
				board.Fields[56] = PieceNone
				board.Fields[59] = WhiteRook
			case 62:
				board.Fields[63] = PieceNone
				board.Fields[61] = WhiteRook
			}
		} else if board.IsChecked(kingPos, board.InvertedMoveColor()) {
			board.Fields[move.ToPos] = move.CapturePiece
			board.Fields[move.FromPos] = piece
			if int(move.ToPos) == board.EnPassantPos && piece&Pawn != PieceNone { // "en passant" ?
				if board.WhiteMove {
					board.Fields[move.ToPos+Width] = BlackPawn
				} else {
					board.Fields[move.ToPos-Width] = WhitePawn
				}
			}
			if piece&King != PieceNone { // wurde ein König gezogen?
				if piece == WhiteKing {
					board.WhiteKingPos = int(move.FromPos)
				} else {
					board.BlackKingPos = int(move.FromPos)
				}
			}
			return false
		}
	}

	board.EnPassantPos = -1
	if piece&Pawn != PieceNone && (move.ToPos-move.FromPos == Width*2 || move.FromPos-move.ToPos == Width*2) {
		board.EnPassantPos = (int(move.FromPos) + int(move.ToPos)) / 2
		posX := board.EnPassantPos % Width
		opPawn := false
		if board.WhiteMove {
			if posX > 0 && board.Fields[board.EnPassantPos-Width-1] == BlackPawn {
				opPawn = true
			}
			if posX < Width-1 && board.Fields[board.EnPassantPos-Width+1] == BlackPawn {
				opPawn = true
			}
		} else {
			if posX > 0 && board.Fields[board.EnPassantPos+Width-1] == WhitePawn {
				opPawn = true
			}
			if posX < Width-1 && board.Fields[board.EnPassantPos+Width+1] == WhitePawn {
				opPawn = true
			}
		}
		if !opPawn {
			board.EnPassantPos = -1
		}
	}

	switch move.FromPos {
	case 0:
		board.BlackCanCastleQueenside = false
	case 4:
		board.BlackCanCastleQueenside = false
		board.BlackCanCastleKingside = false
	case 7:
		board.BlackCanCastleKingside = false
	case 56:
		board.WhiteCanCastleQueenside = false
	case 60:
		board.WhiteCanCastleQueenside = false
		board.WhiteCanCastleKingside = false
	case 63:
		board.WhiteCanCastleKingside = false
	}
	switch move.ToPos {
	case 0:
		board.BlackCanCastleQueenside = false
	case 7:
		board.BlackCanCastleKingside = false
	case 56:
		board.WhiteCanCastleQueenside = false
	case 63:
		board.WhiteCanCastleKingside = false
	}

	board.WhiteMove = !board.WhiteMove
	board.HalfmoveClock++
	if piece == Pawn || move.CapturePiece != PieceNone {
		board.HalfmoveClock = 0
	}
	if board.WhiteMove {
		board.MoveNumber++
	}
	return true
}

func (board *YacBoard) MoveCheck(move Move) bool {
	piece := board.Fields[move.FromPos]

	board.Fields[move.ToPos] = piece
	board.Fields[move.FromPos] = PieceNone

	if int(move.ToPos) == board.EnPassantPos && piece&Pawn != PieceNone { // "en passant"?
		if board.WhiteMove {
			board.Fields[move.ToPos+Width] = PieceNone
		} else {
			board.Fields[move.ToPos-Width] = PieceNone
		}
	}

	if move.PromotionPiece != PieceNone {
		board.Fields[move.ToPos] = move.PromotionPiece
	}

	if piece&King != PieceNone { // kingmove?
		if piece == WhiteKing {
			board.WhiteKingPos = int(move.ToPos)
		} else {
			board.BlackKingPos = int(move.ToPos)
		}
	}

	// --- is the king is in check ---
	{
		var kingPos int
		if board.WhiteMove {
			kingPos = board.WhiteKingPos
		} else {
			kingPos = board.BlackKingPos
		}
		if board.IsChecked(kingPos, board.InvertedMoveColor()) {
			board.Fields[move.ToPos] = move.CapturePiece
			board.Fields[move.FromPos] = piece
			if int(move.ToPos) == board.EnPassantPos && piece&Pawn != PieceNone { // "en passant" ?
				if board.WhiteMove {
					board.Fields[move.ToPos+Width] = BlackPawn
				} else {
					board.Fields[move.ToPos-Width] = WhitePawn
				}
			}
			if piece&King != PieceNone { // wurde ein König gezogen?
				if piece == WhiteKing {
					board.WhiteKingPos = int(move.FromPos)
				} else {
					board.BlackKingPos = int(move.FromPos)
				}
			}
			return false
		}
	}

	board.Fields[move.ToPos] = move.CapturePiece
	board.Fields[move.FromPos] = piece
	if int(move.ToPos) == board.EnPassantPos && piece&Pawn != PieceNone {
		if board.WhiteMove {
			board.Fields[move.ToPos+Width] = BlackPawn
		} else {
			board.Fields[move.ToPos-Width] = WhitePawn
		}
	}
	if piece&King != PieceNone {
		if piece == WhiteKing {
			board.WhiteKingPos = int(move.FromPos)
		} else {
			board.BlackKingPos = int(move.FromPos)
		}
	}
	return true
}
