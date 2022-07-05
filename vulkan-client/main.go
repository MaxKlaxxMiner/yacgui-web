package main

import (
	"os"
	"runtime"
	"vulkan-client/app"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	profile := os.Getenv("PROFILE")

	enableValidationLayers := profile != "prod"

	a := app.New(app.Config{EnableValidationLayers: enableValidationLayers, ValidationLayers: []string{
		"VK_LAYER_KHRONOS_validation\x00",
	}})

	if err := a.Run(); err != nil {
		panic(err)
	}
}
