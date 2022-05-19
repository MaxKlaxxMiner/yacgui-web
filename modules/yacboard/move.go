package yacboard

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

type Move struct {
	PromotionPiece piece.Piece
	CapturePiece   piece.Piece
	FromPos        byte
	ToPos          byte
}

func (m Move) IsValid(optionalBoard *YacBoard) bool {
	if m.FromPos == m.ToPos {
		return false
	}
	if optionalBoard == nil {
		return true
	}
	return optionalBoard.moveCheck(m)
}

func (m Move) String() string {
	if !m.IsValid(nil) {
		return "-"
	}

	result := fmt.Sprintf("%s-%s", pos.Pos(m.FromPos), pos.Pos(m.ToPos))

	if m.PromotionPiece != piece.None {
		result += "->" + m.PromotionPiece.String()
	}

	if m.CapturePiece != piece.None {
		result += " (x" + m.CapturePiece.String() + ")"
	}

	return result
}
