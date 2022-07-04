package canvas

import "syscall/js"

type FixImage struct {
	SizeXY
	ImageInstance js.Value
}

func NewFixImage(bitmap *Bitmap) *FixImage {
	result := FixImage{
		SizeXY:        bitmap.SizeXY,
		ImageInstance: js.Global().Get("Image").New(),
	}
	result.ImageInstance.Set("src", bitmap.CanvasBase.Call("toDataURL", "image/png"))
	return &result
}

func (image *FixImage) GetImageElement() js.Value {
	return image.ImageInstance
}
