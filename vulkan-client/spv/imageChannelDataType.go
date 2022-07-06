package spv

type ImageChannelDataType uint32

const (
	ImageChannelDataTypeSnormInt8       ImageChannelDataType = 0
	ImageChannelDataTypeSnormInt16      ImageChannelDataType = 1
	ImageChannelDataTypeUnormInt8       ImageChannelDataType = 2
	ImageChannelDataTypeUnormInt16      ImageChannelDataType = 3
	ImageChannelDataTypeUnormShort565   ImageChannelDataType = 4
	ImageChannelDataTypeUnormShort555   ImageChannelDataType = 5
	ImageChannelDataTypeUnormInt101010  ImageChannelDataType = 6
	ImageChannelDataTypeSignedInt8      ImageChannelDataType = 7
	ImageChannelDataTypeSignedInt16     ImageChannelDataType = 8
	ImageChannelDataTypeSignedInt32     ImageChannelDataType = 9
	ImageChannelDataTypeUnsignedInt8    ImageChannelDataType = 10
	ImageChannelDataTypeUnsignedInt16   ImageChannelDataType = 11
	ImageChannelDataTypeUnsignedInt32   ImageChannelDataType = 12
	ImageChannelDataTypeHalfFloat       ImageChannelDataType = 13
	ImageChannelDataTypeFloat           ImageChannelDataType = 14
	ImageChannelDataTypeUnormInt24      ImageChannelDataType = 15
	ImageChannelDataTypeUnormInt101010b ImageChannelDataType = 16
	ImageChannelDataTypeMax             ImageChannelDataType = 0x7fffffff
)
