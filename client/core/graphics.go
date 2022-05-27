package core

import (
	"client/canvas"
	"syscall/js"
)

var can canvas.Canvas

func InitCanvas() {
	var err error
	can, err = canvas.New("body")
	if err != nil {
		js.Global().Get("console").Call("error", err.Error())
		return
	}
}
