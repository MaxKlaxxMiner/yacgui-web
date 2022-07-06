package app

import (
	"fmt"
	vk "github.com/vulkan-go/vulkan"
	"strings"
	"unicode"
	"unsafe"
)

func (a *App) createLogicalDevice() error {
	indices := findQueueFamilies(a.physicalDevice, a.winSurface)

	uniqueQueueFamily := map[uint32]bool{
		*indices.graphicsFamily: true,
		*indices.presentFamily:  true,
	}

	var queueCreateInfos []vk.DeviceQueueCreateInfo
	for queueFamilyindex := range uniqueQueueFamily {
		queueCreateInfos = append(queueCreateInfos, vk.DeviceQueueCreateInfo{
			SType:            vk.StructureTypeDeviceQueueCreateInfo,
			QueueFamilyIndex: queueFamilyindex,
			QueueCount:       1,
			PQueuePriorities: []float32{1},
		})
	}

	//deviceFeatures := []vk.PhysicalDeviceFeatures{}

	deviceCreateInfo := vk.DeviceCreateInfo{
		SType:                   vk.StructureTypeDeviceCreateInfo,
		QueueCreateInfoCount:    uint32(len(queueCreateInfos)),
		PQueueCreateInfos:       queueCreateInfos,
		EnabledExtensionCount:   uint32(len(a.config.RequiredDeviceExtensions)),
		PpEnabledExtensionNames: a.config.RequiredDeviceExtensions,
	}

	if a.config.EnableValidationLayers {
		deviceCreateInfo.EnabledLayerCount = uint32(len(a.config.ValidationLayers))
		deviceCreateInfo.PpEnabledLayerNames = a.config.ValidationLayers
	}

	var device vk.Device
	if vk.CreateDevice(a.physicalDevice, &deviceCreateInfo, nil, &device) != vk.Success {
		return fmt.Errorf("could not create logical device")
	}

	a.logicalDevice = device

	return nil
}

func (a *App) isDeviceSuitable(device vk.PhysicalDevice) bool {
	if !checkDeviceExtensionsSupport(device, a.config.RequiredDeviceExtensions) {
		return false
	}

	swapChainSupport := querySwapChainSupport(device, a.winSurface)
	if len(swapChainSupport.surfaceFormats) == 0 || len(swapChainSupport.presentationModes) == 0 {
		return false
	}

	indices := findQueueFamilies(device, a.winSurface)
	if !indices.isComplete() {
		return false
	}

	return true
}

func checkDeviceExtensionsSupport(device vk.PhysicalDevice, requiredDeviceExtensions []string) bool {
	var count uint32
	vk.EnumerateDeviceExtensionProperties(device, "", &count, nil)
	extensionProperties := make([]vk.ExtensionProperties, count)
	vk.EnumerateDeviceExtensionProperties(device, "", &count, extensionProperties)

	supportedExtensions := make(map[string]bool, len(extensionProperties))
	for _, ep := range extensionProperties {
		ep.Deref()
		supportedExtensions[vk.ToString(ep.ExtensionName[:])] = true
		ep.Free()
	}

	for _, requiredExtension := range requiredDeviceExtensions {
		requiredExtension = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, requiredExtension)

		if !supportedExtensions[requiredExtension] {
			return false
		}
	}

	return true
}

type queueFamilyIndices struct {
	graphicsFamily *uint32
	presentFamily  *uint32
}

func (q *queueFamilyIndices) isComplete() bool {
	return q.graphicsFamily != nil && q.presentFamily != nil
}

func (a *App) pickPhysicalDevice() error {
	var deviceCount uint32
	vk.EnumeratePhysicalDevices(a.instance, &deviceCount, nil)
	if deviceCount == 0 {
		return fmt.Errorf("failed to find gpus with vulkan support")
	}

	physicalDevices := make([]vk.PhysicalDevice, deviceCount)
	vk.EnumeratePhysicalDevices(a.instance, &deviceCount, physicalDevices)

	for _, physicalDevice := range physicalDevices {
		if a.isDeviceSuitable(physicalDevice) {
			a.physicalDevice = physicalDevice
			break
		}
	}

	if unsafe.Pointer(a.physicalDevice) == vk.NullHandle {
		return fmt.Errorf("failed to find a suitable gpu")
	}

	return nil
}

func findQueueFamilies(device vk.PhysicalDevice, surface vk.Surface) queueFamilyIndices {
	var indices queueFamilyIndices

	var propCount uint32
	vk.GetPhysicalDeviceQueueFamilyProperties(device, &propCount, nil)

	properties := make([]vk.QueueFamilyProperties, propCount)
	vk.GetPhysicalDeviceQueueFamilyProperties(device, &propCount, properties)

	for i, property := range properties {
		property.Deref()
		queueFlags := property.QueueFlags
		property.Free()

		if (uint32(queueFlags) & uint32(vk.QueueGraphicsBit)) != 0 {
			tmp := uint32(i)
			indices.graphicsFamily = &tmp
		}

		var isSupported vk.Bool32
		vk.GetPhysicalDeviceSurfaceSupport(device, uint32(i), surface, &isSupported)
		if isSupported == vk.True {
			tmp := uint32(i)
			indices.presentFamily = &tmp
		}

		if indices.isComplete() {
			break
		}
	}

	return indices
}
