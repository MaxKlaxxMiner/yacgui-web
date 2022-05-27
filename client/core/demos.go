package core

import (
	"client/canvas"
	"client/keys"
	"client/lineDemo"
	"client/mouse"
	"syscall/js"
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

func (demo *MouseDemo) TickUpdate(c *canvas.CanvasContext, k *keys.Keys) {
	c.Clear("#000")
	c.Line(ms.X, ms.Y, ms.X+100, ms.Y+100, 0x0080ff)
}

func RunMouseDemo() {
	//m := MouseDemo{}
	ms.EventCallback = append(ms.EventCallback, func(m *mouse.Mouse) {
		can.Clear("#000")
		can.Line(ms.X, ms.Y, ms.X+100, ms.Y+100, 0x0080ff)
	})
	//RunMainLoop(&m)
}
