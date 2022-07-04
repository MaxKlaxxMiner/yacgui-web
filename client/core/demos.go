package core

import (
	"client/core/ease"
	"client/jscore/canvas"
	"client/jscore/keys"
	"client/jscore/mouse"
	"client/lineDemo"
	"client/svgpieces"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"math"
	"syscall/js"
	"time"
)

type TickInterface interface {
	TickUpdate(c *canvas.Context, k *keys.Keys)
}

func RunMainLoop(demo TickInterface) {
	var loop js.Func
	loop = js.FuncOf(func(_ js.Value, _ []js.Value) any {
		js.Global().Call("requestAnimationFrame", loop)
		can.ResizeIfNeeded()
		demo.TickUpdate(&can.Context, ks)
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

func Draw(c *canvas.Context) {
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

func (demo *MouseDemo) TickUpdate(c *canvas.Context, k *keys.Keys) {
	Draw(&can.Context)
}

func RunMouseDemo() {
	m := MouseDemo{}
	ms.EventCallback = append(ms.EventCallback, func(m *mouse.Mouse) {
		//Draw(&can.CanvasContext)
	})
	RunMainLoop(&m)
}
