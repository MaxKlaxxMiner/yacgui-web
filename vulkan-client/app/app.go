package app

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
	"strconv"
)

type App struct {
	win            *glfw.Window
	physicalDevice vk.PhysicalDevice
	instance       vk.Instance
	config         Config
	debugMessenger vk.DebugReportCallback
	mouseX, mouseY int
}

type Config struct {
	EnableValidationLayers bool
	ValidationLayers       []string
}

func New(config Config) *App {
	return &App{config: config}
}

func (a *App) Run() (err error) {
	if err = a.initWindow(); err != nil {
		return err
	}

	if err = a.initVulkan(); err != nil {
		return err
	}

	a.win.SetKeyCallback(KeyboardCallBack)

	a.mainLoop()
	a.cleanup()

	return nil
}

func (a *App) mainLoop() {
	fpsCounter := 0
	fpsTime := int(glfw.GetTime())
	var fpsTxt, winTitle string
	for !a.win.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		mx, my := a.win.GetCursorPos()
		a.mouseX = int(mx)
		a.mouseY = int(my)

		fpsCounter++
		tim := int(glfw.GetTime())
		if tim != fpsTime {
			fpsTime = tim
			fpsTxt = "fps: " + strconv.Itoa(fpsCounter)
			fpsCounter = 0
		}

		title := fmt.Sprintf("%d, %d - %s", a.mouseX, a.mouseY, fpsTxt)
		if title != winTitle {
			winTitle = title
			a.win.SetTitle(winTitle)
		}
	}
}

func (a *App) initWindow() (err error) {
	if err = glfw.Init(); err != nil {
		return err
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	a.win, err = glfw.CreateWindow(1792, 1008, "Hello World!", nil, nil)

	return
}

func (a *App) cleanup() {
	a.cleanupVulkan()
	a.win.Destroy()
	glfw.Terminate()
}
