package spv

type LoopControlShift uint32
type LoopControlMask uint32

//goland:noinspection GoUnusedConst
const (
	LoopControlUnrollShift                    LoopControlShift = 0
	LoopControlDontUnrollShift                LoopControlShift = 1
	LoopControlDependencyInfiniteShift        LoopControlShift = 2
	LoopControlDependencyLengthShift          LoopControlShift = 3
	LoopControlMinIterationsShift             LoopControlShift = 4
	LoopControlMaxIterationsShift             LoopControlShift = 5
	LoopControlIterationMultipleShift         LoopControlShift = 6
	LoopControlPeelCountShift                 LoopControlShift = 7
	LoopControlPartialCountShift              LoopControlShift = 8
	LoopControlInitiationIntervalINTELShift   LoopControlShift = 16
	LoopControlMaxConcurrencyINTELShift       LoopControlShift = 17
	LoopControlDependencyArrayINTELShift      LoopControlShift = 18
	LoopControlPipelineEnableINTELShift       LoopControlShift = 19
	LoopControlLoopCoalesceINTELShift         LoopControlShift = 20
	LoopControlMaxInterleavingINTELShift      LoopControlShift = 21
	LoopControlSpeculatedIterationsINTELShift LoopControlShift = 22
	LoopControlNoFusionINTELShift             LoopControlShift = 23
	LoopControlMax                            LoopControlShift = 0x7fffffff

	LoopControlMaskNone                      LoopControlMask = 0
	LoopControlUnrollMask                    LoopControlMask = 0x00000001
	LoopControlDontUnrollMask                LoopControlMask = 0x00000002
	LoopControlDependencyInfiniteMask        LoopControlMask = 0x00000004
	LoopControlDependencyLengthMask          LoopControlMask = 0x00000008
	LoopControlMinIterationsMask             LoopControlMask = 0x00000010
	LoopControlMaxIterationsMask             LoopControlMask = 0x00000020
	LoopControlIterationMultipleMask         LoopControlMask = 0x00000040
	LoopControlPeelCountMask                 LoopControlMask = 0x00000080
	LoopControlPartialCountMask              LoopControlMask = 0x00000100
	LoopControlInitiationIntervalINTELMask   LoopControlMask = 0x00010000
	LoopControlMaxConcurrencyINTELMask       LoopControlMask = 0x00020000
	LoopControlDependencyArrayINTELMask      LoopControlMask = 0x00040000
	LoopControlPipelineEnableINTELMask       LoopControlMask = 0x00080000
	LoopControlLoopCoalesceINTELMask         LoopControlMask = 0x00100000
	LoopControlMaxInterleavingINTELMask      LoopControlMask = 0x00200000
	LoopControlSpeculatedIterationsINTELMask LoopControlMask = 0x00400000
	LoopControlNoFusionINTELMask             LoopControlMask = 0x00800000
)
