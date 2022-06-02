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

func (ctx *CanvasContext) SetLineCap(cap string) {
	ctx.Set("lineCap", cap)
}

func (ctx *CanvasContext) SetLineWidth(width float64) {
	ctx.Set("lineWidth", width)
}

func (ctx *CanvasContext) SetMiterLimit(limit float64) {
	ctx.Set("miterLimit", limit)
}

func (ctx *CanvasContext) BeginPath() {
	ctx.Call("beginPath")
}

func (ctx *CanvasContext) Stroke() {
	ctx.Call("stroke")
}

func (ctx *CanvasContext) StrokePath(path js.Value) {
	ctx.Call("stroke", path)
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

func (ctx *CanvasContext) FillPath(path js.Value) {
	ctx.Call("fill", path)
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

func (ctx *CanvasContext) Translate(x, y int) {
	ctx.Call("translate", x, y)
}

func (ctx *CanvasContext) TranslateF(x, y float64) {
	ctx.Call("translate", x, y)
}

func (ctx *CanvasContext) Rotate(rad float64) {
	ctx.Call("rotate", rad)
}

func (ctx *CanvasContext) Scale(x, y int) {
	ctx.Call("scale", x, y)
}

func (ctx *CanvasContext) ScaleF(x, y float64) {
	ctx.Call("scale", x, y)
}
