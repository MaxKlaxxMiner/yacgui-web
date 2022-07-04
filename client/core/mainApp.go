package core

import (
	"client/jscore/canvas"
	"client/jscore/keys"
	"client/jscore/mouse"
	"math"
)

type MainApp struct {
}

var lastM = ms.PosXY

func (m *MainApp) TickUpdate(c *canvas.Context, k *keys.Keys) {
	//colors := []string{"#f0d9b5", "#b58863"}

	const factor = 3

	cx := int(math.Round(float64(ms.X-lastM.X) * factor))
	cy := int(math.Round(float64(ms.Y-lastM.Y) * factor))

	cur := canvas.PosXY{X: ms.X + cx, Y: ms.Y + cy}
	//cur := ms.PosXY

	can.Context.SetFillStyle("#000")
	can.Context.FillRect(cur.X-10, cur.Y-10, 100+20, 100+20)
	can.Context.SetFillStyle("#f0d9b5")
	can.Context.FillRect(cur.X, cur.Y, 100, 100)
	lastM = ms.PosXY
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
		can.Context.FillRect(ms.X-10, ms.Y-10+120, 100+20, 100+20)
		can.Context.SetFillStyle("#f0d9b5")
		can.Context.FillRect(ms.X, ms.Y+120, 100, 100)

		//m.TickUpdate(&can.Context, ks)
	})
	RunMainLoop(&m)
}
