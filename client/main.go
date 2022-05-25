package main

import (
	"client/canvas"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"syscall/js"
	"time"
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

	c, err := canvas.New("body")
	if err != nil {
		js.Global().Get("console").Call("error", err.Error())
		return
	}
	fmt.Println(c.NewSize)

	go func() {
		for {
			const slower = 5
			colorPos := int((time.Now().UnixMilli() % (1024 * slower)) / slower)
			if colorPos > 512 {
				colorPos = 1024 - colorPos
			}

			if colorPos < 256 {
				c.Clear(fmt.Sprintf("rgb(%d,%d,%d)", 0, colorPos/2, colorPos))
			} else {
				c.Clear(fmt.Sprintf("rgb(%d,%d,%d)", colorPos-256, colorPos/2, 255))
			}

			c.SetStrokeStyle("#fff")
			c.SetLineWidth(1)
			c.BeginPath()
			c.MoveTo(c.Width/2, 0)
			c.LineTo(colorPos*c.Width/512, c.Height/2)
			c.LineTo(c.Width/2, c.Height)
			c.Stroke()

			time.Sleep(16666 * time.Microsecond)
		}
	}()

	<-make(chan bool, 0)
}
