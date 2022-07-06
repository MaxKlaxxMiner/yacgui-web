package spv

type ImageFormat uint32

//goland:noinspection GoUnusedConst
const (
	ImageFormatUnknown      ImageFormat = 0
	ImageFormatRgba32f      ImageFormat = 1
	ImageFormatRgba16f      ImageFormat = 2
	ImageFormatR32f         ImageFormat = 3
	ImageFormatRgba8        ImageFormat = 4
	ImageFormatRgba8Snorm   ImageFormat = 5
	ImageFormatRg32f        ImageFormat = 6
	ImageFormatRg16f        ImageFormat = 7
	ImageFormatR11fG11fB10f ImageFormat = 8
	ImageFormatR16f         ImageFormat = 9
	ImageFormatRgba16       ImageFormat = 10
	ImageFormatRgb10A2      ImageFormat = 11
	ImageFormatRg16         ImageFormat = 12
	ImageFormatRg8          ImageFormat = 13
	ImageFormatR16          ImageFormat = 14
	ImageFormatR8           ImageFormat = 15
	ImageFormatRgba16Snorm  ImageFormat = 16
	ImageFormatRg16Snorm    ImageFormat = 17
	ImageFormatRg8Snorm     ImageFormat = 18
	ImageFormatR16Snorm     ImageFormat = 19
	ImageFormatR8Snorm      ImageFormat = 20
	ImageFormatRgba32i      ImageFormat = 21
	ImageFormatRgba16i      ImageFormat = 22
	ImageFormatRgba8i       ImageFormat = 23
	ImageFormatR32i         ImageFormat = 24
	ImageFormatRg32i        ImageFormat = 25
	ImageFormatRg16i        ImageFormat = 26
	ImageFormatRg8i         ImageFormat = 27
	ImageFormatR16i         ImageFormat = 28
	ImageFormatR8i          ImageFormat = 29
	ImageFormatRgba32ui     ImageFormat = 30
	ImageFormatRgba16ui     ImageFormat = 31
	ImageFormatRgba8ui      ImageFormat = 32
	ImageFormatR32ui        ImageFormat = 33
	ImageFormatRgb10a2ui    ImageFormat = 34
	ImageFormatRg32ui       ImageFormat = 35
	ImageFormatRg16ui       ImageFormat = 36
	ImageFormatRg8ui        ImageFormat = 37
	ImageFormatR16ui        ImageFormat = 38
	ImageFormatR8ui         ImageFormat = 39
	ImageFormatR64ui        ImageFormat = 40
	ImageFormatR64i         ImageFormat = 41
	ImageFormatMax          ImageFormat = 0x7fffffff
)
