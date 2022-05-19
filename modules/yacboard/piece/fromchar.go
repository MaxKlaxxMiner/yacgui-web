package piece

func FromChar(c byte) Piece {
	switch c {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return None
	case 'K':
		return WhiteKing
	case 'k':
		return BlackKing
	case 'Q':
		return WhiteQueen
	case 'q':
		return BlackQueen
	case 'R':
		return WhiteRook
	case 'r':
		return BlackRook
	case 'B':
		return WhiteBishop
	case 'b':
		return BlackBishop
	case 'N':
		return WhiteKnight
	case 'n':
		return BlackKnight
	case 'P':
		return WhitePawn
	case 'p':
		return BlackPawn
	default:
		return Blocked
	}
}
