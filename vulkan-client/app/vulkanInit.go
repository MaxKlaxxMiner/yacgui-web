package app

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
	"unsafe"
)

func (a *App) cleanupVulkan() {
	for i := range a.swapChainImageViews {
		vk.DestroyImageView(a.logicalDevice, a.swapChainImageViews[i], nil)
	}
	vk.DestroySwapchain(a.logicalDevice, a.swapChain, nil)
	vk.DestroyDevice(a.logicalDevice, nil)
	if a.config.EnableValidationLayers {
		vk.DestroyDebugReportCallback(a.instance, a.debugMessenger, nil)
	}
	vk.DestroySurface(a.instance, a.winSurface, nil)
	vk.DestroyInstance(a.instance, nil)
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
		return
	}

	if a.config.EnableValidationLayers {
		if err = a.setupDebugMessenger(); err != nil {
			return
		}
	}

	if err = a.createWindowSurface(); err != nil {
		return
	}

	if err = a.pickPhysicalDevice(); err != nil {
		return
	}

	if err = a.createLogicalDevice(); err != nil {
		return
	}

	if err = a.createSwapChain(); err != nil {
		return
	}

	if err = a.createImageViews(); err != nil {
		return
	}

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

func (a *App) createWindowSurface() error {
	surfaceAddr, err := a.win.CreateWindowSurface(a.instance, nil)
	if err != nil {
		return err
	}

	a.winSurface = vk.SurfaceFromPointer(surfaceAddr)

	return nil
}
