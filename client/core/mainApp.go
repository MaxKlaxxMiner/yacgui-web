package core

import (
	"client/jscore/canvas"
	"client/jscore/keys"
	"client/jscore/mouse"
)

type MainApp struct {
}

func (m *MainApp) TickUpdate(c *canvas.Context, k *keys.Keys) {
	//colors := []string{"#f0d9b5", "#b58863"}
	//c.SetFillStyle("#000")
	//c.FillRect(ms.X-10, ms.Y-10, 100+20, 100+20)
	//c.SetFillStyle("#f0d9b5")
	//c.FillRect(ms.X, ms.Y, 100, 100)

	can.Context.SetFillStyle("#000")
	can.Context.FillRect(ms.X-10, ms.Y-10+150, 100+20, 100+20)
	can.Context.SetFillStyle("#f0d9b5")
	can.Context.FillRect(ms.X, ms.Y+150, 100, 100)
}

func MainAppStart() {
	m := MainApp{}

	//var loop js.Func
	//loop = js.FuncOf(func(_ js.Value, _ []js.Value) any {
	//	js.Global().Call("requestAnimationFrame", loop)

	//js.Global().Call("setInterval", js.FuncOf(func(_ js.Value, _ []js.Value) any {
	//	can.Context.SetFillStyle("#000")
	//	can.Context.FillRect(ms.X-10, ms.Y-10, 100+20, 100+20)
	//	can.Context.SetFillStyle("#f88")
	//	can.Context.FillRect(ms.X, ms.Y, 100, 100)
	//	return nil
	//}), 1)

	ms.EventCallback = append(ms.EventCallback, func(mouse *mouse.Mouse) {
		can.Context.SetFillStyle("#000")
		can.Context.FillRect(ms.X-10, ms.Y-10, 100+20, 100+20)
		can.Context.SetFillStyle("#f0d9b5")
		can.Context.FillRect(ms.X, ms.Y, 100, 100)

		//m.TickUpdate(&can.Context, ks)
	})
	RunMainLoop(&m)
}
