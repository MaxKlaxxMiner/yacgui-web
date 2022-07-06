package spv

type RayFlagsShift uint32
type RayFlagsMask uint32

const (
	RayFlagsOpaqueKHRShift                   RayFlagsShift = 0
	RayFlagsNoOpaqueKHRShift                 RayFlagsShift = 1
	RayFlagsTerminateOnFirstHitKHRShift      RayFlagsShift = 2
	RayFlagsSkipClosestHitShaderKHRShift     RayFlagsShift = 3
	RayFlagsCullBackFacingTrianglesKHRShift  RayFlagsShift = 4
	RayFlagsCullFrontFacingTrianglesKHRShift RayFlagsShift = 5
	RayFlagsCullOpaqueKHRShift               RayFlagsShift = 6
	RayFlagsCullNoOpaqueKHRShift             RayFlagsShift = 7
	RayFlagsSkipTrianglesKHRShift            RayFlagsShift = 8
	RayFlagsSkipAABBsKHRShift                RayFlagsShift = 9
	RayFlagsMax                              RayFlagsShift = 0x7fffffff

	RayFlagsMaskNone                        RayFlagsMask = 0
	RayFlagsOpaqueKHRMask                   RayFlagsMask = 0x00000001
	RayFlagsNoOpaqueKHRMask                 RayFlagsMask = 0x00000002
	RayFlagsTerminateOnFirstHitKHRMask      RayFlagsMask = 0x00000004
	RayFlagsSkipClosestHitShaderKHRMask     RayFlagsMask = 0x00000008
	RayFlagsCullBackFacingTrianglesKHRMask  RayFlagsMask = 0x00000010
	RayFlagsCullFrontFacingTrianglesKHRMask RayFlagsMask = 0x00000020
	RayFlagsCullOpaqueKHRMask               RayFlagsMask = 0x00000040
	RayFlagsCullNoOpaqueKHRMask             RayFlagsMask = 0x00000080
	RayFlagsSkipTrianglesKHRMask            RayFlagsMask = 0x00000100
	RayFlagsSkipAABBsKHRMask                RayFlagsMask = 0x00000200
)
