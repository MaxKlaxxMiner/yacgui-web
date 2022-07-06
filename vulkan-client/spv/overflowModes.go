package spv

type OverflowModes uint32

//goland:noinspection GoUnusedConst
const (
	OverflowModesWRAP    OverflowModes = 0
	OverflowModesSat     OverflowModes = 1
	OverflowModesSatZERO OverflowModes = 2
	OverflowModesSatSYM  OverflowModes = 3
	OverflowModesMax     OverflowModes = 0x7fffffff
)
