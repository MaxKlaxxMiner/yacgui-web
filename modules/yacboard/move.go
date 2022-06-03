package yacboard

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	. "github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

type Move struct {
	PromotionPiece piece.Piece
	CapturePiece   piece.Piece
	FromPosF       byte
	ToPosF         byte
}

func (m Move) IsValid(optionalBoard *YacBoard) bool {
	if m.FromPosF == m.ToPosF {
		return false
	}
	if optionalBoard == nil {
		return true
	}
	return optionalBoard.MoveCheck(m)
}

func (m Move) String() string {
	if !m.IsValid(nil) {
		return "-"
	}

	result := fmt.Sprintf("%s-%s", Pos(FToPb(m.FromPosF)), Pos(FToPb(m.ToPosF)))
	if m.CapturePiece != piece.None {
		result = fmt.Sprintf("%sx%s", Pos(FToPb(m.FromPosF)), Pos(FToPb(m.ToPosF)))
	}

	if m.PromotionPiece != piece.None {
		result += "->" + m.PromotionPiece.String()
	}

	if m.CapturePiece != piece.None {
		result += " (x" + m.CapturePiece.String() + ")"
	}

	return result
}
