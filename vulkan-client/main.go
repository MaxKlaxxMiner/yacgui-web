package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
	"runtime"
	"strconv"
	"vulkan-client/app"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// Samples: https://github.com/cstegel/opengl-samples-golang

var fpsTxt = ""
var winTitle = ""
var mouseX, mouseY int

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(1792, 1008, "Hello World!", nil, nil)
	if err != nil {
		panic(err)
	}

	procAddr := glfw.GetVulkanGetInstanceProcAddress()
	if procAddr == nil {
		panic("GetInstanceProcAddress is nil")
	}
	vulkan.SetGetInstanceProcAddr(procAddr)

	if err = vulkan.Init(); err != nil {
		panic(err)
	}

	window.SetKeyCallback(app.KeyboardCallBack)

	var extensionCount uint32 = 0
	vulkan.EnumerateInstanceExtensionProperties("", &extensionCount, nil)
	fmt.Println(fmt.Sprint(extensionCount) + " extensions supported")

	fpsCounter := 0
	fpsTime := int(glfw.GetTime())
	for !window.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		mx, my := window.GetCursorPos()
		mouseX = int(mx)
		mouseY = int(my)

		fpsCounter++
		tim := int(glfw.GetTime())
		if tim != fpsTime {
			fpsTime = tim
			fpsTxt = "fps: " + strconv.Itoa(fpsCounter)
			fpsCounter = 0
		}

		title := fmt.Sprintf("%d, %d - %s", mouseX, mouseY, fpsTxt)
		if title != winTitle {
			winTitle = title
			window.SetTitle(winTitle)
		}
	}
}
