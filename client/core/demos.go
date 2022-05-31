package core

import (
	"client/canvas"
	"client/core/ease"
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

	c.Save()

	var board = yacboard.New()
	fieldSize := c.Height / boardsize.Height

	c.Clear("#888")

	colors := []string{"#f0d9b5", "#b58863"}

	anipos := math.Pi / 1800.0 * float64(time.Now().UnixMilli()%3600)
	anipos /= math.Pi * 2
	anipos = ease.InOutSine(anipos)

	for y := 0; y < boardsize.Height; y++ {
		for x := 0; x < boardsize.Width; x++ {
			c.ResetTransform()
			if ms.Buttons&1 != 0 {
				c.Translate(can.Width/2, can.Height/2)
				c.Rotate(anipos * math.Pi * 2)
				c.Translate(-can.Width/2, -can.Height/2)
			}
			c.Translate(x*fieldSize+can.Width/2-fieldSize*4, y*fieldSize)
			if ms.Buttons&2 != 0 {
				c.Translate(fieldSize/2, fieldSize/2)
				c.Rotate(anipos * math.Pi * -2)
				c.Translate(fieldSize/-2, fieldSize/-2)
			}
			c.SetFillStyle(colors[(x+y)&1])
			c.FillRect(0, 0, fieldSize, fieldSize)
			svgpieces.Draw(c, 0, 0, fieldSize, board.GetField(pos.FromXY(x, y)))
		}
	}

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
