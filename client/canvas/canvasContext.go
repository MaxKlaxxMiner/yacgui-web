package canvas

import "syscall/js"

type CanvasContext struct {
	js.Value
	SizeXY
}

func (ctx *CanvasContext) ImageSmoothingEnabled(smooth bool) {
	// --- legacy ---
	if !ctx.Get("imageSmoothingEnabled").IsUndefined() {
		ctx.Set("imageSmoothingEnabled", smooth)
	}
	// --- internet explorer ---
	if !ctx.Get("msImageSmoothingEnabled").IsUndefined() {
		ctx.Set("msImageSmoothingEnabled", smooth)
	}
}

func (ctx *CanvasContext) SetFillStyle(style string) {
	ctx.Set("fillStyle", style)
}

func (ctx *CanvasContext) SetStrokeStyle(style string) {
	ctx.Set("strokeStyle", style)
}

func (ctx *CanvasContext) SetLineWidth(width float64) {
	ctx.Set("lineWidth", width)
}

func (ctx *CanvasContext) BeginPath() {
	ctx.Call("beginPath")
}

func (ctx *CanvasContext) Stroke() {
	ctx.Call("stroke")
}

func (ctx *CanvasContext) MoveTo(x, y int) {
	ctx.Call("moveTo", x, y)
}

func (ctx *CanvasContext) MoveToF(x, y float64) {
	ctx.Call("moveTo", x, y)
}

func (ctx *CanvasContext) LineTo(x, y int) {
	ctx.Call("lineTo", x, y)
}

func (ctx *CanvasContext) LineToF(x, y float64) {
	ctx.Call("lineTo", x, y)
}

func (ctx *CanvasContext) FillRect(x, y, w, h int) {
	ctx.Call("fillRect", x, y, w, h)
}

func (ctx *CanvasContext) Clear(colorCode string) {
	ctx.SetFillStyle(colorCode)
	ctx.FillRect(0, 0, ctx.Width, ctx.Height)
	ctx.SetFillStyle("#fff")
}

func (ctx *CanvasContext) Save() {
	ctx.Call("save")
}

func (ctx *CanvasContext) Restore() {
	ctx.Call("restore")
}

func (ctx *CanvasContext) ResetTransform() {
	ctx.Call("resetTransform")
}
