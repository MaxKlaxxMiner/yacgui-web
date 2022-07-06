package spv

type KernelEnqueueFlags uint32

const (
	KernelEnqueueFlagsNoWait        KernelEnqueueFlags = 0
	KernelEnqueueFlagsWaitKernel    KernelEnqueueFlags = 1
	KernelEnqueueFlagsWaitWorkGroup KernelEnqueueFlags = 2
	KernelEnqueueFlagsMax           KernelEnqueueFlags = 0x7fffffff
)
