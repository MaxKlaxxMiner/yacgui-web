package yacboard

func (board *YacBoard) Clear() {
	for i := 0; i < len(board.Fields); i++ {
		board.Fields[i] = 0
	}

	board.HalfmoveClock = 0
	board.MoveNumber = 1
	board.WhiteKingPos = -1
	board.BlackKingPos = -1
	board.EnPassantPos = -1

	board.WhiteMove = true
	board.WhiteCanCastleKingside = false
	board.WhiteCanCastleQueenside = false
	board.BlackCanCastleKingside = false
	board.BlackCanCastleQueenside = false
}
