package canvas

import (
	"errors"
	"syscall/js"
)

type Canvas struct {
	OuterHtml  js.Value
	CanvasHtml js.Value
	NewSize    SizeXY
	CanvasContext
}

func New(elementSelector string) (Canvas, error) {
	var c Canvas

	doc := js.Global().Get("document")

	// --- search html-element ---
	c.OuterHtml = doc.Call("querySelector", elementSelector)
	if c.OuterHtml.IsNull() {
		return c, errors.New("html-element not found: " + elementSelector)
	}

	// --- create canvas-container & insert at first ---
	c.CanvasHtml = doc.Call("createElement", "canvas")
	c.OuterHtml.Call("insertBefore", c.CanvasHtml, c.OuterHtml.Get("childNodes").Index(0))

	// --- create (2D) render-context ---
	c.CanvasContext = CanvasContext{c.CanvasHtml.Call("getContext", "2d"), SizeXY{}}
	if c.CanvasContext.IsNull() {
		return c, errors.New("error creating CanvasRenderingContext2D")
	}

	// --- resize canvas ---
	c.ResizeIfNeeded()

	// --- default settings ---
	c.CanvasContext.ImageSmoothingEnabled(false)

	return c, nil
}
