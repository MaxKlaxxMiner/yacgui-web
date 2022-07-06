package spv

type FPFastMathModeShift uint32
type FPFastMathModeMask uint32

//goland:noinspection GoUnusedConst
const (
	FPFastMathModeNotNaNShift                 FPFastMathModeShift = 0
	FPFastMathModeNotInfShift                 FPFastMathModeShift = 1
	FPFastMathModeNSZShift                    FPFastMathModeShift = 2
	FPFastMathModeAllowRecipShift             FPFastMathModeShift = 3
	FPFastMathModeFastShift                   FPFastMathModeShift = 4
	FPFastMathModeAllowContractFastINTELShift FPFastMathModeShift = 16
	FPFastMathModeAllowReassocINTELShift      FPFastMathModeShift = 17
	FPFastMathModeMax                         FPFastMathModeShift = 0x7fffffff

	FPFastMathModeMaskNone                   FPFastMathModeMask = 0
	FPFastMathModeNotNaNMask                 FPFastMathModeMask = 0x00000001
	FPFastMathModeNotInfMask                 FPFastMathModeMask = 0x00000002
	FPFastMathModeNSZMask                    FPFastMathModeMask = 0x00000004
	FPFastMathModeAllowRecipMask             FPFastMathModeMask = 0x00000008
	FPFastMathModeFastMask                   FPFastMathModeMask = 0x00000010
	FPFastMathModeAllowContractFastINTELMask FPFastMathModeMask = 0x00010000
	FPFastMathModeAllowReassocINTELMask      FPFastMathModeMask = 0x00020000
)
