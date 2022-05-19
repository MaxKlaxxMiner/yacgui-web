package WasmGlobal

import (
	"syscall/js"
)

func WgInit() {
	wg := js.Global().Get("window").Get("wg")

	wg.Set("test", js.FuncOf(WgTest))
	wg.Set("keyDown", js.FuncOf(WgKeyDown))
	wg.Set("loglog", js.FuncOf(WgLoglog))
	wg.Set("getMoveMapFromFEN", js.FuncOf(WgGetMoveMapFromFEN))
	wg.Set("doMove", js.FuncOf(WgDoMove))

	wg.Call("ready")
}
