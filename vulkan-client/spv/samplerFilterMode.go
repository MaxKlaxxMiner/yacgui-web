package spv

type SamplerFilterMode uint32

const (
	SamplerFilterModeNearest SamplerFilterMode = 0
	SamplerFilterModeLinear  SamplerFilterMode = 1
	SamplerFilterModeMax     SamplerFilterMode = 0x7fffffff
)
