package YacBoard

func PieceFromChar(c rune) Piece {
	switch c {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return PieceNone
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

func PieceToChar(piece Piece) byte {
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

func PosFrom2Chars(str string) int {
	if len(str) != 2 {
		return -1
	}
	if str[0] < 'a' || str[0]-'a' >= Width {
		return -1
	}
	if str[1] < '1' || str[1]-'1' >= Height {
		return -1
	}
	return int(str[0] - 'a' + (Height+'0'-str[1])*Width)
}

func PosToChars(pos int) string {
	if uint(pos) >= Width*Height {
		return "-"
	}
	return PosXYToChars(pos%Width, pos/Width)
}

func PosXYToChars(x, y int) string {
	return string([]byte{byte(x + 'a'), byte((Height - y) + '0')})
}
