package core

import (
	"client/canvas"
	"client/keys"
	"client/lineDemo"
	"client/mouse"
	"client/svgpieces"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
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
	can.Clear("#888")
	can.Save()

	can.ResetTransform()
	can.Translate(ms.X, ms.Y)
	can.Scale(3, 3)
	can.TranslateF(-45.0/2, -45.0/2)
	svgpieces.Draw(c, piece.WhitePawn)
	can.Translate(0, 45)
	svgpieces.Draw(c, piece.BlackPawn)
	can.Translate(45, -45)
	svgpieces.Draw(c, piece.WhiteBishop)
	can.Translate(0, 45)
	svgpieces.Draw(c, piece.BlackBishop)
	can.Translate(45, -45)
	svgpieces.Draw(c, piece.WhiteKnight)
	can.Translate(0, 45)
	svgpieces.Draw(c, piece.BlackKnight)
	can.Translate(45, -45)
	svgpieces.Draw(c, piece.WhiteRook)
	can.Translate(0, 45)
	svgpieces.Draw(c, piece.BlackRook)

	can.Restore()
}

func RunMouseDemo() {
	m := MouseDemo{}
	ms.EventCallback = append(ms.EventCallback, func(m *mouse.Mouse) {
		//	draw()
	})
	RunMainLoop(&m)
}
