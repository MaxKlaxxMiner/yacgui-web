package app

import (
	"fmt"
	vk "github.com/vulkan-go/vulkan"
	"unsafe"
)

func (a *App) createLogicalDevice() error {
	indices := findQueueFamilies(a.physicalDevice)

	queueCreateInfo := []vk.DeviceQueueCreateInfo{{
		SType:            vk.StructureTypeDeviceQueueCreateInfo,
		PNext:            nil,
		Flags:            0,
		QueueFamilyIndex: *indices.graphicsFamily,
		QueueCount:       1,
		PQueuePriorities: []float32{1},
	}}

	//deviceFeatures := []vk.PhysicalDeviceFeatures{}

	deviceCreateInfo := vk.DeviceCreateInfo{
		SType:                   vk.StructureTypeDeviceCreateInfo,
		PNext:                   nil,
		Flags:                   0,
		QueueCreateInfoCount:    1,
		PQueueCreateInfos:       queueCreateInfo,
		EnabledLayerCount:       0,
		PpEnabledLayerNames:     nil,
		EnabledExtensionCount:   0,
		PpEnabledExtensionNames: nil,
		PEnabledFeatures:        nil,
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

	var deviceQueue vk.Queue
	vk.GetDeviceQueue(device, *indices.graphicsFamily, 0, &deviceQueue)

	return nil
}

func isDeviceSuitable(device vk.PhysicalDevice) bool {
	var deviceProperties vk.PhysicalDeviceProperties
	var deviceFeatures vk.PhysicalDeviceFeatures
	vk.GetPhysicalDeviceProperties(device, &deviceProperties)
	vk.GetPhysicalDeviceFeatures(device, &deviceFeatures)

	indices := findQueueFamilies(device)
	if indices.graphicsFamily == nil {
		return false
	}

	return true
}

type queueFamilyIndices struct {
	graphicsFamily *uint32
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
		if isDeviceSuitable(physicalDevice) {
			a.physicalDevice = physicalDevice
			break
		}
	}

	if unsafe.Pointer(a.physicalDevice) == vk.NullHandle {
		return fmt.Errorf("failed to find a suitable gpu")
	}

	return nil
}

func findQueueFamilies(device vk.PhysicalDevice) queueFamilyIndices {
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
			break
		}
	}

	return indices
}
