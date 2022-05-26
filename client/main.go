package main

import (
	"client/canvas"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"math"
	"math/rand"
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

const speed = 0.5
const pih = math.Pi / 180.0
const pih2 = math.Pi / 360.0

var edgeCount = 21
var grad, gra, eg float64

func LoopTick() {
	can.ResizeIfNeeded()

	can.Clear("#000")
	can.SetLineWidth(2)

	grad += speed
	gra += speed

	if grad > 360.0 {
		grad -= 360.0
		if eg == 0 {
			eg = 1
		} else {
			eg = 0
		}
	}
	if gra > 720.0 {
		gra -= 720.0
	}

	fr := int(math.Sin(pih*grad)*127) + 127
	fb := 127

	sw1 := pih * grad
	sw2 := pih2 * gra

	rnd := rand.New(rand.NewSource(12345))

	radStep := math.Pi * 2.0 / float64(edgeCount)
	widthH := float64(can.Width / 2)
	heightH := float64(can.Height / 2)
	widthK := widthH - widthH/5
	heightK := heightH - heightH/10
	for y := 0; y < edgeCount; y++ {
		x2 := math.Sin(sw1+radStep*float64(y)+sw2)*widthK + widthH
		y2 := -math.Cos(sw1+radStep*float64(y))*heightK + heightH

		for x := 0; x < y; x++ {
			x1 := math.Sin(sw1+radStep*float64(x)+sw2)*widthK + widthH
			y1 := -math.Cos(sw1+radStep*float64(x)+(sw1*eg))*heightK + heightH

			fg := rnd.Intn(256)

			can.BeginPath()
			can.SetStrokeStyle(fmt.Sprintf("rgb(%d,%d,%d)", fr, fg, fb))
			can.MoveToF(x1, y1)
			can.LineToF(x2, y2)
			can.Stroke()
		}
	}
}

func main() {
	fmt.Println("Hello World!")

	wg := js.Global().Get("window").Get("wg")
	wg.Set("test", js.FuncOf(WgTest))
	wg.Set("perfTest", js.FuncOf(WgPerfTest))

	var err error
	can, err = canvas.New("body")
	if err != nil {
		js.Global().Get("console").Call("error", err.Error())
		return
	}
	fmt.Println(can.NewSize)

	var loop js.Func
	loop = js.FuncOf(func(_ js.Value, _ []js.Value) any {
		js.Global().Call("requestAnimationFrame", loop)
		LoopTick()
		return nil
	})
	js.Global().Call("requestAnimationFrame", loop)

	<-make(chan bool, 0)
}
