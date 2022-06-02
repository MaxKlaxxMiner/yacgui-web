package core

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"strconv"
	"syscall/js"
)

func WgTest(_ js.Value, _ []js.Value) any {
	fmt.Println("wasm: test()")
	return nil
}

func WgPerfTest(_ js.Value, args []js.Value) any {
	trim := 3
	if len(args) > 0 && args[0].Type() == js.TypeNumber {
		trim = args[0].Int()
	}
	yacboard.PerftTest(trim)
	return nil
}

func InitWg() {
	wg := js.Global().Get("window").Get("wg")
	wg.Set("test", js.FuncOf(WgTest))
	wg.Set("perfTest", js.FuncOf(WgPerfTest))
	fmt.Println(strconv.IntSize)
}
