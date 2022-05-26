package main

import (
	"client/canvas"
	"client/keys"
	"client/lineDemo"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"syscall/js"
)

var can canvas.Canvas

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

func InitCanvas() {
	var err error
	can, err = canvas.New("body")
	if err != nil {
		js.Global().Get("console").Call("error", err.Error())
		return
	}
}

func RunLineDemo() {
	k := keys.New()
	demo := lineDemo.New()
	var loop js.Func
	loop = js.FuncOf(func(_ js.Value, _ []js.Value) any {
		js.Global().Call("requestAnimationFrame", loop)
		can.ResizeIfNeeded()
		demo.TickUpdate(&can.CanvasContext, k)
		return nil
	})
	js.Global().Call("requestAnimationFrame", loop)
}

func main() {
	fmt.Println("Hello World!")

	wg := js.Global().Get("window").Get("wg")
	wg.Set("test", js.FuncOf(WgTest))
	wg.Set("perfTest", js.FuncOf(WgPerfTest))

	InitCanvas()
	RunLineDemo()

	<-make(chan bool, 0)
}
