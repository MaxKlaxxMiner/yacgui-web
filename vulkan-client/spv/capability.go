package spv

type Capability uint32

const (
	CapabilityMatrix                                       Capability = 0
	CapabilityShader                                       Capability = 1
	CapabilityGeometry                                     Capability = 2
	CapabilityTessellation                                 Capability = 3
	CapabilityAddresses                                    Capability = 4
	CapabilityLinkage                                      Capability = 5
	CapabilityKernel                                       Capability = 6
	CapabilityVector16                                     Capability = 7
	CapabilityFloat16Buffer                                Capability = 8
	CapabilityFloat16                                      Capability = 9
	CapabilityFloat64                                      Capability = 10
	CapabilityInt64                                        Capability = 11
	CapabilityInt64Atomics                                 Capability = 12
	CapabilityImageBasic                                   Capability = 13
	CapabilityImageReadWrite                               Capability = 14
	CapabilityImageMipmap                                  Capability = 15
	CapabilityPipes                                        Capability = 17
	CapabilityGroups                                       Capability = 18
	CapabilityDeviceEnqueue                                Capability = 19
	CapabilityLiteralSampler                               Capability = 20
	CapabilityAtomicStorage                                Capability = 21
	CapabilityInt16                                        Capability = 22
	CapabilityTessellationPointSize                        Capability = 23
	CapabilityGeometryPointSize                            Capability = 24
	CapabilityImageGatherExtended                          Capability = 25
	CapabilityStorageImageMultisample                      Capability = 27
	CapabilityUniformBufferArrayDynamicIndexing            Capability = 28
	CapabilitySampledImageArrayDynamicIndexing             Capability = 29
	CapabilityStorageBufferArrayDynamicIndexing            Capability = 30
	CapabilityStorageImageArrayDynamicIndexing             Capability = 31
	CapabilityClipDistance                                 Capability = 32
	CapabilityCullDistance                                 Capability = 33
	CapabilityImageCubeArray                               Capability = 34
	CapabilitySampleRateShading                            Capability = 35
	CapabilityImageRect                                    Capability = 36
	CapabilitySampledRect                                  Capability = 37
	CapabilityGenericPointer                               Capability = 38
	CapabilityInt8                                         Capability = 39
	CapabilityInputAttachment                              Capability = 40
	CapabilitySparseResidency                              Capability = 41
	CapabilityMinLod                                       Capability = 42
	CapabilitySampled1D                                    Capability = 43
	CapabilityImage1D                                      Capability = 44
	CapabilitySampledCubeArray                             Capability = 45
	CapabilitySampledBuffer                                Capability = 46
	CapabilityImageBuffer                                  Capability = 47
	CapabilityImageMSArray                                 Capability = 48
	CapabilityStorageImageExtendedFormats                  Capability = 49
	CapabilityImageQuery                                   Capability = 50
	CapabilityDerivativeControl                            Capability = 51
	CapabilityInterpolationFunction                        Capability = 52
	CapabilityTransformFeedback                            Capability = 53
	CapabilityGeometryStreams                              Capability = 54
	CapabilityStorageImageReadWithoutFormat                Capability = 55
	CapabilityStorageImageWriteWithoutFormat               Capability = 56
	CapabilityMultiViewport                                Capability = 57
	CapabilitySubgroupDispatch                             Capability = 58
	CapabilityNamedBarrier                                 Capability = 59
	CapabilityPipeStorage                                  Capability = 60
	CapabilityGroupNonUniform                              Capability = 61
	CapabilityGroupNonUniformVote                          Capability = 62
	CapabilityGroupNonUniformArithmetic                    Capability = 63
	CapabilityGroupNonUniformBallot                        Capability = 64
	CapabilityGroupNonUniformShuffle                       Capability = 65
	CapabilityGroupNonUniformShuffleRelative               Capability = 66
	CapabilityGroupNonUniformClustered                     Capability = 67
	CapabilityGroupNonUniformQuad                          Capability = 68
	CapabilityShaderLayer                                  Capability = 69
	CapabilityShaderViewportIndex                          Capability = 70
	CapabilityUniformDecoration                            Capability = 71
	CapabilityFragmentShadingRateKHR                       Capability = 4422
	CapabilitySubgroupBallotKHR                            Capability = 4423
	CapabilityDrawParameters                               Capability = 4427
	CapabilityWorkgroupMemoryExplicitLayoutKHR             Capability = 4428
	CapabilityWorkgroupMemoryExplicitLayout8BitAccessKHR   Capability = 4429
	CapabilityWorkgroupMemoryExplicitLayout16BitAccessKHR  Capability = 4430
	CapabilitySubgroupVoteKHR                              Capability = 4431
	CapabilityStorageBuffer16BitAccess                     Capability = 4433
	CapabilityStorageUniformBufferBlock16                  Capability = 4433
	CapabilityStorageUniform16                             Capability = 4434
	CapabilityUniformAndStorageBuffer16BitAccess           Capability = 4434
	CapabilityStoragePushConstant16                        Capability = 4435
	CapabilityStorageInputOutput16                         Capability = 4436
	CapabilityDeviceGroup                                  Capability = 4437
	CapabilityMultiView                                    Capability = 4439
	CapabilityVariablePointersStorageBuffer                Capability = 4441
	CapabilityVariablePointers                             Capability = 4442
	CapabilityAtomicStorageOps                             Capability = 4445
	CapabilitySampleMaskPostDepthCoverage                  Capability = 4447
	CapabilityStorageBuffer8BitAccess                      Capability = 4448
	CapabilityUniformAndStorageBuffer8BitAccess            Capability = 4449
	CapabilityStoragePushConstant8                         Capability = 4450
	CapabilityDenormPreserve                               Capability = 4464
	CapabilityDenormFlushToZero                            Capability = 4465
	CapabilitySignedZeroInfNanPreserve                     Capability = 4466
	CapabilityRoundingModeRTE                              Capability = 4467
	CapabilityRoundingModeRTZ                              Capability = 4468
	CapabilityRayQueryProvisionalKHR                       Capability = 4471
	CapabilityRayQueryKHR                                  Capability = 4472
	CapabilityRayTraversalPrimitiveCullingKHR              Capability = 4478
	CapabilityRayTracingKHR                                Capability = 4479
	CapabilityFloat16ImageAMD                              Capability = 5008
	CapabilityImageGatherBiasLodAMD                        Capability = 5009
	CapabilityFragmentMaskAMD                              Capability = 5010
	CapabilityStencilExportEXT                             Capability = 5013
	CapabilityImageReadWriteLodAMD                         Capability = 5015
	CapabilityInt64ImageEXT                                Capability = 5016
	CapabilityShaderClockKHR                               Capability = 5055
	CapabilitySampleMaskOverrideCoverageNV                 Capability = 5249
	CapabilityGeometryShaderPassthroughNV                  Capability = 5251
	CapabilityShaderViewportIndexLayerEXT                  Capability = 5254
	CapabilityShaderViewportIndexLayerNV                   Capability = 5254
	CapabilityShaderViewportMaskNV                         Capability = 5255
	CapabilityShaderStereoViewNV                           Capability = 5259
	CapabilityPerViewAttributesNV                          Capability = 5260
	CapabilityFragmentFullyCoveredEXT                      Capability = 5265
	CapabilityMeshShadingNV                                Capability = 5266
	CapabilityImageFootprintNV                             Capability = 5282
	CapabilityFragmentBarycentricKHR                       Capability = 5284
	CapabilityFragmentBarycentricNV                        Capability = 5284
	CapabilityComputeDerivativeGroupQuadsNV                Capability = 5288
	CapabilityFragmentDensityEXT                           Capability = 5291
	CapabilityShadingRateNV                                Capability = 5291
	CapabilityGroupNonUniformPartitionedNV                 Capability = 5297
	CapabilityShaderNonUniform                             Capability = 5301
	CapabilityShaderNonUniformEXT                          Capability = 5301
	CapabilityRuntimeDescriptorArray                       Capability = 5302
	CapabilityRuntimeDescriptorArrayEXT                    Capability = 5302
	CapabilityInputAttachmentArrayDynamicIndexing          Capability = 5303
	CapabilityInputAttachmentArrayDynamicIndexingEXT       Capability = 5303
	CapabilityUniformTexelBufferArrayDynamicIndexing       Capability = 5304
	CapabilityUniformTexelBufferArrayDynamicIndexingEXT    Capability = 5304
	CapabilityStorageTexelBufferArrayDynamicIndexing       Capability = 5305
	CapabilityStorageTexelBufferArrayDynamicIndexingEXT    Capability = 5305
	CapabilityUniformBufferArrayNonUniformIndexing         Capability = 5306
	CapabilityUniformBufferArrayNonUniformIndexingEXT      Capability = 5306
	CapabilitySampledImageArrayNonUniformIndexing          Capability = 5307
	CapabilitySampledImageArrayNonUniformIndexingEXT       Capability = 5307
	CapabilityStorageBufferArrayNonUniformIndexing         Capability = 5308
	CapabilityStorageBufferArrayNonUniformIndexingEXT      Capability = 5308
	CapabilityStorageImageArrayNonUniformIndexing          Capability = 5309
	CapabilityStorageImageArrayNonUniformIndexingEXT       Capability = 5309
	CapabilityInputAttachmentArrayNonUniformIndexing       Capability = 5310
	CapabilityInputAttachmentArrayNonUniformIndexingEXT    Capability = 5310
	CapabilityUniformTexelBufferArrayNonUniformIndexing    Capability = 5311
	CapabilityUniformTexelBufferArrayNonUniformIndexingEXT Capability = 5311
	CapabilityStorageTexelBufferArrayNonUniformIndexing    Capability = 5312
	CapabilityStorageTexelBufferArrayNonUniformIndexingEXT Capability = 5312
	CapabilityRayTracingNV                                 Capability = 5340
	CapabilityRayTracingMotionBlurNV                       Capability = 5341
	CapabilityVulkanMemoryModel                            Capability = 5345
	CapabilityVulkanMemoryModelKHR                         Capability = 5345
	CapabilityVulkanMemoryModelDeviceScope                 Capability = 5346
	CapabilityVulkanMemoryModelDeviceScopeKHR              Capability = 5346
	CapabilityPhysicalStorageBufferAddresses               Capability = 5347
	CapabilityPhysicalStorageBufferAddressesEXT            Capability = 5347
	CapabilityComputeDerivativeGroupLinearNV               Capability = 5350
	CapabilityRayTracingProvisionalKHR                     Capability = 5353
	CapabilityCooperativeMatrixNV                          Capability = 5357
	CapabilityFragmentShaderSampleInterlockEXT             Capability = 5363
	CapabilityFragmentShaderShadingRateInterlockEXT        Capability = 5372
	CapabilityShaderSMBuiltinsNV                           Capability = 5373
	CapabilityFragmentShaderPixelInterlockEXT              Capability = 5378
	CapabilityDemoteToHelperInvocation                     Capability = 5379
	CapabilityDemoteToHelperInvocationEXT                  Capability = 5379
	CapabilityBindlessTextureNV                            Capability = 5390
	CapabilitySubgroupShuffleINTEL                         Capability = 5568
	CapabilitySubgroupBufferBlockIOINTEL                   Capability = 5569
	CapabilitySubgroupImageBlockIOINTEL                    Capability = 5570
	CapabilitySubgroupImageMediaBlockIOINTEL               Capability = 5579
	CapabilityRoundToInfinityINTEL                         Capability = 5582
	CapabilityFloatingPointModeINTEL                       Capability = 5583
	CapabilityIntegerFunctions2INTEL                       Capability = 5584
	CapabilityFunctionPointersINTEL                        Capability = 5603
	CapabilityIndirectReferencesINTEL                      Capability = 5604
	CapabilityAsmINTEL                                     Capability = 5606
	CapabilityAtomicFloat32MinMaxEXT                       Capability = 5612
	CapabilityAtomicFloat64MinMaxEXT                       Capability = 5613
	CapabilityAtomicFloat16MinMaxEXT                       Capability = 5616
	CapabilityVectorComputeINTEL                           Capability = 5617
	CapabilityVectorAnyINTEL                               Capability = 5619
	CapabilityExpectAssumeKHR                              Capability = 5629
	CapabilitySubgroupAvcMotionEstimationINTEL             Capability = 5696
	CapabilitySubgroupAvcMotionEstimationIntraINTEL        Capability = 5697
	CapabilitySubgroupAvcMotionEstimationChromaINTEL       Capability = 5698
	CapabilityVariableLengthArrayINTEL                     Capability = 5817
	CapabilityFunctionFloatControlINTEL                    Capability = 5821
	CapabilityFPGAMemoryAttributesINTEL                    Capability = 5824
	CapabilityFPFastMathModeINTEL                          Capability = 5837
	CapabilityArbitraryPrecisionIntegersINTEL              Capability = 5844
	CapabilityArbitraryPrecisionFloatingPointINTEL         Capability = 5845
	CapabilityUnstructuredLoopControlsINTEL                Capability = 5886
	CapabilityFPGALoopControlsINTEL                        Capability = 5888
	CapabilityKernelAttributesINTEL                        Capability = 5892
	CapabilityFPGAKernelAttributesINTEL                    Capability = 5897
	CapabilityFPGAMemoryAccessesINTEL                      Capability = 5898
	CapabilityFPGAClusterAttributesINTEL                   Capability = 5904
	CapabilityLoopFuseINTEL                                Capability = 5906
	CapabilityMemoryAccessAliasingINTEL                    Capability = 5910
	CapabilityFPGABufferLocationINTEL                      Capability = 5920
	CapabilityArbitraryPrecisionFixedPointINTEL            Capability = 5922
	CapabilityUSMStorageClassesINTEL                       Capability = 5935
	CapabilityIOPipesINTEL                                 Capability = 5943
	CapabilityBlockingPipesINTEL                           Capability = 5945
	CapabilityFPGARegINTEL                                 Capability = 5948
	CapabilityDotProductInputAll                           Capability = 6016
	CapabilityDotProductInputAllKHR                        Capability = 6016
	CapabilityDotProductInput4x8Bit                        Capability = 6017
	CapabilityDotProductInput4x8BitKHR                     Capability = 6017
	CapabilityDotProductInput4x8BitPacked                  Capability = 6018
	CapabilityDotProductInput4x8BitPackedKHR               Capability = 6018
	CapabilityDotProduct                                   Capability = 6019
	CapabilityDotProductKHR                                Capability = 6019
	CapabilityRayCullMaskKHR                               Capability = 6020
	CapabilityBitInstructions                              Capability = 6025
	CapabilityGroupNonUniformRotateKHR                     Capability = 6026
	CapabilityAtomicFloat32AddEXT                          Capability = 6033
	CapabilityAtomicFloat64AddEXT                          Capability = 6034
	CapabilityLongConstantCompositeINTEL                   Capability = 6089
	CapabilityOptNoneINTEL                                 Capability = 6094
	CapabilityAtomicFloat16AddEXT                          Capability = 6095
	CapabilityDebugInfoModuleINTEL                         Capability = 6114
	CapabilitySplitBarrierINTEL                            Capability = 6141
	CapabilityGroupUniformArithmeticKHR                    Capability = 6400
	CapabilityMax                                          Capability = 0x7fffffff
)
