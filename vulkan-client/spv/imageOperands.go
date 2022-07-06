package spv

type ImageOperandsShift uint32
type ImageOperandsMask uint32

const (
	ImageOperandsBiasShift                  ImageOperandsShift = 0
	ImageOperandsLodShift                   ImageOperandsShift = 1
	ImageOperandsGradShift                  ImageOperandsShift = 2
	ImageOperandsConstOffsetShift           ImageOperandsShift = 3
	ImageOperandsOffsetShift                ImageOperandsShift = 4
	ImageOperandsConstOffsetsShift          ImageOperandsShift = 5
	ImageOperandsSampleShift                ImageOperandsShift = 6
	ImageOperandsMinLodShift                ImageOperandsShift = 7
	ImageOperandsMakeTexelAvailableShift    ImageOperandsShift = 8
	ImageOperandsMakeTexelAvailableKHRShift ImageOperandsShift = 8
	ImageOperandsMakeTexelVisibleShift      ImageOperandsShift = 9
	ImageOperandsMakeTexelVisibleKHRShift   ImageOperandsShift = 9
	ImageOperandsNonPrivateTexelShift       ImageOperandsShift = 10
	ImageOperandsNonPrivateTexelKHRShift    ImageOperandsShift = 10
	ImageOperandsVolatileTexelShift         ImageOperandsShift = 11
	ImageOperandsVolatileTexelKHRShift      ImageOperandsShift = 11
	ImageOperandsSignExtendShift            ImageOperandsShift = 12
	ImageOperandsZeroExtendShift            ImageOperandsShift = 13
	ImageOperandsNontemporalShift           ImageOperandsShift = 14
	ImageOperandsOffsetsShift               ImageOperandsShift = 16
	ImageOperandsMax                        ImageOperandsShift = 0x7fffffff

	ImageOperandsMaskNone                  ImageOperandsMask = 0
	ImageOperandsBiasMask                  ImageOperandsMask = 0x00000001
	ImageOperandsLodMask                   ImageOperandsMask = 0x00000002
	ImageOperandsGradMask                  ImageOperandsMask = 0x00000004
	ImageOperandsConstOffsetMask           ImageOperandsMask = 0x00000008
	ImageOperandsOffsetMask                ImageOperandsMask = 0x00000010
	ImageOperandsConstOffsetsMask          ImageOperandsMask = 0x00000020
	ImageOperandsSampleMask                ImageOperandsMask = 0x00000040
	ImageOperandsMinLodMask                ImageOperandsMask = 0x00000080
	ImageOperandsMakeTexelAvailableMask    ImageOperandsMask = 0x00000100
	ImageOperandsMakeTexelAvailableKHRMask ImageOperandsMask = 0x00000100
	ImageOperandsMakeTexelVisibleMask      ImageOperandsMask = 0x00000200
	ImageOperandsMakeTexelVisibleKHRMask   ImageOperandsMask = 0x00000200
	ImageOperandsNonPrivateTexelMask       ImageOperandsMask = 0x00000400
	ImageOperandsNonPrivateTexelKHRMask    ImageOperandsMask = 0x00000400
	ImageOperandsVolatileTexelMask         ImageOperandsMask = 0x00000800
	ImageOperandsVolatileTexelKHRMask      ImageOperandsMask = 0x00000800
	ImageOperandsSignExtendMask            ImageOperandsMask = 0x00001000
	ImageOperandsZeroExtendMask            ImageOperandsMask = 0x00002000
	ImageOperandsNontemporalMask           ImageOperandsMask = 0x00004000
	ImageOperandsOffsetsMask               ImageOperandsMask = 0x00010000
)
