package spv

type ExecutionMode uint32

const (
	ExecutionModeInvocations                      ExecutionMode = 0
	ExecutionModeSpacingEqual                     ExecutionMode = 1
	ExecutionModeSpacingFractionalEven            ExecutionMode = 2
	ExecutionModeSpacingFractionalOdd             ExecutionMode = 3
	ExecutionModeVertexOrderCw                    ExecutionMode = 4
	ExecutionModeVertexOrderCcw                   ExecutionMode = 5
	ExecutionModePixelCenterInteger               ExecutionMode = 6
	ExecutionModeOriginUpperLeft                  ExecutionMode = 7
	ExecutionModeOriginLowerLeft                  ExecutionMode = 8
	ExecutionModeEarlyFragmentTests               ExecutionMode = 9
	ExecutionModePointMode                        ExecutionMode = 10
	ExecutionModeXfb                              ExecutionMode = 11
	ExecutionModeDepthReplacing                   ExecutionMode = 12
	ExecutionModeDepthGreater                     ExecutionMode = 14
	ExecutionModeDepthLess                        ExecutionMode = 15
	ExecutionModeDepthUnchanged                   ExecutionMode = 16
	ExecutionModeLocalSize                        ExecutionMode = 17
	ExecutionModeLocalSizeHint                    ExecutionMode = 18
	ExecutionModeInputPoints                      ExecutionMode = 19
	ExecutionModeInputLines                       ExecutionMode = 20
	ExecutionModeInputLinesAdjacency              ExecutionMode = 21
	ExecutionModeTriangles                        ExecutionMode = 22
	ExecutionModeInputTrianglesAdjacency          ExecutionMode = 23
	ExecutionModeQuads                            ExecutionMode = 24
	ExecutionModeIsolines                         ExecutionMode = 25
	ExecutionModeOutputVertices                   ExecutionMode = 26
	ExecutionModeOutputPoints                     ExecutionMode = 27
	ExecutionModeOutputLineStrip                  ExecutionMode = 28
	ExecutionModeOutputTriangleStrip              ExecutionMode = 29
	ExecutionModeVecTypeHint                      ExecutionMode = 30
	ExecutionModeContractionOff                   ExecutionMode = 31
	ExecutionModeInitializer                      ExecutionMode = 33
	ExecutionModeFinalizer                        ExecutionMode = 34
	ExecutionModeSubgroupSize                     ExecutionMode = 35
	ExecutionModeSubgroupsPerWorkgroup            ExecutionMode = 36
	ExecutionModeSubgroupsPerWorkgroupId          ExecutionMode = 37
	ExecutionModeLocalSizeId                      ExecutionMode = 38
	ExecutionModeLocalSizeHintId                  ExecutionMode = 39
	ExecutionModeSubgroupUniformControlFlowKHR    ExecutionMode = 4421
	ExecutionModePostDepthCoverage                ExecutionMode = 4446
	ExecutionModeDenormPreserve                   ExecutionMode = 4459
	ExecutionModeDenormFlushToZero                ExecutionMode = 4460
	ExecutionModeSignedZeroInfNanPreserve         ExecutionMode = 4461
	ExecutionModeRoundingModeRTE                  ExecutionMode = 4462
	ExecutionModeRoundingModeRTZ                  ExecutionMode = 4463
	ExecutionModeEarlyAndLateFragmentTestsAMD     ExecutionMode = 5017
	ExecutionModeStencilRefReplacingEXT           ExecutionMode = 5027
	ExecutionModeStencilRefUnchangedFrontAMD      ExecutionMode = 5079
	ExecutionModeStencilRefGreaterFrontAMD        ExecutionMode = 5080
	ExecutionModeStencilRefLessFrontAMD           ExecutionMode = 5081
	ExecutionModeStencilRefUnchangedBackAMD       ExecutionMode = 5082
	ExecutionModeStencilRefGreaterBackAMD         ExecutionMode = 5083
	ExecutionModeStencilRefLessBackAMD            ExecutionMode = 5084
	ExecutionModeOutputLinesNV                    ExecutionMode = 5269
	ExecutionModeOutputPrimitivesNV               ExecutionMode = 5270
	ExecutionModeDerivativeGroupQuadsNV           ExecutionMode = 5289
	ExecutionModeDerivativeGroupLinearNV          ExecutionMode = 5290
	ExecutionModeOutputTrianglesNV                ExecutionMode = 5298
	ExecutionModePixelInterlockOrderedEXT         ExecutionMode = 5366
	ExecutionModePixelInterlockUnorderedEXT       ExecutionMode = 5367
	ExecutionModeSampleInterlockOrderedEXT        ExecutionMode = 5368
	ExecutionModeSampleInterlockUnorderedEXT      ExecutionMode = 5369
	ExecutionModeShadingRateInterlockOrderedEXT   ExecutionMode = 5370
	ExecutionModeShadingRateInterlockUnorderedEXT ExecutionMode = 5371
	ExecutionModeSharedLocalMemorySizeINTEL       ExecutionMode = 5618
	ExecutionModeRoundingModeRTPINTEL             ExecutionMode = 5620
	ExecutionModeRoundingModeRTNINTEL             ExecutionMode = 5621
	ExecutionModeFloatingPointModeALTINTEL        ExecutionMode = 5622
	ExecutionModeFloatingPointModeIEEEINTEL       ExecutionMode = 5623
	ExecutionModeMaxWorkgroupSizeINTEL            ExecutionMode = 5893
	ExecutionModeMaxWorkDimINTEL                  ExecutionMode = 5894
	ExecutionModeNoGlobalOffsetINTEL              ExecutionMode = 5895
	ExecutionModeNumSIMDWorkitemsINTEL            ExecutionMode = 5896
	ExecutionModeSchedulerTargetFmaxMhzINTEL      ExecutionMode = 5903
	ExecutionModeNamedBarrierCountINTEL           ExecutionMode = 6417
	ExecutionModeMax                              ExecutionMode = 0x7fffffff
)