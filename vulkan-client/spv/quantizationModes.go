package spv

type QuantizationModes uint32

//goland:noinspection GoUnusedConst
const (
	QuantizationModesTRN        QuantizationModes = 0
	QuantizationModesTrnZERO    QuantizationModes = 1
	QuantizationModesRnd        QuantizationModes = 2
	QuantizationModesRndZERO    QuantizationModes = 3
	QuantizationModesRndINF     QuantizationModes = 4
	QuantizationModesRndMinINF  QuantizationModes = 5
	QuantizationModesRndConv    QuantizationModes = 6
	QuantizationModesRndConvODD QuantizationModes = 7
	QuantizationModesMax        QuantizationModes = 0x7fffffff
)
