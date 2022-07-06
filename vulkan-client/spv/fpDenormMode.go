package spv

type FPDenormMode uint32

const (
	FPDenormModePreserve    FPDenormMode = 0
	FPDenormModeFlushToZero FPDenormMode = 1
	FPDenormModeMax         FPDenormMode = 0x7fffffff
)
