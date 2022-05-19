package yacboard

import (
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (board *YacBoard) DoMove(move Move) {
	p := board.Fields[move.FromPos]

	board.Fields[move.ToPos] = p
	board.Fields[move.FromPos] = piece.None

	if Pos(move.ToPos) == board.EnPassantPos && p&piece.Pawn != piece.None { // "en passant"?
		if board.WhiteMove {
			board.Fields[move.ToPos+Width] = piece.None
		} else {
			board.Fields[move.ToPos-Width] = piece.None
		}
	}

	if move.PromotionPiece != piece.None { // pawn move with promotion?
		board.Fields[move.ToPos] = move.PromotionPiece
	}

	if p&piece.King != piece.None { // kingmove?
		if p == piece.WhiteKing {
			board.WhiteKingPos = Pos(move.ToPos)
		} else {
			board.BlackKingPos = Pos(move.ToPos)
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
		if kingPos == Pos(move.ToPos) && (int(move.ToPos)-int(move.FromPos) == 2 || int(move.ToPos)-int(move.FromPos) == -2) {
			switch kingPos {
			case 2:
				board.Fields[0] = piece.None
				board.Fields[3] = piece.BlackRook
			case 6:
				board.Fields[7] = piece.None
				board.Fields[5] = piece.BlackRook
			case 58:
				board.Fields[56] = piece.None
				board.Fields[59] = piece.WhiteRook
			case 62:
				board.Fields[63] = piece.None
				board.Fields[61] = piece.WhiteRook
			}
		}
	}

	board.EnPassantPos = -1
	if p&piece.Pawn != piece.None && (move.ToPos-move.FromPos == Width*2 || move.FromPos-move.ToPos == Width*2) {
		board.EnPassantPos = Pos(int(move.FromPos)+int(move.ToPos)) / 2
		posX := board.EnPassantPos % Width
		opPawn := false
		if board.WhiteMove {
			if posX > 0 && board.Fields[board.EnPassantPos-Width-1] == piece.BlackPawn {
				opPawn = true
			}
			if posX < Width-1 && board.Fields[board.EnPassantPos-Width+1] == piece.BlackPawn {
				opPawn = true
			}
		} else {
			if posX > 0 && board.Fields[board.EnPassantPos+Width-1] == piece.WhitePawn {
				opPawn = true
			}
			if posX < Width-1 && board.Fields[board.EnPassantPos+Width+1] == piece.WhitePawn {
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
	if p == piece.Pawn || move.CapturePiece != piece.None {
		board.HalfmoveClock = 0
	}
	if board.WhiteMove {
		board.MoveNumber++
	}
}

func (board *YacBoard) DoMoveBackward(move Move, lastBoardInfos BoardInfo) {
	p := board.Fields[move.ToPos]
	board.Fields[move.FromPos] = p
	board.Fields[move.ToPos] = move.CapturePiece

	if move.PromotionPiece != piece.None {
		board.Fields[move.FromPos] = (p & piece.Colors) | piece.Pawn
	}

	if p&piece.Pawn != piece.None &&
		move.FromPos%Width != move.ToPos%Width &&
		move.CapturePiece == piece.None {
		if board.WhiteMove {
			board.Fields[(uint)(lastBoardInfos&EnPassantMask)-Width] = piece.WhitePawn
		} else {
			board.Fields[(uint)(lastBoardInfos&EnPassantMask)+Width] = piece.BlackPawn
		}
	}

	if p&piece.King != piece.None {
		if p == piece.WhiteKing {
			board.WhiteKingPos = Pos(move.FromPos)
		} else {
			board.BlackKingPos = Pos(move.FromPos)
		}

		posXdif := int(move.FromPos%Width) - int(move.ToPos%Width)
		if posXdif > 1 || posXdif < -1 {
			switch move.ToPos {
			case 2: // black O-O-O
				board.Fields[0] = piece.BlackRook
				board.Fields[3] = piece.None
			case 6: // black O-O
				board.Fields[7] = piece.BlackRook
				board.Fields[5] = piece.None
			case 58: // white O-O-O
				board.Fields[56] = piece.WhiteRook
				board.Fields[59] = piece.None
			case 62: // white O-O
				board.Fields[63] = piece.WhiteRook
				board.Fields[61] = piece.None
			}
		}
	}

	if board.WhiteMove {
		board.MoveNumber--
	}
	board.WhiteMove = !board.WhiteMove
	board.SetBoardInfo(lastBoardInfos)
}
