package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"glfw-client/app"
	"runtime"
	"strconv"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// Samples: https://github.com/cstegel/opengl-samples-golang

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(1792, 1008, "Hello World!", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	app.MainInit(window)

	fpsCounter := 0
	fpsTime := int(glfw.GetTime())
	for !window.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		app.MainDraw(window)

		window.SwapBuffers()

		fpsCounter++
		tim := int(glfw.GetTime())
		if tim != fpsTime {
			fpsTime = tim
			window.SetTitle("fps: " + strconv.Itoa(fpsCounter))
			fpsCounter = 0
		}
	}
}
