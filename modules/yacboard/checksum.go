package yacboard

import "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/crc64"

func ChecksumFull(board *YacBoard) crc64.Value {
	return crc64.CrcStart.UpdatePieces(board.FieldsF[:]).
		UpdateBool(board.WhiteMove).
		UpdateBool(board.WhiteCanCastleKingside).UpdateBool(board.WhiteCanCastleQueenside).
		UpdateBool(board.BlackCanCastleKingside).UpdateBool(board.BlackCanCastleQueenside).
		UpdateInt(int(board.EnPassantPos)).
		//UpdateInt(board.HalfmoveClock).
		UpdateInt(board.MoveNumber)
}

func Checksum(board *YacBoard) crc64.Value {
	return crc64.CrcStart.UpdatePieces(board.FieldsF[:]).
		UpdateBool(board.WhiteMove).
		UpdateBool(board.WhiteCanCastleKingside).UpdateBool(board.WhiteCanCastleQueenside).
		UpdateBool(board.BlackCanCastleKingside).UpdateBool(board.BlackCanCastleQueenside).
		UpdateInt(int(board.EnPassantPos))
}
