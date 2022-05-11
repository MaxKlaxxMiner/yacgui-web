package WasmGlobal

import (
	"syscall/js"
)

// WgKeyDown wg: keyDown(key: string, defaultPrevented: boolean): boolean
func WgKeyDown(_ js.Value, args []js.Value) any {
	if len(args) != 2 || args[0].Type() != js.TypeString || args[1].Type() != js.TypeBoolean {
		return false
	}
	code := args[0].String()
	defaultPrevented := args[1].Bool()

	if defaultPrevented {
		return false
	}

	if code == "KeyF" {
		js.Global().Get("ground").Call("toggleOrientation")
		return true
	}

	return false
}
