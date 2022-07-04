package canvas

import "syscall/js"

type Context struct {
	js.Value
	SizeXY
}

func (ctx *Context) ImageSmoothingEnabled(smooth bool) {
	// --- legacy ---
	if !ctx.Get("imageSmoothingEnabled").IsUndefined() {
		ctx.Set("imageSmoothingEnabled", smooth)
	}
	// --- internet explorer ---
	if !ctx.Get("msImageSmoothingEnabled").IsUndefined() {
		ctx.Set("msImageSmoothingEnabled", smooth)
	}
}

func (ctx *Context) SetFillStyle(style string) {
	ctx.Set("fillStyle", style)
}

func (ctx *Context) SetStrokeStyle(style string) {
	ctx.Set("strokeStyle", style)
}

func (ctx *Context) SetLineCap(cap string) {
	ctx.Set("lineCap", cap)
}

func (ctx *Context) SetLineWidth(width float64) {
	ctx.Set("lineWidth", width)
}

func (ctx *Context) SetMiterLimit(limit float64) {
	ctx.Set("miterLimit", limit)
}

func (ctx *Context) BeginPath() {
	ctx.Call("beginPath")
}

func (ctx *Context) Stroke() {
	ctx.Call("stroke")
}

func (ctx *Context) StrokePath(path js.Value) {
	ctx.Call("stroke", path)
}

func (ctx *Context) MoveTo(x, y int) {
	ctx.Call("moveTo", x, y)
}

func (ctx *Context) MoveToF(x, y float64) {
	ctx.Call("moveTo", x, y)
}

func (ctx *Context) LineTo(x, y int) {
	ctx.Call("lineTo", x, y)
}

func (ctx *Context) LineToF(x, y float64) {
	ctx.Call("lineTo", x, y)
}

func (ctx *Context) FillPath(path js.Value) {
	ctx.Call("fill", path)
}

func (ctx *Context) FillRect(x, y, w, h int) {
	ctx.Call("fillRect", x, y, w, h)
}

func (ctx *Context) Clear(colorCode string) {
	ctx.SetFillStyle(colorCode)
	ctx.FillRect(0, 0, ctx.Width, ctx.Height)
	ctx.SetFillStyle("#fff")
}

func (ctx *Context) Save() {
	ctx.Call("save")
}

func (ctx *Context) Restore() {
	ctx.Call("restore")
}

func (ctx *Context) ResetTransform() {
	ctx.Call("resetTransform")
}

func (ctx *Context) Translate(x, y int) {
	ctx.Call("translate", x, y)
}

func (ctx *Context) TranslateF(x, y float64) {
	ctx.Call("translate", x, y)
}

func (ctx *Context) Rotate(rad float64) {
	ctx.Call("rotate", rad)
}

func (ctx *Context) Scale(x, y int) {
	ctx.Call("scale", x, y)
}

func (ctx *Context) ScaleF(x, y float64) {
	ctx.Call("scale", x, y)
}
