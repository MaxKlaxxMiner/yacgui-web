package piece

func ToChar(piece Piece) byte {
	switch piece {
	case WhiteKing:
		return 'K'
	case BlackKing:
		return 'k'
	case WhiteQueen:
		return 'Q'
	case BlackQueen:
		return 'q'
	case WhiteRook:
		return 'R'
	case BlackRook:
		return 'r'
	case WhiteBishop:
		return 'B'
	case BlackBishop:
		return 'b'
	case WhiteKnight:
		return 'N'
	case BlackKnight:
		return 'n'
	case WhitePawn:
		return 'P'
	case BlackPawn:
		return 'p'
	default:
		return '.'
	}
}

func (piece Piece) String() string {
	return string(ToChar(piece))
}
