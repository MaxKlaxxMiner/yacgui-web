package canvas

import "syscall/js"

type Bitmap struct {
	Context
	CanvasBase js.Value
}

type IBitmap interface {
	GetImageElement() js.Value
}

func NewBitmap(width, height int) *Bitmap {
	doc := js.Global().Get("document")
	CanvasBase := doc.Call("createElement", "canvas")
	CanvasBase.Set("width", width)
	CanvasBase.Set("height", height)
	ctx := Context{Value: CanvasBase.Call("getContext", "2d"), SizeXY: SizeXY{Width: width, Height: height}}
	ctx.ImageSmoothingEnabled(false)

	return &Bitmap{
		Context:    ctx,
		CanvasBase: CanvasBase,
	}
}

func (bitmap *Bitmap) GetImageElement() js.Value {
	return bitmap.CanvasBase
}

func (ctx *Context) DrawImage(bitmap IBitmap, destX, destY int) {
	ctx.Call("drawImage", bitmap.GetImageElement(), destX, destY)
}

func (ctx *Context) DrawSprite(bitmap IBitmap, sourceX, sourceY, width, height, destX, destY int) {
	ctx.Call("drawImage", bitmap.GetImageElement(), sourceX, sourceY, width, height, destX, destY, width, height)
}

func (ctx *Context) DrawSpriteResize(bitmap IBitmap, sourceX, sourceY, sourceWidth, sourceHeight, destX, destY, destWidth, destHeight int) {
	ctx.Call("drawImage", bitmap.GetImageElement(), sourceX, sourceY, sourceWidth, sourceHeight, destX, destY, destWidth, destHeight)
}
