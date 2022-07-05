package app

import (
	"fmt"
	vk "github.com/vulkan-go/vulkan"
	"unsafe"
)

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
