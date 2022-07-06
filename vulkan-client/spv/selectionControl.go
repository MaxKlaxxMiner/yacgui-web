package spv

type SelectionControlShift uint32
type SelectionControlMask uint32

const (
	SelectionControlFlattenShift     SelectionControlShift = 0
	SelectionControlDontFlattenShift SelectionControlShift = 1
	SelectionControlMax              SelectionControlShift = 0x7fffffff

	SelectionControlMaskNone        SelectionControlMask = 0
	SelectionControlFlattenMask     SelectionControlMask = 0x00000001
	SelectionControlDontFlattenMask SelectionControlMask = 0x00000002
)
