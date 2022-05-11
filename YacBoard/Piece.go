package YacBoard

type Piece byte

const (
	PieceNone Piece = 0x00
	Blocked         = White | Black

	White  Piece = 0x40
	Black  Piece = 0x80
	Colors       = White | Black

	King        Piece = 0x01
	Queen       Piece = 0x02
	Rook        Piece = 0x04
	Bishop      Piece = 0x08
	Knight      Piece = 0x10
	Pawn        Piece = 0x20
	BasicPieces       = King | Queen | Rook | Bishop | Knight | Pawn

	WhiteKing   = White | King
	WhiteQueen  = White | Queen
	WhiteRook   = White | Rook
	WhiteBishop = White | Bishop
	WhiteKnight = White | Knight
	WhitePawn   = White | Pawn
	BlackKing   = Black | King
	BlackQueen  = Black | Queen
	BlackRook   = Black | Rook
	BlackBishop = Black | Bishop
	BlackKnight = Black | Knight
	BlackPawn   = Black | Pawn
)
