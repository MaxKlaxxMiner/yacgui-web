package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"glfw-client/app"
	"runtime"
	"strconv"
	"time"
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

	lastFrameTime := time.Now()
	lastActionTime := time.Now()
	fpsCounter := 0
	fpsTime := int(glfw.GetTime())
	for !window.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		refresh := time.Since(lastFrameTime) > 100*time.Millisecond

		mx, my := window.GetCursorPos()
		if mx != app.MouseX || my != app.MouseY {
			app.MouseX = mx
			app.MouseY = my
			lastActionTime = time.Now()
			refresh = true
		}

		if refresh {
			app.MainDraw(window)
			window.SwapBuffers()
			lastFrameTime = time.Now()
			fpsCounter++
			tim := int(glfw.GetTime())
			if tim != fpsTime {
				fpsTime = tim
				window.SetTitle("fps: " + strconv.Itoa(fpsCounter))
				fpsCounter = 0
			}
		} else {
			since := time.Since(lastActionTime)
			if since > 100*time.Millisecond {
				if since < 500*time.Millisecond {
					time.Sleep(time.Millisecond)
				} else {
					time.Sleep(time.Millisecond * 20)
				}
			}
		}
	}
}
