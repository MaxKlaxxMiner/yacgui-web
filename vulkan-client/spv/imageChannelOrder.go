package spv

type ImageChannelOrder uint32

const (
	ImageChannelOrderR            ImageChannelOrder = 0
	ImageChannelOrderA            ImageChannelOrder = 1
	ImageChannelOrderRG           ImageChannelOrder = 2
	ImageChannelOrderRA           ImageChannelOrder = 3
	ImageChannelOrderRGB          ImageChannelOrder = 4
	ImageChannelOrderRGBA         ImageChannelOrder = 5
	ImageChannelOrderBGRA         ImageChannelOrder = 6
	ImageChannelOrderARGB         ImageChannelOrder = 7
	ImageChannelOrderIntensity    ImageChannelOrder = 8
	ImageChannelOrderLuminance    ImageChannelOrder = 9
	ImageChannelOrderRx           ImageChannelOrder = 10
	ImageChannelOrderRGx          ImageChannelOrder = 11
	ImageChannelOrderRGBx         ImageChannelOrder = 12
	ImageChannelOrderDepth        ImageChannelOrder = 13
	ImageChannelOrderDepthStencil ImageChannelOrder = 14
	ImageChannelOrdersRGB         ImageChannelOrder = 15
	ImageChannelOrdersRGBx        ImageChannelOrder = 16
	ImageChannelOrdersRGBA        ImageChannelOrder = 17
	ImageChannelOrdersBGRA        ImageChannelOrder = 18
	ImageChannelOrderABGR         ImageChannelOrder = 19
	ImageChannelOrderMax          ImageChannelOrder = 0x7fffffff
)
