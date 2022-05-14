package WasmGlobal

import (
	"fmt"
	. "github.com/MaxKlaxxMiner/yacgui-web/YacBoard"
	"syscall/js"
)

// WgDoMove wg: doMove(fen: string, from: Key, to: Key): string
func WgDoMove(_ js.Value, args []js.Value) any {
	if len(args) != 3 || args[0].Type() != js.TypeString || args[1].Type() != js.TypeString || args[2].Type() != js.TypeString {
		return nil
	}

	fen := args[0].String()
	from := args[1].String()
	to := args[2].String()
	fmt.Println("blub2: " + fen)

	var board YacBoard
	err := board.SetFEN(fen)
	if err != nil {
		return nil
	}

	var moves [256]Move
	moveCount := int(board.GetMoves(&moves))

	for m := 0; m < moveCount; m++ {
		mFrom := PosToChars(int(moves[m].FromPos))
		mTo := PosToChars(int(moves[m].ToPos))
		if mFrom == from && mTo == to {
			board.DoMove(moves[m])
			return board.GetFEN()
		}
	}

	return nil
}
