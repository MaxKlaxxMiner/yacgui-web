package spv

type MemoryModel uint32

//goland:noinspection GoUnusedConst
const (
	MemoryModelSimple    MemoryModel = 0
	MemoryModelGLSL450   MemoryModel = 1
	MemoryModelOpenCL    MemoryModel = 2
	MemoryModelVulkan    MemoryModel = 3
	MemoryModelVulkanKHR MemoryModel = 3
	MemoryModelMax       MemoryModel = 0x7fffffff
)
