package core

import (
	"client/canvas"
	"client/keys"
	"client/lineDemo"
	"client/mouse"
	"client/svgpieces"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
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

func Draw(c *canvas.CanvasContext) {
	var board = yacboard.New()

	c.Save()

	c.Clear("#888")

	fieldSize := c.Height / boardsize.Height
	colors := []string{"#f0d9b5", "#b58863"}
	scale := float64(fieldSize) / 45
	for y := 0; y < boardsize.Height; y++ {
		for x := 0; x < boardsize.Width; x++ {
			c.ResetTransform()
			if ms.Buttons&1 != 0 {
				c.Translate(can.Width/2, can.Height/2)
				c.Rotate(math.Pi / 1800.0 * float64(time.Now().UnixMilli()%3600))
				c.Translate(-can.Width/2, -can.Height/2)
			}
			c.Translate(x*fieldSize+can.Width/2-fieldSize*4, y*fieldSize)
			if ms.Buttons&2 != 0 {
				c.Translate(fieldSize/2, fieldSize/2)
				c.Rotate(math.Pi / -1800.0 * float64(time.Now().UnixMilli()%3600))
				c.Translate(fieldSize/-2, fieldSize/-2)
			}
			c.SetFillStyle(colors[(x+y)&1])
			c.FillRect(0, 0, fieldSize, fieldSize)
			c.ScaleF(scale, scale)
			svgpieces.Draw(c, board.GetField(pos.FromXY(x, y)))
		}
	}
	//c.SetFillStyle("#")

	//c.ResetTransform()
	//c.Translate(ms.X, ms.Y)
	//c.Scale(3, 3)
	//c.TranslateF(-45.0/2, -45.0/2)
	//svgpieces.Draw(c, piece.WhitePawn)
	//c.Translate(0, 45)
	//svgpieces.Draw(c, piece.BlackPawn)
	//c.Translate(45, -45)
	//svgpieces.Draw(c, piece.WhiteBishop)
	//c.Translate(0, 45)
	//svgpieces.Draw(c, piece.BlackBishop)
	//c.Translate(45, -45)
	//svgpieces.Draw(c, piece.WhiteKnight)
	//c.Translate(0, 45)
	//svgpieces.Draw(c, piece.BlackKnight)
	//c.Translate(45, -45)
	//svgpieces.Draw(c, piece.WhiteRook)
	//c.Translate(0, 45)
	//svgpieces.Draw(c, piece.BlackRook)
	//c.Translate(45, -45)
	//svgpieces.Draw(c, piece.WhiteQueen)
	//c.Translate(0, 45)
	//svgpieces.Draw(c, piece.BlackQueen)
	//c.Translate(45, -45)
	//svgpieces.Draw(c, piece.WhiteKing)
	//c.Translate(0, 45)
	//svgpieces.Draw(c, piece.BlackKing)

	c.Restore()
}

func (demo *MouseDemo) TickUpdate(c *canvas.CanvasContext, k *keys.Keys) {
	Draw(&can.CanvasContext)
}

func RunMouseDemo() {
	m := MouseDemo{}
	ms.EventCallback = append(ms.EventCallback, func(m *mouse.Mouse) {
		//Draw(&can.CanvasContext)
	})
	RunMainLoop(&m)
}
