package spv

type FunctionControlShift uint32
type FunctionControlMask uint32

const (
	FunctionControlInlineShift       FunctionControlShift = 0
	FunctionControlDontInlineShift   FunctionControlShift = 1
	FunctionControlPureShift         FunctionControlShift = 2
	FunctionControlConstShift        FunctionControlShift = 3
	FunctionControlOptNoneINTELShift FunctionControlShift = 16
	FunctionControlMax               FunctionControlShift = 0x7fffffff

	FunctionControlMaskNone         FunctionControlMask = 0
	FunctionControlInlineMask       FunctionControlMask = 0x00000001
	FunctionControlDontInlineMask   FunctionControlMask = 0x00000002
	FunctionControlPureMask         FunctionControlMask = 0x00000004
	FunctionControlConstMask        FunctionControlMask = 0x00000008
	FunctionControlOptNoneINTELMask FunctionControlMask = 0x00010000
)
