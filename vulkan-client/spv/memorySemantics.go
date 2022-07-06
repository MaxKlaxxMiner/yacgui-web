package spv

type MemorySemanticsShift uint32
type MemorySemanticsMask uint32

const (
	MemorySemanticsAcquireShift                MemorySemanticsShift = 1
	MemorySemanticsReleaseShift                MemorySemanticsShift = 2
	MemorySemanticsAcquireReleaseShift         MemorySemanticsShift = 3
	MemorySemanticsSequentiallyConsistentShift MemorySemanticsShift = 4
	MemorySemanticsUniformMemoryShift          MemorySemanticsShift = 6
	MemorySemanticsSubgroupMemoryShift         MemorySemanticsShift = 7
	MemorySemanticsWorkgroupMemoryShift        MemorySemanticsShift = 8
	MemorySemanticsCrossWorkgroupMemoryShift   MemorySemanticsShift = 9
	MemorySemanticsAtomicCounterMemoryShift    MemorySemanticsShift = 10
	MemorySemanticsImageMemoryShift            MemorySemanticsShift = 11
	MemorySemanticsOutputMemoryShift           MemorySemanticsShift = 12
	MemorySemanticsOutputMemoryKHRShift        MemorySemanticsShift = 12
	MemorySemanticsMakeAvailableShift          MemorySemanticsShift = 13
	MemorySemanticsMakeAvailableKHRShift       MemorySemanticsShift = 13
	MemorySemanticsMakeVisibleShift            MemorySemanticsShift = 14
	MemorySemanticsMakeVisibleKHRShift         MemorySemanticsShift = 14
	MemorySemanticsVolatileShift               MemorySemanticsShift = 15
	MemorySemanticsMax                         MemorySemanticsShift = 0x7fffffff

	MemorySemanticsMaskNone                   MemorySemanticsMask = 0
	MemorySemanticsAcquireMask                MemorySemanticsMask = 0x00000002
	MemorySemanticsReleaseMask                MemorySemanticsMask = 0x00000004
	MemorySemanticsAcquireReleaseMask         MemorySemanticsMask = 0x00000008
	MemorySemanticsSequentiallyConsistentMask MemorySemanticsMask = 0x00000010
	MemorySemanticsUniformMemoryMask          MemorySemanticsMask = 0x00000040
	MemorySemanticsSubgroupMemoryMask         MemorySemanticsMask = 0x00000080
	MemorySemanticsWorkgroupMemoryMask        MemorySemanticsMask = 0x00000100
	MemorySemanticsCrossWorkgroupMemoryMask   MemorySemanticsMask = 0x00000200
	MemorySemanticsAtomicCounterMemoryMask    MemorySemanticsMask = 0x00000400
	MemorySemanticsImageMemoryMask            MemorySemanticsMask = 0x00000800
	MemorySemanticsOutputMemoryMask           MemorySemanticsMask = 0x00001000
	MemorySemanticsOutputMemoryKHRMask        MemorySemanticsMask = 0x00001000
	MemorySemanticsMakeAvailableMask          MemorySemanticsMask = 0x00002000
	MemorySemanticsMakeAvailableKHRMask       MemorySemanticsMask = 0x00002000
	MemorySemanticsMakeVisibleMask            MemorySemanticsMask = 0x00004000
	MemorySemanticsMakeVisibleKHRMask         MemorySemanticsMask = 0x00004000
	MemorySemanticsVolatileMask               MemorySemanticsMask = 0x00008000
)
