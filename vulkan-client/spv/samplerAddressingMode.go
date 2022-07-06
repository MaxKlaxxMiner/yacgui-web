package spv

type SamplerAddressingMode uint32

const (
	SamplerAddressingModeNone           SamplerAddressingMode = 0
	SamplerAddressingModeClampToEdge    SamplerAddressingMode = 1
	SamplerAddressingModeClamp          SamplerAddressingMode = 2
	SamplerAddressingModeRepeat         SamplerAddressingMode = 3
	SamplerAddressingModeRepeatMirrored SamplerAddressingMode = 4
	SamplerAddressingModeMax            SamplerAddressingMode = 0x7fffffff
)
