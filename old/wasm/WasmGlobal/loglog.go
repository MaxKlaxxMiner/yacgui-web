package WasmGlobal

import (
	"strings"
	"syscall/js"
)

// WgLoglog wg: loglog(line: string)
func WgLoglog(_ js.Value, args []js.Value) any {
	if len(args) != 1 || args[0].Type() != js.TypeString {
		return nil
	}
	line := args[0].String()

	doc := js.Global().Get("document")
	output := doc.Call("getElementById", "output")
	board := doc.Call("getElementById", "board")

	html := output.Get("innerHTML").String() + line + "<br>"

	boardHeight := board.Get("clientHeight").Int()
	outputHeight := output.Get("clientHeight").Int()

	for boardHeight-15 < outputHeight {
		i := strings.Index(html, "<br>")
		if i >= 0 {
			html = html[i+4:]
			outputHeight -= 15
		} else {
			break
		}
	}
	output.Set("innerHTML", html)
	return nil
}
