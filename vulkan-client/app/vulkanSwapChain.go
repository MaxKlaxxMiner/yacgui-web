package app

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

func (a *App) createSwapChain() error {
	swapChainSupport := querySwapChainSupport(a.physicalDevice, a.winSurface)
	swapChainSupport.capabilities.Deref()
	swapChainSupport.capabilities.Free()

	surfaceFormat := chooseSwapSurfaceFormat(swapChainSupport.surfaceFormats...)
	presentationMode := chooseSwapPresentMode(swapChainSupport.presentationModes...)
	swapExtent := chooseSwapExtent(swapChainSupport.capabilities, a.win)
	imageCount := swapChainSupport.capabilities.MinImageCount + 1

	if swapChainSupport.capabilities.MaxImageCount > 0 && imageCount > swapChainSupport.capabilities.MaxImageCount {
		imageCount = swapChainSupport.capabilities.MaxImageCount
	}

	createInfo := vk.SwapchainCreateInfo{
		SType:            vk.StructureTypeSwapchainCreateInfo,
		Surface:          a.winSurface,
		MinImageCount:    imageCount,
		ImageFormat:      surfaceFormat.Format,
		ImageColorSpace:  surfaceFormat.ColorSpace,
		ImageExtent:      swapExtent,
		ImageArrayLayers: 1,
		ImageUsage:       vk.ImageUsageFlags(vk.ImageUsageColorAttachmentBit),
		PreTransform:     swapChainSupport.capabilities.CurrentTransform,
		CompositeAlpha:   vk.CompositeAlphaOpaqueBit,
		PresentMode:      presentationMode,
		Clipped:          vk.True,
		OldSwapchain:     vk.NullSwapchain,
	}

	indices := findQueueFamilies(a.physicalDevice, a.winSurface)
	queueFamilies := []uint32{*indices.presentFamily, *indices.graphicsFamily}

	if *indices.graphicsFamily != *indices.presentFamily {
		createInfo.ImageSharingMode = vk.SharingModeConcurrent
		createInfo.QueueFamilyIndexCount = uint32(len(queueFamilies))
		createInfo.PQueueFamilyIndices = queueFamilies
	} else {
		createInfo.ImageSharingMode = vk.SharingModeExclusive
	}

	var swapChain vk.Swapchain
	err := vk.Error(vk.CreateSwapchain(a.logicalDevice, &createInfo, nil, &swapChain))
	if err != nil {
		return err
	}

	a.swapChain = swapChain

	var imagesCount uint32
	vk.GetSwapchainImages(a.logicalDevice, a.swapChain, &imagesCount, nil)
	a.swapChainImages = make([]vk.Image, imageCount)
	vk.GetSwapchainImages(a.logicalDevice, a.swapChain, &imagesCount, a.swapChainImages)

	a.swapChainExtent = swapExtent
	a.swapChainImageFormat = surfaceFormat.Format

	return nil
}

type swapChainSupportDetails struct {
	capabilities      vk.SurfaceCapabilities
	surfaceFormats    []vk.SurfaceFormat
	presentationModes []vk.PresentMode
}

func querySwapChainSupport(device vk.PhysicalDevice, surface vk.Surface) swapChainSupportDetails {
	var details swapChainSupportDetails

	vk.GetPhysicalDeviceSurfaceCapabilities(device, surface, &details.capabilities)

	var formatCount uint32
	vk.GetPhysicalDeviceSurfaceFormats(device, surface, &formatCount, nil)
	if formatCount != 0 {
		surfaceFormats := make([]vk.SurfaceFormat, formatCount)
		vk.GetPhysicalDeviceSurfaceFormats(device, surface, &formatCount, surfaceFormats)
		details.surfaceFormats = surfaceFormats
	}

	var modeCount uint32
	vk.GetPhysicalDeviceSurfacePresentModes(device, surface, &modeCount, nil)
	if modeCount != 0 {
		presentationModes := make([]vk.PresentMode, modeCount)
		vk.GetPhysicalDeviceSurfacePresentModes(device, surface, &modeCount, presentationModes)
		details.presentationModes = presentationModes
	}

	return details
}

func chooseSwapSurfaceFormat(surfaceFormats ...vk.SurfaceFormat) vk.SurfaceFormat {
	if len(surfaceFormats) < 1 {
		return vk.SurfaceFormat{}
	}

	for _, surfaceFormat := range surfaceFormats {
		surfaceFormat.Deref()
		surfaceFormat.Free()

		if surfaceFormat.Format == vk.FormatB8g8r8a8Srgb && surfaceFormat.ColorSpace == vk.ColorspaceSrgbNonlinear {
			return surfaceFormat
		}
	}

	return surfaceFormats[0]
}

func chooseSwapPresentMode(presentModes ...vk.PresentMode) vk.PresentMode {
	if len(presentModes) < 1 {
		return 0
	}

	for _, presentMode := range presentModes {
		if presentMode == vk.PresentModeMailbox {
			return presentMode
		}
	}

	return vk.PresentModeFifo
}

func chooseSwapExtent(surfaceCapabilities vk.SurfaceCapabilities, win *glfw.Window) vk.Extent2D {
	surfaceCapabilities.Deref()
	surfaceCapabilities.Free()
	surfaceCapabilities.CurrentExtent.Deref()
	surfaceCapabilities.CurrentExtent.Free()
	surfaceCapabilities.MaxImageExtent.Deref()
	surfaceCapabilities.MaxImageExtent.Free()
	surfaceCapabilities.MinImageExtent.Deref()
	surfaceCapabilities.MinImageExtent.Free()

	//if surfaceCapabilities.CurrentExtent.Width != vk.MaxUint32 {
	//	return surfaceCapabilities.CurrentExtent
	//}

	w, h := win.GetFramebufferSize()

	actualExtent := vk.Extent2D{
		Width:  uint32(w),
		Height: uint32(h),
	}

	if actualExtent.Width > surfaceCapabilities.MaxImageExtent.Width {
		actualExtent.Width = surfaceCapabilities.MaxImageExtent.Width
	}
	if actualExtent.Width < surfaceCapabilities.MinImageExtent.Width {
		actualExtent.Width = surfaceCapabilities.MinImageExtent.Width
	}

	if actualExtent.Height > surfaceCapabilities.MaxImageExtent.Height {
		actualExtent.Height = surfaceCapabilities.MaxImageExtent.Height
	}
	if actualExtent.Height < surfaceCapabilities.MinImageExtent.Height {
		actualExtent.Height = surfaceCapabilities.MinImageExtent.Height
	}

	return actualExtent
}
