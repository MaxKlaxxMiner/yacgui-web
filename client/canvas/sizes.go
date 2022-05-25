package canvas

import "syscall/js"

type SizeXY struct {
	Width  int
	Height int
}

func GetWindowSize() SizeXY {
	w := js.Global()

	return SizeXY{
		Width:  w.Get("innerWidth").Int(),
		Height: w.Get("innerHeight").Int(),
	}
}

func (c *Canvas) ResizeIfNeeded() bool {
	c.NewSize = GetWindowSize()

	if c.NewSize != c.CanvasContext.SizeXY {
		c.CanvasHtml.Set("width", c.NewSize.Width)
		c.CanvasHtml.Set("height", c.NewSize.Height)
		c.CanvasContext.SizeXY = c.NewSize
	}

	return false
}
