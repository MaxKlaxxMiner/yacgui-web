package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) DoMove(move Move) {
	p := board.FieldsF[move.FromPosF]

	board.FieldsF[move.ToPosF] = p
	board.FieldsF[move.FromPosF] = piece.None

	if Pos(move.ToPosF) == board.EnPassantPosF && p&piece.Pawn != piece.None { // "en passant"?
		if board.WhiteMove {
			board.FieldsF[move.ToPosF+WidthF] = piece.None
		} else {
			board.FieldsF[move.ToPosF-WidthF] = piece.None
		}
	}

	if move.PromotionPiece != piece.None { // pawn move with promotion?
		board.FieldsF[move.ToPosF] = move.PromotionPiece
	}

	if p&piece.King != piece.None { // kingmove?
		if p == piece.WhiteKing {
			board.WhiteKingPos = Pos(FToPb(move.ToPosF))
		} else {
			board.BlackKingPos = Pos(FToPb(move.ToPosF))
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
		if kingPos == Pos(FToPb(move.ToPosF)) && (move.ToPosF-move.FromPosF == 2 || int(move.ToPosF)-int(move.FromPosF) == -2) {
			switch kingPos {
			case 2:
				board.FieldsF[PToF(0)] = piece.None
				board.FieldsF[PToF(3)] = piece.BlackRook
			case 6:
				board.FieldsF[PToF(7)] = piece.None
				board.FieldsF[PToF(5)] = piece.BlackRook
			case 58:
				board.FieldsF[PToF(56)] = piece.None
				board.FieldsF[PToF(59)] = piece.WhiteRook
			case 62:
				board.FieldsF[PToF(63)] = piece.None
				board.FieldsF[PToF(61)] = piece.WhiteRook
			}
		}
	}

	board.EnPassantPosF = -1
	if p&piece.Pawn != piece.None && (move.ToPosF-move.FromPosF == WidthF*2 || move.FromPosF-move.ToPosF == WidthF*2) {
		board.EnPassantPosF = Pos((int(move.FromPosF) + int(move.ToPosF)) / 2)
		posX := board.EnPassantPosF%WidthF + 1
		opPawn := false
		if board.WhiteMove {
			if posX > 0 && board.FieldsF[board.EnPassantPosF-WidthF-1] == piece.BlackPawn {
				opPawn = true
			}
			if posX < Width-1 && board.FieldsF[board.EnPassantPosF-WidthF+1] == piece.BlackPawn {
				opPawn = true
			}
		} else {
			if posX > 0 && board.FieldsF[board.EnPassantPosF+WidthF-1] == piece.WhitePawn {
				opPawn = true
			}
			if posX < Width-1 && board.FieldsF[board.EnPassantPosF+WidthF+1] == piece.WhitePawn {
				opPawn = true
			}
		}
		if !opPawn {
			board.EnPassantPosF = -1
		}
	}

	switch FToPb(move.FromPosF) {
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
	switch FToPb(move.ToPosF) {
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
	if p&piece.Pawn == piece.Pawn || move.CapturePiece != piece.None {
		board.HalfmoveClock = 0
	}
	if board.WhiteMove {
		board.MoveNumber++
	}
}

func (board *YacBoard) DoMoveBackward(move Move, lastBoardInfos BoardInfo) {
	p := board.FieldsF[move.ToPosF]
	board.FieldsF[move.FromPosF] = p
	board.FieldsF[move.ToPosF] = move.CapturePiece

	if move.PromotionPiece != piece.None {
		board.FieldsF[move.FromPosF] = (p & piece.Colors) | piece.Pawn
	}

	if p&piece.Pawn != piece.None &&
		move.FromPosF%WidthF != move.ToPosF%WidthF &&
		move.CapturePiece == piece.None {
		if board.WhiteMove {
			board.FieldsF[PToF((int)(lastBoardInfos&EnPassantMask)-Width)] = piece.WhitePawn
		} else {
			board.FieldsF[PToF((int)(lastBoardInfos&EnPassantMask)+Width)] = piece.BlackPawn
		}
	}

	if p&piece.King != piece.None {
		if p == piece.WhiteKing {
			board.WhiteKingPos = Pos(FToPb(move.FromPosF))
		} else {
			board.BlackKingPos = Pos(FToPb(move.FromPosF))
		}

		posXdif := int(move.FromPosF%WidthF) - int(move.ToPosF%WidthF)
		if posXdif > 1 || posXdif < -1 {
			switch FToPb(move.ToPosF) {
			case 2: // black O-O-O
				board.FieldsF[PToF(0)] = piece.BlackRook
				board.FieldsF[PToF(3)] = piece.None
			case 6: // black O-O
				board.FieldsF[PToF(7)] = piece.BlackRook
				board.FieldsF[PToF(5)] = piece.None
			case 58: // white O-O-O
				board.FieldsF[PToF(56)] = piece.WhiteRook
				board.FieldsF[PToF(59)] = piece.None
			case 62: // white O-O
				board.FieldsF[PToF(63)] = piece.WhiteRook
				board.FieldsF[PToF(61)] = piece.None
			}
		}
	}

	if board.WhiteMove {
		board.MoveNumber--
	}
	board.WhiteMove = !board.WhiteMove
	board.SetBoardInfo(lastBoardInfos)
}
