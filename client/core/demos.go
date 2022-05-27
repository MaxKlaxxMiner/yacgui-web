package core

import (
	"client/canvas"
	"client/keys"
	"client/lineDemo"
	"client/mouse"
	"client/svgpieces"
	"math"
	"syscall/js"
	"time"
)

type DemoInterface interface {
	TickUpdate(c *canvas.CanvasContext, k *keys.Keys)
}

func RunMainLoop(demo DemoInterface) {
	var loop js.Func
	loop = js.FuncOf(func(_ js.Value, _ []js.Value) any {
		js.Global().Call("requestAnimationFrame", loop)
		can.ResizeIfNeeded()
		demo.TickUpdate(&can.CanvasContext, ks)
		return nil
	})
	js.Global().Call("requestAnimationFrame", loop)
}

func RunLineDemo() {
	demo := lineDemo.New()
	RunMainLoop(&demo)
}

type MouseDemo struct {
}

var whitePawnPath = js.Global().Get("Path2D").New(svgpieces.WhitePawn)

func (demo *MouseDemo) TickUpdate(c *canvas.CanvasContext, k *keys.Keys) {
	can.Clear("#888")
	can.Save()

	can.ResetTransform()
	can.Translate(ms.X, ms.Y)
	can.Scale(3, 3)
	can.Rotate(math.Pi / -1800.0 * float64(time.Now().UnixMilli()%3600))
	can.TranslateF(-45.0/2, -45.0/2)

	can.SetFillStyle("#fff")
	can.FillPath(whitePawnPath)

	can.SetStrokeStyle("#000")
	can.StrokePath(whitePawnPath)

	can.Restore()
}

func RunMouseDemo() {
	m := MouseDemo{}
	ms.EventCallback = append(ms.EventCallback, func(m *mouse.Mouse) {
		//	draw()
	})
	RunMainLoop(&m)
}
