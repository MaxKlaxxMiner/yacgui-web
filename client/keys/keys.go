package keys

import (
	"syscall/js"
)

type Keys struct {
	pressed     map[string]bool
	LastPressed string
}

func New() *Keys {
	keys := Keys{pressed: make(map[string]bool)}

	js.Global().Call("addEventListener", "keydown", js.FuncOf(func(_ js.Value, args []js.Value) any {
		code := args[0].Get("code").String()
		keys.pressed[code] = true
		keys.LastPressed = code
		//args[0].Call("preventDefault")
		return nil
	}))

	js.Global().Call("addEventListener", "keyup", js.FuncOf(func(_ js.Value, args []js.Value) any {
		code := args[0].Get("code").String()
		keys.pressed[code] = false
		return nil
	}))

	return &keys
}

func (k *Keys) Pressed(checkKeys ...string) bool {
	for _, checkKey := range checkKeys {
		if k.pressed[checkKey] {
			return true
		}
	}
	return false
}

func (k *Keys) Release(releaseKeys ...string) {
	for _, releaseKey := range releaseKeys {
		if k.pressed[releaseKey] {
			k.pressed[releaseKey] = false
		}
	}
}
