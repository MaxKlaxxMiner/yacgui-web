package spv

type KernelProfilingInfoShift uint32
type KernelProfilingInfoMask uint32

//goland:noinspection GoUnusedConst
const (
	KernelProfilingInfoCmdExecTimeShift KernelProfilingInfoShift = 0
	KernelProfilingInfoMax              KernelProfilingInfoShift = 0x7fffffff

	KernelProfilingInfoMaskNone        KernelProfilingInfoMask = 0
	KernelProfilingInfoCmdExecTimeMask KernelProfilingInfoMask = 0x00000001
)
