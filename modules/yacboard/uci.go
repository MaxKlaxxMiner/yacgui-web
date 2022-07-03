package yacboard

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
)

func (m Move) Uci() string {
	if m.PromotionPiece != piece.None {
		return pos.Pos(m.FromPos).String() + pos.Pos(m.ToPos).String() + m.PromotionPiece.String()
	}
	return pos.Pos(m.FromPos).String() + pos.Pos(m.ToPos).String()
}

func (board *YacBoard) GetUciMove(uciMove string) Move {
	switch {
	case uciMove == "e1h1" && board.WhiteCanCastleKingside:
		uciMove = "e1g1"
	case uciMove == "e1a1" && board.WhiteCanCastleQueenside:
		uciMove = "e1c1"
	case uciMove == "e8h8" && board.BlackCanCastleKingside:
		uciMove = "e8g8"
	case uciMove == "e8a8" && board.BlackCanCastleQueenside:
		uciMove = "e8c8"
	}
	moves := board.GetMoves()
	for _, m := range moves {
		if m.Uci() == uciMove {
			return m
		}
	}
	return Move{}
}

func (board *YacBoard) DoUciMove(uciMove ...string) bool {
	for _, uci := range uciMove {
		m := board.GetUciMove(uci)
		if !m.IsValid(nil) {
			return false
		}
		board.DoMove(m)
	}
	return true
}
