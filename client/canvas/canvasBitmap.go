package canvas

import "syscall/js"

type CanvasBitmap struct {
	CanvasContext
	CanvasBase js.Value
}

type IBitmap interface {
	GetImageElement() js.Value
}

func NewBitmap(width, height int) *CanvasBitmap {
	doc := js.Global().Get("document")
	CanvasBase := doc.Call("createElement", "canvas")
	CanvasBase.Set("width", width)
	CanvasBase.Set("height", height)
	ctx := CanvasContext{Value: CanvasBase.Call("getContext", "2d"), SizeXY: SizeXY{Width: width, Height: height}}
	ctx.ImageSmoothingEnabled(false)

	return &CanvasBitmap{
		CanvasContext: ctx,
		CanvasBase:    CanvasBase,
	}
}

func (bitmap *CanvasBitmap) GetImageElement() js.Value {
	return bitmap.CanvasBase
}

func (ctx *CanvasContext) DrawImage(bitmap IBitmap, destX, destY int) {
	ctx.Call("drawImage", bitmap.GetImageElement(), destX, destY)
}

func (ctx *CanvasContext) DrawSprite(bitmap IBitmap, sourceX, sourceY, width, height, destX, destY int) {
	ctx.Call("drawImage", bitmap.GetImageElement(), sourceX, sourceY, width, height, destX, destY, width, height)
}

func (ctx *CanvasContext) DrawSpriteResize(bitmap IBitmap, sourceX, sourceY, sourceWidth, sourceHeight, destX, destY, destWidth, destHeight int) {
	ctx.Call("drawImage", bitmap.GetImageElement(), sourceX, sourceY, sourceWidth, sourceHeight, destX, destY, destWidth, destHeight)
}
