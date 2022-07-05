package app

import "github.com/go-gl/glfw/v3.3/glfw"

func KeyboardCallBack(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
