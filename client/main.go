package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
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

func main() {
	fmt.Println("Hello World!")

	wg := js.Global().Get("window").Get("wg")
	wg.Set("test", js.FuncOf(WgTest))
	wg.Set("perfTest", js.FuncOf(WgPerfTest))

	c := make(chan bool, 0)
	<-c
}
