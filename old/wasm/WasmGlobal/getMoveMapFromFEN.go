package WasmGlobal

import (
	. "github.com/MaxKlaxxMiner/yacgui-web/YacBoard"
	"strings"
	"syscall/js"
)

func WgGetMoveMapFromFEN(_ js.Value, args []js.Value) any {
	if len(args) != 1 || args[0].Type() != js.TypeString {
		return nil
	}

	fen := args[0].String()

	var board YacBoard
	err := board.SetFEN(fen)
	if err != nil {
		return nil
	}

	var moves [256]Move
	moveCount := int(board.GetMoves(&moves))

	result := ""
	lastFrom := ""

	for m := 0; m < moveCount; m++ {
		if moves[m].PromotionPiece != PieceNone && moves[m].PromotionPiece&Queen != Queen {
			continue
		}
		from := PosToChars(int(moves[m].FromPos))
		to := PosToChars(int(moves[m].ToPos))
		if from != lastFrom {
			lastFrom = from
			result += "," + from + ","
		}
		result += "|" + to
	}

	if len(result) > 0 {
		return strings.ReplaceAll(result[1:], ",|", ",")
	}
	return ""
}
