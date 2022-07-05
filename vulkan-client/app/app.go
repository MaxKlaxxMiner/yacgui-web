package app

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
	"log"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
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

func (a *App) initVulkan() (err error) {
	procAddr := glfw.GetVulkanGetInstanceProcAddress()
	if procAddr == nil {
		return fmt.Errorf("GetInstanceProcAddress is nil")
	}
	vk.SetGetInstanceProcAddr(procAddr)

	if err = vk.Init(); err != nil {
		return
	}

	if err = a.createInstance(); err != nil {
		return err
	}

	if a.config.EnableValidationLayers {
		if err = a.setupDebugMessenger(); err != nil {
			return
		}
	}

	return nil
}

func (a *App) cleanup() {
	if a.config.EnableValidationLayers {
		vk.DestroyDebugReportCallback(a.instance, a.debugMessenger, nil)
	}
	vk.DestroyInstance(a.instance, nil)
	a.win.Destroy()
	glfw.Terminate()
}

func defaultDebugCreateInfo() vk.DebugReportCallbackCreateInfo {
	return vk.DebugReportCallbackCreateInfo{
		SType: vk.StructureTypeDebugReportCallbackCreateInfo,
		Flags: vk.DebugReportFlags(vk.DebugReportErrorBit | vk.DebugReportWarningBit),
		PfnCallback: func(flags vk.DebugReportFlags, objectType vk.DebugReportObjectType, object uint64, location uint, messageCode int32, pLayerPrefix string, pMessage string, pUserData unsafe.Pointer) vk.Bool32 {
			switch {
			case flags&vk.DebugReportFlags(vk.DebugReportErrorBit) != 0:
				log.Printf("[ERROR %d] %s on layer %s", messageCode, pMessage, pLayerPrefix)
			case flags&vk.DebugReportFlags(vk.DebugReportWarningBit) != 0:
				log.Printf("[WARN %d] %s on layer %s", messageCode, pMessage, pLayerPrefix)
			default:
				log.Printf("[WARN] unknown debug message %d (layer %s)", messageCode, pLayerPrefix)
			}
			return vk.Bool32(vk.False)
		},
	}
}

func (a *App) setupDebugMessenger() error {
	dbgCreateInfo := defaultDebugCreateInfo()
	var dbg vk.DebugReportCallback
	if err := vk.Error(vk.CreateDebugReportCallback(a.instance, &dbgCreateInfo, nil, &dbg)); err != nil {
		return fmt.Errorf("vk.CreateDebugReportCallback failed with %s", err)
	}
	a.debugMessenger = dbg
	return nil
}

func (a *App) createInstance() (err error) {
	requiredExtensions := a.win.GetRequiredInstanceExtensions()
	if a.config.EnableValidationLayers {
		requiredExtensions = append(requiredExtensions, "VK_EXT_debug_report\x00")
	}
	if err = checkExtensionSupport(requiredExtensions); err != nil {
		return err
	}

	applicationInfo := vk.ApplicationInfo{
		SType:              vk.StructureTypeApplicationInfo,
		PApplicationName:   "Hello Triangle",
		ApplicationVersion: vk.MakeVersion(1, 0, 0),
		PEngineName:        "No Engine",
		EngineVersion:      vk.MakeVersion(1, 0, 0),
		ApiVersion:         vk.MakeVersion(1, 0, 0),
	}

	dbgCreateInfo := defaultDebugCreateInfo()
	instanceCreateInfo := vk.InstanceCreateInfo{
		SType:                   vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo:        &applicationInfo,
		EnabledExtensionCount:   uint32(len(requiredExtensions)),
		PpEnabledExtensionNames: requiredExtensions,
		PNext:                   unsafe.Pointer(dbgCreateInfo.Ref()),
	}

	if a.config.EnableValidationLayers {
		if err = checkValidationLayerSupport(a.config.ValidationLayers); err != nil {
			return err
		}
		instanceCreateInfo.PpEnabledLayerNames = a.config.ValidationLayers
		instanceCreateInfo.EnabledLayerCount = uint32(len(a.config.ValidationLayers))
	}

	var instance vk.Instance
	res := vk.CreateInstance(&instanceCreateInfo, nil, &instance)
	if res != vk.Success {
		return fmt.Errorf("failed to create instance")
	}
	a.instance = instance
	return nil
}

func checkExtensionSupport(requiredExtensions []string) error {
	var extensionCount uint32
	vk.EnumerateInstanceExtensionProperties("", &extensionCount, nil)
	extensionProperties := make([]vk.ExtensionProperties, extensionCount)
	vk.EnumerateInstanceExtensionProperties("", &extensionCount, extensionProperties)

	supportedExtensions := make(map[string]bool)
	for _, extensionProperty := range extensionProperties {
		extensionProperty.Deref()
		supportedExtensions[vk.ToString(extensionProperty.ExtensionName[:])] = true
		extensionProperty.Free()
	}

	for _, requiredExtension := range requiredExtensions {
		requiredExtension = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, requiredExtension)
		if !supportedExtensions[requiredExtension] {
			return fmt.Errorf(requiredExtension + " - is not a supported extension")
		}
	}

	return nil
}

func checkValidationLayerSupport(requiredLayers []string) error {
	var layerCount uint32
	vk.EnumerateInstanceLayerProperties(&layerCount, nil)
	layerProperties := make([]vk.LayerProperties, layerCount)
	vk.EnumerateInstanceLayerProperties(&layerCount, layerProperties)

	supportedLayers := make(map[string]bool)
	for _, layerProperty := range layerProperties {
		layerProperty.Deref()
		supportedLayers[vk.ToString(layerProperty.LayerName[:])] = true
		layerProperty.Free()
	}

	for _, requiredLayer := range requiredLayers {
		requiredLayer = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, requiredLayer)
		if !supportedLayers[requiredLayer] {
			return fmt.Errorf(requiredLayer + " - is not a supported layer")
		}
	}

	return nil
}
