package spv

type MemoryAccessShift uint32
type MemoryAccessMask uint32

//goland:noinspection GoUnusedConst
const (
	MemoryAccessVolatileShift                MemoryAccessShift = 0
	MemoryAccessAlignedShift                 MemoryAccessShift = 1
	MemoryAccessNontemporalShift             MemoryAccessShift = 2
	MemoryAccessMakePointerAvailableShift    MemoryAccessShift = 3
	MemoryAccessMakePointerAvailableKHRShift MemoryAccessShift = 3
	MemoryAccessMakePointerVisibleShift      MemoryAccessShift = 4
	MemoryAccessMakePointerVisibleKHRShift   MemoryAccessShift = 4
	MemoryAccessNonPrivatePointerShift       MemoryAccessShift = 5
	MemoryAccessNonPrivatePointerKHRShift    MemoryAccessShift = 5
	MemoryAccessAliasScopeINTELMaskShift     MemoryAccessShift = 16
	MemoryAccessNoAliasINTELMaskShift        MemoryAccessShift = 17
	MemoryAccessMax                          MemoryAccessShift = 0x7fffffff

	MemoryAccessMaskNone                    MemoryAccessMask = 0
	MemoryAccessVolatileMask                MemoryAccessMask = 0x00000001
	MemoryAccessAlignedMask                 MemoryAccessMask = 0x00000002
	MemoryAccessNontemporalMask             MemoryAccessMask = 0x00000004
	MemoryAccessMakePointerAvailableMask    MemoryAccessMask = 0x00000008
	MemoryAccessMakePointerAvailableKHRMask MemoryAccessMask = 0x00000008
	MemoryAccessMakePointerVisibleMask      MemoryAccessMask = 0x00000010
	MemoryAccessMakePointerVisibleKHRMask   MemoryAccessMask = 0x00000010
	MemoryAccessNonPrivatePointerMask       MemoryAccessMask = 0x00000020
	MemoryAccessNonPrivatePointerKHRMask    MemoryAccessMask = 0x00000020
	MemoryAccessAliasScopeINTELMaskMask     MemoryAccessMask = 0x00010000
	MemoryAccessNoAliasINTELMaskMask        MemoryAccessMask = 0x00020000
)
