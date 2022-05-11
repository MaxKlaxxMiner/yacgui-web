package Crc64

import "github.com/MaxKlaxxMiner/yacgui-web/YacBoard"

func FromBoardFull(board *YacBoard.YacBoard) Value {
	return CrcStart.UpdatePieces(board.Fields[:]).
		UpdateBool(board.WhiteMove).
		UpdateBool(board.WhiteCanCastleKingside).UpdateBool(board.WhiteCanCastleQueenside).
		UpdateBool(board.BlackCanCastleKingside).UpdateBool(board.BlackCanCastleQueenside).
		UpdateInt(board.EnPassantPos).
		//UpdateInt(board.HalfmoveClock).
		UpdateInt(board.MoveNumber)
}

func FromBoard(board *YacBoard.YacBoard) Value {
	return CrcStart.UpdatePieces(board.Fields[:]).
		UpdateBool(board.WhiteMove).
		UpdateBool(board.WhiteCanCastleKingside).UpdateBool(board.WhiteCanCastleQueenside).
		UpdateBool(board.BlackCanCastleKingside).UpdateBool(board.BlackCanCastleQueenside).
		UpdateInt(board.EnPassantPos)
}
