package canvas

const hexchars = "0123456789abcdef"

func HexColor(color int) string {
	result := make([]byte, 7)
	result[0] = '#'
	result[1] = hexchars[color>>20&0xf]
	result[2] = hexchars[color>>16&0xf]
	result[3] = hexchars[color>>12&0xf]
	result[4] = hexchars[color>>8&0xf]
	result[5] = hexchars[color>>4&0xf]
	result[6] = hexchars[color&0xf]
	return string(result)
}

func (ctx *Context) Line(x1, y1, x2, y2, color int) {
	ctx.BeginPath()
	ctx.SetStrokeStyle(HexColor(color))
	ctx.MoveTo(x1, y1)
	ctx.LineTo(x2, y2)
	ctx.Stroke()
}

func (ctx *Context) LineF(x1, y1, x2, y2 float64, color int) {
	ctx.BeginPath()
	ctx.SetStrokeStyle(HexColor(color))
	ctx.MoveToF(x1, y1)
	ctx.LineToF(x2, y2)
	ctx.Stroke()
}
