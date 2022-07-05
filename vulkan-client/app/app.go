package app

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
	"strconv"
)

type App struct {
	win            *glfw.Window
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

func (a *App) initWindow() (err error) {
	if err = glfw.Init(); err != nil {
		return err
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	a.win, err = glfw.CreateWindow(1792, 1008, "Hello World!", nil, nil)

	return
}

func (a *App) initVulkan() (err error) {
	procAddr := glfw.GetVulkanGetInstanceProcAddress()
	if procAddr == nil {
		return fmt.Errorf("GetInstanceProcAddress is nil")
	}

	vk.SetGetInstanceProcAddr(procAddr)

	if err = vk.Init(); err != nil {
		return
	}

	return a.createInstance()
}

func (a *App) createInstance() (err error) {
	glfwExtensions := a.win.GetRequiredInstanceExtensions()

	var extensionCount uint32
	vk.EnumerateInstanceExtensionProperties("", &extensionCount, nil)
	extensionProperties := make([]vk.ExtensionProperties, extensionCount)
	vk.EnumerateInstanceExtensionProperties("", &extensionCount, extensionProperties)

	supportedExtensions := make(map[string]bool)
	for _, extensionProperty := range extensionProperties {
		extensionProperty.Deref()
		supportedExtensions[vk.ToString(extensionProperty.ExtensionName[:])] = true
	}

	for _, glfwExtension := range glfwExtensions {
		if !supportedExtensions[glfwExtension] {
			return fmt.Errorf("glfwExtension - " + glfwExtension + " - is not supported by vulkan")
		}
	}

	applicationInfo := vk.ApplicationInfo{
		SType:              vk.StructureTypeApplicationInfo,
		PApplicationName:   "Hello Triangle",
		ApplicationVersion: vk.MakeVersion(1, 0, 0),
		PEngineName:        "No Engine",
		EngineVersion:      vk.MakeVersion(1, 0, 0),
		ApiVersion:         vk.MakeVersion(1, 0, 0),
	}

	instanceCreateInfo := vk.InstanceCreateInfo{
		SType:                   vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo:        &applicationInfo,
		EnabledExtensionCount:   uint32(len(glfwExtensions)),
		PpEnabledExtensionNames: glfwExtensions,
	}

	var instance vk.Instance
	res := vk.CreateInstance(&instanceCreateInfo, nil, &instance)
	if res != vk.Success {
		return fmt.Errorf("failed to create instance")
	}
	a.instance = instance
	return nil
}

func (a *App) cleanup() {
	vk.DestroyInstance(a.instance, nil)
	a.win.Destroy()
	glfw.Terminate()
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
