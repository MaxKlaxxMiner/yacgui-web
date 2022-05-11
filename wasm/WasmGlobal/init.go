package WasmGlobal

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func WgInit() {
	wg := js.Global().Get("window").Get("wg")

	wg.Set("test", js.FuncOf(WgTest))
	wg.Set("keyDown", js.FuncOf(WgKeyDown))
	wg.Set("loglog", js.FuncOf(WgLoglog))
	wg.Set("mm", js.FuncOf(WgMm))

	wg.Call("ready")
}

func WgMm(_ js.Value, _ []js.Value) any {

	fmt.Println(strconv.IntSize)

	return nil
}
