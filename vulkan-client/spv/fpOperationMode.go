package spv

type FPOperationMode uint32

//goland:noinspection GoUnusedConst
const (
	FPOperationModeIEEE FPOperationMode = 0
	FPOperationModeALT  FPOperationMode = 1
	FPOperationModeMax  FPOperationMode = 0x7fffffff
)
