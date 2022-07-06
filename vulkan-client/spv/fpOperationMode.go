package spv

type FPOperationMode uint32

const (
	FPOperationModeIEEE FPOperationMode = 0
	FPOperationModeALT  FPOperationMode = 1
	FPOperationModeMax  FPOperationMode = 0x7fffffff
)
