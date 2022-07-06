package spv

type FPRoundingMode uint32

const (
	FPRoundingModeRTE FPRoundingMode = 0
	FPRoundingModeRTZ FPRoundingMode = 1
	FPRoundingModeRTP FPRoundingMode = 2
	FPRoundingModeRTN FPRoundingMode = 3
	FPRoundingModeMax FPRoundingMode = 0x7fffffff
)
