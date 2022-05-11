package YacBoard

type BoardInfo uint32

const (
	BoardInfoNone           BoardInfo = 0
	EnPassantNone           BoardInfo = 0xff
	EnPassantMask           BoardInfo = 0xff
	EnPassantBlackA6        BoardInfo = 16
	EnPassantBlackB6        BoardInfo = 17
	EnPassantBlackC6        BoardInfo = 18
	EnPassantBlackD6        BoardInfo = 19
	EnPassantBlackE6        BoardInfo = 20
	EnPassantBlackF6        BoardInfo = 21
	EnPassantBlackG6        BoardInfo = 22
	EnPassantBlackH6        BoardInfo = 23
	EnPassantWhiteA3        BoardInfo = 40
	EnPassantWhiteB3        BoardInfo = 41
	EnPassantWhiteC3        BoardInfo = 42
	EnPassantWhiteD3        BoardInfo = 43
	EnPassantWhiteE3        BoardInfo = 44
	EnPassantWhiteF3        BoardInfo = 45
	EnPassantWhiteG3        BoardInfo = 46
	EnPassantWhiteH3        BoardInfo = 47
	WhiteCanCastleKingside  BoardInfo = 0x0100
	WhiteCanCastleQueenside BoardInfo = 0x0200
	BlackCanCastleKingside  BoardInfo = 0x0400
	BlackCanCastleQueenside BoardInfo = 0x0800
	HalfmoveCounterMask     BoardInfo = 0xffff0000
)
