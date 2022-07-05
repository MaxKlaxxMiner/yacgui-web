package main

import (
	"runtime"
	"vulkan-client/app"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	a := app.New()
	if err := a.Run(); err != nil {
		panic(err)
	}
}
