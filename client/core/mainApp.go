package core

import (
	"client/jscore/canvas"
	"client/jscore/keys"
)

type MainApp struct {
}

type MouseTimeInfo struct {
	canvas.PosXY
	tick int
}

var msLastPos canvas.PosXY
var lastM canvas.PosXY
var tickCounter = 0
var msTimeList = make([]MouseTimeInfo, 2)

const msTimeThreshold = 3

func (m *MainApp) TickUpdate(c *canvas.Context, k *keys.Keys) {
	//colors := []string{"#f0d9b5", "#b58863"}
	tickCounter++

	if ms.PosXY != msLastPos || ms.PosXY != lastM {
		if ms.PosXY != msLastPos {
			copy(msTimeList[1:], msTimeList)
			msTimeList[0] = MouseTimeInfo{ms.PosXY, tickCounter}
		}

		can.Context.SetFillStyle("#000")
		can.Context.FillRect(lastM.X, lastM.Y, 100, 100)
		lastM = ms.PosXY
		if tickCounter == msTimeList[len(msTimeList)-1].tick+len(msTimeList)-1 {
			addX := msTimeList[0].X - msTimeList[len(msTimeList)-1].X
			addY := msTimeList[0].Y - msTimeList[len(msTimeList)-1].Y
			if addX > msTimeThreshold || addY > msTimeThreshold || addX < -msTimeThreshold || addY < -msTimeThreshold {
				lastM.X += addX * 100 / 100
				lastM.Y += addY * 100 / 100
			}
		}
		can.Context.SetFillStyle("#f0d9b5")
		//can.Context.SetFillStyle("#b58863")
		can.Context.FillRect(lastM.X, lastM.Y, 100, 100)

		msLastPos = ms.PosXY
	}
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

	//ms.EventCallback = append(ms.EventCallback, func(mouse *mouse.Mouse) {
	//	if ms.PosXY != lastM {
	//		can.Context.SetFillStyle("#000")
	//		can.Context.FillRect(lastM.X, lastM.Y, 100, 100)
	//		can.Context.SetFillStyle("#f0d9b5")
	//		can.Context.FillRect(ms.X, ms.Y, 100, 100)
	//		lastM = ms.PosXY
	//	}
	//})
	RunMainLoop(&m)
}
