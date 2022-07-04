package mouse

import (
	"client/jscore/canvas"
	"strconv"
	"syscall/js"
)

type Mouse struct {
	canvas.PosXY
	Wheel         int
	Buttons       int
	EventCallback []func(m *Mouse)
}

func New() *Mouse {
	var startPos canvas.PosXY
	sStorage := js.Global().Get("sessionStorage")
	if mx := sStorage.Get("mouseX"); !mx.IsUndefined() {
		startPos.X, _ = strconv.Atoi(mx.String())
	}
	if my := sStorage.Get("mouseY"); !my.IsUndefined() {
		startPos.Y, _ = strconv.Atoi(my.String())
	}

	mouse := Mouse{
		PosXY:         startPos,
		EventCallback: []func(m *Mouse){},
	}

	js.Global().Call("addEventListener", "contextmenu", js.FuncOf(func(_ js.Value, args []js.Value) any {
		args[0].Call("preventDefault")
		return nil
	}))

	mouseEvent := js.FuncOf(func(_ js.Value, args []js.Value) any {
		m := args[0]
		mouse.X = m.Get("x").Int()
		mouse.Y = m.Get("y").Int()
		sStorage.Set("mouseX", mouse.X)
		sStorage.Set("mouseY", mouse.Y)
		mouse.Buttons = m.Get("buttons").Int()
		if mouse.Buttons&8|16 != 0 { // supress browser back/forward
			args[0].Call("preventDefault")
		}
		for _, call := range mouse.EventCallback {
			call(&mouse)
		}
		return nil
	})

	wheelEvent := js.FuncOf(func(_ js.Value, args []js.Value) any {
		m := args[0]
		if !m.Get("wheelDelta").IsUndefined() { // Internet Explorer
			if m.Get("wheelDelta").Int() < 0 {
				mouse.Wheel++
			} else {
				mouse.Wheel--
			}
			for _, call := range mouse.EventCallback {
				call(&mouse)
			}
			return nil
		}
		if !m.Get("deltaY").IsUndefined() { // Legacy
			if m.Get("deltaY").Int() > 0 {
				mouse.Wheel++
			} else {
				mouse.Wheel--
			}
		} else {
			if m.Get("detail").Int() > 0 { // Firefox
				mouse.Wheel++
			} else {
				mouse.Wheel--
			}
		}
		for _, call := range mouse.EventCallback {
			call(&mouse)
		}
		return nil
	})

	js.Global().Get("document").Get("body").Set("onmousedown", mouseEvent)
	js.Global().Get("document").Get("body").Set("onmousemove", mouseEvent)
	js.Global().Get("document").Get("body").Set("onmouseup", mouseEvent)

	if !js.Global().Get("window").Get("onmousewheel").IsUndefined() { // Legacy
		js.Global().Get("window").Set("onmousewheel", wheelEvent)
	} else { // Firefox
		js.Global().Call("addEventListener", "DOMMouseScroll", wheelEvent, false)
	}

	return &mouse
}
