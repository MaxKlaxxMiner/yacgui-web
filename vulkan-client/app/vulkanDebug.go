package app

import (
	"fmt"
	vk "github.com/vulkan-go/vulkan"
	"log"
	"strings"
	"unicode"
	"unsafe"
)

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
