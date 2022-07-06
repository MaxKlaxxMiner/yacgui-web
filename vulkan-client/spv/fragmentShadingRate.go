package spv

type FragmentShadingRateShift uint32
type FragmentShadingRateMask uint32

//goland:noinspection GoUnusedConst
const (
	FragmentShadingRateVertical2PixelsShift   FragmentShadingRateShift = 0
	FragmentShadingRateVertical4PixelsShift   FragmentShadingRateShift = 1
	FragmentShadingRateHorizontal2PixelsShift FragmentShadingRateShift = 2
	FragmentShadingRateHorizontal4PixelsShift FragmentShadingRateShift = 3
	FragmentShadingRateMax                    FragmentShadingRateShift = 0x7fffffff

	FragmentShadingRateMaskNone              FragmentShadingRateMask = 0
	FragmentShadingRateVertical2PixelsMask   FragmentShadingRateMask = 0x00000001
	FragmentShadingRateVertical4PixelsMask   FragmentShadingRateMask = 0x00000002
	FragmentShadingRateHorizontal2PixelsMask FragmentShadingRateMask = 0x00000004
	FragmentShadingRateHorizontal4PixelsMask FragmentShadingRateMask = 0x00000008
)
