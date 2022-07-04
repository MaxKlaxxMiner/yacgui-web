package canvas

type Window struct {
	RectXY
	Transparent bool
	Alive       bool
	Focus       int
	Draw        func(ctx *Context, that *Window, drawRect RectXY)
}

type IWindow interface {
	GetWindow() *Window
}
