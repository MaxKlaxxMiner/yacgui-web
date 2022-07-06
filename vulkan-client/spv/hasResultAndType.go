package spv

//goland:noinspection GoUnusedExportedFunction
func (opcode Op) HasResultAndType(hasResult bool, hasResultType bool) {
	switch opcode {
	default: /* unknown opcode */
	case OpNop:
		hasResult = false
		hasResultType = false
	case OpUndef:
		hasResult = true
		hasResultType = true
	case OpSourceContinued:
		hasResult = false
		hasResultType = false
	case OpSource:
		hasResult = false
		hasResultType = false
	case OpSourceExtension:
		hasResult = false
		hasResultType = false
	case OpName:
		hasResult = false
		hasResultType = false
	case OpMemberName:
		hasResult = false
		hasResultType = false
	case OpString:
		hasResult = true
		hasResultType = false
	case OpLine:
		hasResult = false
		hasResultType = false
	case OpExtension:
		hasResult = false
		hasResultType = false
	case OpExtInstImport:
		hasResult = true
		hasResultType = false
	case OpExtInst:
		hasResult = true
		hasResultType = true
	case OpMemoryModel:
		hasResult = false
		hasResultType = false
	case OpEntryPoint:
		hasResult = false
		hasResultType = false
	case OpExecutionMode:
		hasResult = false
		hasResultType = false
	case OpCapability:
		hasResult = false
		hasResultType = false
	case OpTypeVoid:
		hasResult = true
		hasResultType = false
	case OpTypeBool:
		hasResult = true
		hasResultType = false
	case OpTypeInt:
		hasResult = true
		hasResultType = false
	case OpTypeFloat:
		hasResult = true
		hasResultType = false
	case OpTypeVector:
		hasResult = true
		hasResultType = false
	case OpTypeMatrix:
		hasResult = true
		hasResultType = false
	case OpTypeImage:
		hasResult = true
		hasResultType = false
	case OpTypeSampler:
		hasResult = true
		hasResultType = false
	case OpTypeSampledImage:
		hasResult = true
		hasResultType = false
	case OpTypeArray:
		hasResult = true
		hasResultType = false
	case OpTypeRuntimeArray:
		hasResult = true
		hasResultType = false
	case OpTypeStruct:
		hasResult = true
		hasResultType = false
	case OpTypeOpaque:
		hasResult = true
		hasResultType = false
	case OpTypePointer:
		hasResult = true
		hasResultType = false
	case OpTypeFunction:
		hasResult = true
		hasResultType = false
	case OpTypeEvent:
		hasResult = true
		hasResultType = false
	case OpTypeDeviceEvent:
		hasResult = true
		hasResultType = false
	case OpTypeReserveId:
		hasResult = true
		hasResultType = false
	case OpTypeQueue:
		hasResult = true
		hasResultType = false
	case OpTypePipe:
		hasResult = true
		hasResultType = false
	case OpTypeForwardPointer:
		hasResult = false
		hasResultType = false
	case OpConstantTrue:
		hasResult = true
		hasResultType = true
	case OpConstantFalse:
		hasResult = true
		hasResultType = true
	case OpConstant:
		hasResult = true
		hasResultType = true
	case OpConstantComposite:
		hasResult = true
		hasResultType = true
	case OpConstantSampler:
		hasResult = true
		hasResultType = true
	case OpConstantNull:
		hasResult = true
		hasResultType = true
	case OpSpecConstantTrue:
		hasResult = true
		hasResultType = true
	case OpSpecConstantFalse:
		hasResult = true
		hasResultType = true
	case OpSpecConstant:
		hasResult = true
		hasResultType = true
	case OpSpecConstantComposite:
		hasResult = true
		hasResultType = true
	case OpSpecConstantOp:
		hasResult = true
		hasResultType = true
	case OpFunction:
		hasResult = true
		hasResultType = true
	case OpFunctionParameter:
		hasResult = true
		hasResultType = true
	case OpFunctionEnd:
		hasResult = false
		hasResultType = false
	case OpFunctionCall:
		hasResult = true
		hasResultType = true
	case OpVariable:
		hasResult = true
		hasResultType = true
	case OpImageTexelPointer:
		hasResult = true
		hasResultType = true
	case OpLoad:
		hasResult = true
		hasResultType = true
	case OpStore:
		hasResult = false
		hasResultType = false
	case OpCopyMemory:
		hasResult = false
		hasResultType = false
	case OpCopyMemorySized:
		hasResult = false
		hasResultType = false
	case OpAccessChain:
		hasResult = true
		hasResultType = true
	case OpInBoundsAccessChain:
		hasResult = true
		hasResultType = true
	case OpPtrAccessChain:
		hasResult = true
		hasResultType = true
	case OpArrayLength:
		hasResult = true
		hasResultType = true
	case OpGenericPtrMemSemantics:
		hasResult = true
		hasResultType = true
	case OpInBoundsPtrAccessChain:
		hasResult = true
		hasResultType = true
	case OpDecorate:
		hasResult = false
		hasResultType = false
	case OpMemberDecorate:
		hasResult = false
		hasResultType = false
	case OpDecorationGroup:
		hasResult = true
		hasResultType = false
	case OpGroupDecorate:
		hasResult = false
		hasResultType = false
	case OpGroupMemberDecorate:
		hasResult = false
		hasResultType = false
	case OpVectorExtractDynamic:
		hasResult = true
		hasResultType = true
	case OpVectorInsertDynamic:
		hasResult = true
		hasResultType = true
	case OpVectorShuffle:
		hasResult = true
		hasResultType = true
	case OpCompositeConstruct:
		hasResult = true
		hasResultType = true
	case OpCompositeExtract:
		hasResult = true
		hasResultType = true
	case OpCompositeInsert:
		hasResult = true
		hasResultType = true
	case OpCopyObject:
		hasResult = true
		hasResultType = true
	case OpTranspose:
		hasResult = true
		hasResultType = true
	case OpSampledImage:
		hasResult = true
		hasResultType = true
	case OpImageSampleImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleDrefImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleDrefExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleProjImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleProjExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleProjDrefImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSampleProjDrefExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageFetch:
		hasResult = true
		hasResultType = true
	case OpImageGather:
		hasResult = true
		hasResultType = true
	case OpImageDrefGather:
		hasResult = true
		hasResultType = true
	case OpImageRead:
		hasResult = true
		hasResultType = true
	case OpImageWrite:
		hasResult = false
		hasResultType = false
	case OpImage:
		hasResult = true
		hasResultType = true
	case OpImageQueryFormat:
		hasResult = true
		hasResultType = true
	case OpImageQueryOrder:
		hasResult = true
		hasResultType = true
	case OpImageQuerySizeLod:
		hasResult = true
		hasResultType = true
	case OpImageQuerySize:
		hasResult = true
		hasResultType = true
	case OpImageQueryLod:
		hasResult = true
		hasResultType = true
	case OpImageQueryLevels:
		hasResult = true
		hasResultType = true
	case OpImageQuerySamples:
		hasResult = true
		hasResultType = true
	case OpConvertFToU:
		hasResult = true
		hasResultType = true
	case OpConvertFToS:
		hasResult = true
		hasResultType = true
	case OpConvertSToF:
		hasResult = true
		hasResultType = true
	case OpConvertUToF:
		hasResult = true
		hasResultType = true
	case OpUConvert:
		hasResult = true
		hasResultType = true
	case OpSConvert:
		hasResult = true
		hasResultType = true
	case OpFConvert:
		hasResult = true
		hasResultType = true
	case OpQuantizeToF16:
		hasResult = true
		hasResultType = true
	case OpConvertPtrToU:
		hasResult = true
		hasResultType = true
	case OpSatConvertSToU:
		hasResult = true
		hasResultType = true
	case OpSatConvertUToS:
		hasResult = true
		hasResultType = true
	case OpConvertUToPtr:
		hasResult = true
		hasResultType = true
	case OpPtrCastToGeneric:
		hasResult = true
		hasResultType = true
	case OpGenericCastToPtr:
		hasResult = true
		hasResultType = true
	case OpGenericCastToPtrExplicit:
		hasResult = true
		hasResultType = true
	case OpBitcast:
		hasResult = true
		hasResultType = true
	case OpSNegate:
		hasResult = true
		hasResultType = true
	case OpFNegate:
		hasResult = true
		hasResultType = true
	case OpIAdd:
		hasResult = true
		hasResultType = true
	case OpFAdd:
		hasResult = true
		hasResultType = true
	case OpISub:
		hasResult = true
		hasResultType = true
	case OpFSub:
		hasResult = true
		hasResultType = true
	case OpIMul:
		hasResult = true
		hasResultType = true
	case OpFMul:
		hasResult = true
		hasResultType = true
	case OpUDiv:
		hasResult = true
		hasResultType = true
	case OpSDiv:
		hasResult = true
		hasResultType = true
	case OpFDiv:
		hasResult = true
		hasResultType = true
	case OpUMod:
		hasResult = true
		hasResultType = true
	case OpSRem:
		hasResult = true
		hasResultType = true
	case OpSMod:
		hasResult = true
		hasResultType = true
	case OpFRem:
		hasResult = true
		hasResultType = true
	case OpFMod:
		hasResult = true
		hasResultType = true
	case OpVectorTimesScalar:
		hasResult = true
		hasResultType = true
	case OpMatrixTimesScalar:
		hasResult = true
		hasResultType = true
	case OpVectorTimesMatrix:
		hasResult = true
		hasResultType = true
	case OpMatrixTimesVector:
		hasResult = true
		hasResultType = true
	case OpMatrixTimesMatrix:
		hasResult = true
		hasResultType = true
	case OpOuterProduct:
		hasResult = true
		hasResultType = true
	case OpDot:
		hasResult = true
		hasResultType = true
	case OpIAddCarry:
		hasResult = true
		hasResultType = true
	case OpISubBorrow:
		hasResult = true
		hasResultType = true
	case OpUMulExtended:
		hasResult = true
		hasResultType = true
	case OpSMulExtended:
		hasResult = true
		hasResultType = true
	case OpAny:
		hasResult = true
		hasResultType = true
	case OpAll:
		hasResult = true
		hasResultType = true
	case OpIsNan:
		hasResult = true
		hasResultType = true
	case OpIsInf:
		hasResult = true
		hasResultType = true
	case OpIsFinite:
		hasResult = true
		hasResultType = true
	case OpIsNormal:
		hasResult = true
		hasResultType = true
	case OpSignBitSet:
		hasResult = true
		hasResultType = true
	case OpLessOrGreater:
		hasResult = true
		hasResultType = true
	case OpOrdered:
		hasResult = true
		hasResultType = true
	case OpUnordered:
		hasResult = true
		hasResultType = true
	case OpLogicalEqual:
		hasResult = true
		hasResultType = true
	case OpLogicalNotEqual:
		hasResult = true
		hasResultType = true
	case OpLogicalOr:
		hasResult = true
		hasResultType = true
	case OpLogicalAnd:
		hasResult = true
		hasResultType = true
	case OpLogicalNot:
		hasResult = true
		hasResultType = true
	case OpSelect:
		hasResult = true
		hasResultType = true
	case OpIEqual:
		hasResult = true
		hasResultType = true
	case OpINotEqual:
		hasResult = true
		hasResultType = true
	case OpUGreaterThan:
		hasResult = true
		hasResultType = true
	case OpSGreaterThan:
		hasResult = true
		hasResultType = true
	case OpUGreaterThanEqual:
		hasResult = true
		hasResultType = true
	case OpSGreaterThanEqual:
		hasResult = true
		hasResultType = true
	case OpULessThan:
		hasResult = true
		hasResultType = true
	case OpSLessThan:
		hasResult = true
		hasResultType = true
	case OpULessThanEqual:
		hasResult = true
		hasResultType = true
	case OpSLessThanEqual:
		hasResult = true
		hasResultType = true
	case OpFOrdEqual:
		hasResult = true
		hasResultType = true
	case OpFUnordEqual:
		hasResult = true
		hasResultType = true
	case OpFOrdNotEqual:
		hasResult = true
		hasResultType = true
	case OpFUnordNotEqual:
		hasResult = true
		hasResultType = true
	case OpFOrdLessThan:
		hasResult = true
		hasResultType = true
	case OpFUnordLessThan:
		hasResult = true
		hasResultType = true
	case OpFOrdGreaterThan:
		hasResult = true
		hasResultType = true
	case OpFUnordGreaterThan:
		hasResult = true
		hasResultType = true
	case OpFOrdLessThanEqual:
		hasResult = true
		hasResultType = true
	case OpFUnordLessThanEqual:
		hasResult = true
		hasResultType = true
	case OpFOrdGreaterThanEqual:
		hasResult = true
		hasResultType = true
	case OpFUnordGreaterThanEqual:
		hasResult = true
		hasResultType = true
	case OpShiftRightLogical:
		hasResult = true
		hasResultType = true
	case OpShiftRightArithmetic:
		hasResult = true
		hasResultType = true
	case OpShiftLeftLogical:
		hasResult = true
		hasResultType = true
	case OpBitwiseOr:
		hasResult = true
		hasResultType = true
	case OpBitwiseXor:
		hasResult = true
		hasResultType = true
	case OpBitwiseAnd:
		hasResult = true
		hasResultType = true
	case OpNot:
		hasResult = true
		hasResultType = true
	case OpBitFieldInsert:
		hasResult = true
		hasResultType = true
	case OpBitFieldSExtract:
		hasResult = true
		hasResultType = true
	case OpBitFieldUExtract:
		hasResult = true
		hasResultType = true
	case OpBitReverse:
		hasResult = true
		hasResultType = true
	case OpBitCount:
		hasResult = true
		hasResultType = true
	case OpDPdx:
		hasResult = true
		hasResultType = true
	case OpDPdy:
		hasResult = true
		hasResultType = true
	case OpFwidth:
		hasResult = true
		hasResultType = true
	case OpDPdxFine:
		hasResult = true
		hasResultType = true
	case OpDPdyFine:
		hasResult = true
		hasResultType = true
	case OpFwidthFine:
		hasResult = true
		hasResultType = true
	case OpDPdxCoarse:
		hasResult = true
		hasResultType = true
	case OpDPdyCoarse:
		hasResult = true
		hasResultType = true
	case OpFwidthCoarse:
		hasResult = true
		hasResultType = true
	case OpEmitVertex:
		hasResult = false
		hasResultType = false
	case OpEndPrimitive:
		hasResult = false
		hasResultType = false
	case OpEmitStreamVertex:
		hasResult = false
		hasResultType = false
	case OpEndStreamPrimitive:
		hasResult = false
		hasResultType = false
	case OpControlBarrier:
		hasResult = false
		hasResultType = false
	case OpMemoryBarrier:
		hasResult = false
		hasResultType = false
	case OpAtomicLoad:
		hasResult = true
		hasResultType = true
	case OpAtomicStore:
		hasResult = false
		hasResultType = false
	case OpAtomicExchange:
		hasResult = true
		hasResultType = true
	case OpAtomicCompareExchange:
		hasResult = true
		hasResultType = true
	case OpAtomicCompareExchangeWeak:
		hasResult = true
		hasResultType = true
	case OpAtomicIIncrement:
		hasResult = true
		hasResultType = true
	case OpAtomicIDecrement:
		hasResult = true
		hasResultType = true
	case OpAtomicIAdd:
		hasResult = true
		hasResultType = true
	case OpAtomicISub:
		hasResult = true
		hasResultType = true
	case OpAtomicSMin:
		hasResult = true
		hasResultType = true
	case OpAtomicUMin:
		hasResult = true
		hasResultType = true
	case OpAtomicSMax:
		hasResult = true
		hasResultType = true
	case OpAtomicUMax:
		hasResult = true
		hasResultType = true
	case OpAtomicAnd:
		hasResult = true
		hasResultType = true
	case OpAtomicOr:
		hasResult = true
		hasResultType = true
	case OpAtomicXor:
		hasResult = true
		hasResultType = true
	case OpPhi:
		hasResult = true
		hasResultType = true
	case OpLoopMerge:
		hasResult = false
		hasResultType = false
	case OpSelectionMerge:
		hasResult = false
		hasResultType = false
	case OpLabel:
		hasResult = true
		hasResultType = false
	case OpBranch:
		hasResult = false
		hasResultType = false
	case OpBranchConditional:
		hasResult = false
		hasResultType = false
	case OpSwitch:
		hasResult = false
		hasResultType = false
	case OpKill:
		hasResult = false
		hasResultType = false
	case OpReturn:
		hasResult = false
		hasResultType = false
	case OpReturnValue:
		hasResult = false
		hasResultType = false
	case OpUnreachable:
		hasResult = false
		hasResultType = false
	case OpLifetimeStart:
		hasResult = false
		hasResultType = false
	case OpLifetimeStop:
		hasResult = false
		hasResultType = false
	case OpGroupAsyncCopy:
		hasResult = true
		hasResultType = true
	case OpGroupWaitEvents:
		hasResult = false
		hasResultType = false
	case OpGroupAll:
		hasResult = true
		hasResultType = true
	case OpGroupAny:
		hasResult = true
		hasResultType = true
	case OpGroupBroadcast:
		hasResult = true
		hasResultType = true
	case OpGroupIAdd:
		hasResult = true
		hasResultType = true
	case OpGroupFAdd:
		hasResult = true
		hasResultType = true
	case OpGroupFMin:
		hasResult = true
		hasResultType = true
	case OpGroupUMin:
		hasResult = true
		hasResultType = true
	case OpGroupSMin:
		hasResult = true
		hasResultType = true
	case OpGroupFMax:
		hasResult = true
		hasResultType = true
	case OpGroupUMax:
		hasResult = true
		hasResultType = true
	case OpGroupSMax:
		hasResult = true
		hasResultType = true
	case OpReadPipe:
		hasResult = true
		hasResultType = true
	case OpWritePipe:
		hasResult = true
		hasResultType = true
	case OpReservedReadPipe:
		hasResult = true
		hasResultType = true
	case OpReservedWritePipe:
		hasResult = true
		hasResultType = true
	case OpReserveReadPipePackets:
		hasResult = true
		hasResultType = true
	case OpReserveWritePipePackets:
		hasResult = true
		hasResultType = true
	case OpCommitReadPipe:
		hasResult = false
		hasResultType = false
	case OpCommitWritePipe:
		hasResult = false
		hasResultType = false
	case OpIsValidReserveId:
		hasResult = true
		hasResultType = true
	case OpGetNumPipePackets:
		hasResult = true
		hasResultType = true
	case OpGetMaxPipePackets:
		hasResult = true
		hasResultType = true
	case OpGroupReserveReadPipePackets:
		hasResult = true
		hasResultType = true
	case OpGroupReserveWritePipePackets:
		hasResult = true
		hasResultType = true
	case OpGroupCommitReadPipe:
		hasResult = false
		hasResultType = false
	case OpGroupCommitWritePipe:
		hasResult = false
		hasResultType = false
	case OpEnqueueMarker:
		hasResult = true
		hasResultType = true
	case OpEnqueueKernel:
		hasResult = true
		hasResultType = true
	case OpGetKernelNDrangeSubGroupCount:
		hasResult = true
		hasResultType = true
	case OpGetKernelNDrangeMaxSubGroupSize:
		hasResult = true
		hasResultType = true
	case OpGetKernelWorkGroupSize:
		hasResult = true
		hasResultType = true
	case OpGetKernelPreferredWorkGroupSizeMultiple:
		hasResult = true
		hasResultType = true
	case OpRetainEvent:
		hasResult = false
		hasResultType = false
	case OpReleaseEvent:
		hasResult = false
		hasResultType = false
	case OpCreateUserEvent:
		hasResult = true
		hasResultType = true
	case OpIsValidEvent:
		hasResult = true
		hasResultType = true
	case OpSetUserEventStatus:
		hasResult = false
		hasResultType = false
	case OpCaptureEventProfilingInfo:
		hasResult = false
		hasResultType = false
	case OpGetDefaultQueue:
		hasResult = true
		hasResultType = true
	case OpBuildNDRange:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleDrefImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleDrefExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleProjImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleProjExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleProjDrefImplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseSampleProjDrefExplicitLod:
		hasResult = true
		hasResultType = true
	case OpImageSparseFetch:
		hasResult = true
		hasResultType = true
	case OpImageSparseGather:
		hasResult = true
		hasResultType = true
	case OpImageSparseDrefGather:
		hasResult = true
		hasResultType = true
	case OpImageSparseTexelsResident:
		hasResult = true
		hasResultType = true
	case OpNoLine:
		hasResult = false
		hasResultType = false
	case OpAtomicFlagTestAndSet:
		hasResult = true
		hasResultType = true
	case OpAtomicFlagClear:
		hasResult = false
		hasResultType = false
	case OpImageSparseRead:
		hasResult = true
		hasResultType = true
	case OpSizeOf:
		hasResult = true
		hasResultType = true
	case OpTypePipeStorage:
		hasResult = true
		hasResultType = false
	case OpConstantPipeStorage:
		hasResult = true
		hasResultType = true
	case OpCreatePipeFromPipeStorage:
		hasResult = true
		hasResultType = true
	case OpGetKernelLocalSizeForSubgroupCount:
		hasResult = true
		hasResultType = true
	case OpGetKernelMaxNumSubgroups:
		hasResult = true
		hasResultType = true
	case OpTypeNamedBarrier:
		hasResult = true
		hasResultType = false
	case OpNamedBarrierInitialize:
		hasResult = true
		hasResultType = true
	case OpMemoryNamedBarrier:
		hasResult = false
		hasResultType = false
	case OpModuleProcessed:
		hasResult = false
		hasResultType = false
	case OpExecutionModeId:
		hasResult = false
		hasResultType = false
	case OpDecorateId:
		hasResult = false
		hasResultType = false
	case OpGroupNonUniformElect:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformAll:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformAny:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformAllEqual:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBroadcast:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBroadcastFirst:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBallot:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformInverseBallot:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBallotBitExtract:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBallotBitCount:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBallotFindLSB:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBallotFindMSB:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformShuffle:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformShuffleXor:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformShuffleUp:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformShuffleDown:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformIAdd:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformFAdd:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformIMul:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformFMul:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformSMin:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformUMin:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformFMin:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformSMax:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformUMax:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformFMax:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBitwiseAnd:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBitwiseOr:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformBitwiseXor:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformLogicalAnd:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformLogicalOr:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformLogicalXor:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformQuadBroadcast:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformQuadSwap:
		hasResult = true
		hasResultType = true
	case OpCopyLogical:
		hasResult = true
		hasResultType = true
	case OpPtrEqual:
		hasResult = true
		hasResultType = true
	case OpPtrNotEqual:
		hasResult = true
		hasResultType = true
	case OpPtrDiff:
		hasResult = true
		hasResultType = true
	case OpTerminateInvocation:
		hasResult = false
		hasResultType = false
	case OpSubgroupBallotKHR:
		hasResult = true
		hasResultType = true
	case OpSubgroupFirstInvocationKHR:
		hasResult = true
		hasResultType = true
	case OpSubgroupAllKHR:
		hasResult = true
		hasResultType = true
	case OpSubgroupAnyKHR:
		hasResult = true
		hasResultType = true
	case OpSubgroupAllEqualKHR:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformRotateKHR:
		hasResult = true
		hasResultType = true
	case OpSubgroupReadInvocationKHR:
		hasResult = true
		hasResultType = true
	case OpTraceRayKHR:
		hasResult = false
		hasResultType = false
	case OpExecuteCallableKHR:
		hasResult = false
		hasResultType = false
	case OpConvertUToAccelerationStructureKHR:
		hasResult = true
		hasResultType = true
	case OpIgnoreIntersectionKHR:
		hasResult = false
		hasResultType = false
	case OpTerminateRayKHR:
		hasResult = false
		hasResultType = false
	case OpSDot:
		hasResult = true
		hasResultType = true
	case OpUDot:
		hasResult = true
		hasResultType = true
	case OpSUDot:
		hasResult = true
		hasResultType = true
	case OpSDotAccSat:
		hasResult = true
		hasResultType = true
	case OpUDotAccSat:
		hasResult = true
		hasResultType = true
	case OpSUDotAccSat:
		hasResult = true
		hasResultType = true
	case OpTypeRayQueryKHR:
		hasResult = true
		hasResultType = false
	case OpRayQueryInitializeKHR:
		hasResult = false
		hasResultType = false
	case OpRayQueryTerminateKHR:
		hasResult = false
		hasResultType = false
	case OpRayQueryGenerateIntersectionKHR:
		hasResult = false
		hasResultType = false
	case OpRayQueryConfirmIntersectionKHR:
		hasResult = false
		hasResultType = false
	case OpRayQueryProceedKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionTypeKHR:
		hasResult = true
		hasResultType = true
	case OpGroupIAddNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupFAddNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupFMinNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupUMinNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupSMinNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupFMaxNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupUMaxNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpGroupSMaxNonUniformAMD:
		hasResult = true
		hasResultType = true
	case OpFragmentMaskFetchAMD:
		hasResult = true
		hasResultType = true
	case OpFragmentFetchAMD:
		hasResult = true
		hasResultType = true
	case OpReadClockKHR:
		hasResult = true
		hasResultType = true
	case OpImageSampleFootprintNV:
		hasResult = true
		hasResultType = true
	case OpGroupNonUniformPartitionNV:
		hasResult = true
		hasResultType = true
	case OpWritePackedPrimitiveIndices4x8NV:
		hasResult = false
		hasResultType = false
	case OpReportIntersectionNV:
		hasResult = true
		hasResultType = true
	case OpIgnoreIntersectionNV:
		hasResult = false
		hasResultType = false
	case OpTerminateRayNV:
		hasResult = false
		hasResultType = false
	case OpTraceNV:
		hasResult = false
		hasResultType = false
	case OpTraceMotionNV:
		hasResult = false
		hasResultType = false
	case OpTraceRayMotionNV:
		hasResult = false
		hasResultType = false
	case OpTypeAccelerationStructureNV:
		hasResult = true
		hasResultType = false
	case OpExecuteCallableNV:
		hasResult = false
		hasResultType = false
	case OpTypeCooperativeMatrixNV:
		hasResult = true
		hasResultType = false
	case OpCooperativeMatrixLoadNV:
		hasResult = true
		hasResultType = true
	case OpCooperativeMatrixStoreNV:
		hasResult = false
		hasResultType = false
	case OpCooperativeMatrixMulAddNV:
		hasResult = true
		hasResultType = true
	case OpCooperativeMatrixLengthNV:
		hasResult = true
		hasResultType = true
	case OpBeginInvocationInterlockEXT:
		hasResult = false
		hasResultType = false
	case OpEndInvocationInterlockEXT:
		hasResult = false
		hasResultType = false
	case OpDemoteToHelperInvocation:
		hasResult = false
		hasResultType = false
	case OpIsHelperInvocationEXT:
		hasResult = true
		hasResultType = true
	case OpConvertUToImageNV:
		hasResult = true
		hasResultType = true
	case OpConvertUToSamplerNV:
		hasResult = true
		hasResultType = true
	case OpConvertImageToUNV:
		hasResult = true
		hasResultType = true
	case OpConvertSamplerToUNV:
		hasResult = true
		hasResultType = true
	case OpConvertUToSampledImageNV:
		hasResult = true
		hasResultType = true
	case OpConvertSampledImageToUNV:
		hasResult = true
		hasResultType = true
	case OpSamplerImageAddressingModeNV:
		hasResult = false
		hasResultType = false
	case OpSubgroupShuffleINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupShuffleDownINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupShuffleUpINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupShuffleXorINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupBlockReadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupBlockWriteINTEL:
		hasResult = false
		hasResultType = false
	case OpSubgroupImageBlockReadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupImageBlockWriteINTEL:
		hasResult = false
		hasResultType = false
	case OpSubgroupImageMediaBlockReadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupImageMediaBlockWriteINTEL:
		hasResult = false
		hasResultType = false
	case OpUCountLeadingZerosINTEL:
		hasResult = true
		hasResultType = true
	case OpUCountTrailingZerosINTEL:
		hasResult = true
		hasResultType = true
	case OpAbsISubINTEL:
		hasResult = true
		hasResultType = true
	case OpAbsUSubINTEL:
		hasResult = true
		hasResultType = true
	case OpIAddSatINTEL:
		hasResult = true
		hasResultType = true
	case OpUAddSatINTEL:
		hasResult = true
		hasResultType = true
	case OpIAverageINTEL:
		hasResult = true
		hasResultType = true
	case OpUAverageINTEL:
		hasResult = true
		hasResultType = true
	case OpIAverageRoundedINTEL:
		hasResult = true
		hasResultType = true
	case OpUAverageRoundedINTEL:
		hasResult = true
		hasResultType = true
	case OpISubSatINTEL:
		hasResult = true
		hasResultType = true
	case OpUSubSatINTEL:
		hasResult = true
		hasResultType = true
	case OpIMul32x16INTEL:
		hasResult = true
		hasResultType = true
	case OpUMul32x16INTEL:
		hasResult = true
		hasResultType = true
	case OpConstantFunctionPointerINTEL:
		hasResult = true
		hasResultType = true
	case OpFunctionPointerCallINTEL:
		hasResult = true
		hasResultType = true
	case OpAsmTargetINTEL:
		hasResult = true
		hasResultType = true
	case OpAsmINTEL:
		hasResult = true
		hasResultType = true
	case OpAsmCallINTEL:
		hasResult = true
		hasResultType = true
	case OpAtomicFMinEXT:
		hasResult = true
		hasResultType = true
	case OpAtomicFMaxEXT:
		hasResult = true
		hasResultType = true
	case OpAssumeTrueKHR:
		hasResult = false
		hasResultType = false
	case OpExpectKHR:
		hasResult = true
		hasResultType = true
	case OpDecorateString:
		hasResult = false
		hasResultType = false
	case OpMemberDecorateString:
		hasResult = false
		hasResultType = false
	case OpVmeImageINTEL:
		hasResult = true
		hasResultType = true
	case OpTypeVmeImageINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcImePayloadINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcRefPayloadINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcSicPayloadINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcMcePayloadINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcMceResultINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcImeResultINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcImeResultSingleReferenceStreamoutINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcImeResultDualReferenceStreamoutINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcImeSingleReferenceStreaminINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcImeDualReferenceStreaminINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcRefResultINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeAvcSicResultINTEL:
		hasResult = true
		hasResultType = false
	case OpSubgroupAvcMceGetDefaultInterBaseMultiReferencePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetInterBaseMultiReferencePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultInterShapePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetInterShapePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultInterDirectionPenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetInterDirectionPenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultIntraLumaShapePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultInterMotionVectorCostTableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultHighPenaltyCostTableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultMediumPenaltyCostTableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultLowPenaltyCostTableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetMotionVectorCostFunctionINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultIntraLumaModePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultNonDcLumaIntraPenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetDefaultIntraChromaModeBasePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetAcOnlyHaarINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetSourceInterlacedFieldPolarityINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetSingleReferenceInterlacedFieldPolarityINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceSetDualReferenceInterlacedFieldPolaritiesINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceConvertToImePayloadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceConvertToImeResultINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceConvertToRefPayloadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceConvertToRefResultINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceConvertToSicPayloadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceConvertToSicResultINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetMotionVectorsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterDistortionsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetBestInterDistortionsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterMajorShapeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterMinorShapeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterDirectionsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterMotionVectorCountINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterReferenceIdsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcMceGetInterReferenceInterlacedFieldPolaritiesINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeInitializeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeSetSingleReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeSetDualReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeRefWindowSizeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeAdjustRefOffsetINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeConvertToMcePayloadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeSetMaxMotionVectorCountINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeSetUnidirectionalMixDisableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeSetEarlySearchTerminationThresholdINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeSetWeightedSadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithSingleReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithDualReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithSingleReferenceStreaminINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithDualReferenceStreaminINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithSingleReferenceStreamoutINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithDualReferenceStreamoutINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithSingleReferenceStreaminoutINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeEvaluateWithDualReferenceStreaminoutINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeConvertToMceResultINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetSingleReferenceStreaminINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetDualReferenceStreaminINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeStripSingleReferenceStreamoutINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeStripDualReferenceStreamoutINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetStreamoutSingleReferenceMajorShapeMotionVectorsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetStreamoutSingleReferenceMajorShapeDistortionsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetStreamoutSingleReferenceMajorShapeReferenceIdsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetStreamoutDualReferenceMajorShapeMotionVectorsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetStreamoutDualReferenceMajorShapeDistortionsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetStreamoutDualReferenceMajorShapeReferenceIdsINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetBorderReachedINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetTruncatedSearchIndicationINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetUnidirectionalEarlySearchTerminationINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetWeightingPatternMinimumMotionVectorINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcImeGetWeightingPatternMinimumDistortionINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcFmeInitializeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcBmeInitializeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefConvertToMcePayloadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefSetBidirectionalMixDisableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefSetBilinearFilterEnableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefEvaluateWithSingleReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefEvaluateWithDualReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefEvaluateWithMultiReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefEvaluateWithMultiReferenceInterlacedINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcRefConvertToMceResultINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicInitializeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicConfigureSkcINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicConfigureIpeLumaINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicConfigureIpeLumaChromaINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetMotionVectorMaskINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicConvertToMcePayloadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicSetIntraLumaShapePenaltyINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicSetIntraLumaModeCostFunctionINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicSetIntraChromaModeCostFunctionINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicSetBilinearFilterEnableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicSetSkcForwardTransformEnableINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicSetBlockBasedRawSkipSadINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicEvaluateIpeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicEvaluateWithSingleReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicEvaluateWithDualReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicEvaluateWithMultiReferenceINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicEvaluateWithMultiReferenceInterlacedINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicConvertToMceResultINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetIpeLumaShapeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetBestIpeLumaDistortionINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetBestIpeChromaDistortionINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetPackedIpeLumaModesINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetIpeChromaModeINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetPackedSkcLumaCountThresholdINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetPackedSkcLumaSumThresholdINTEL:
		hasResult = true
		hasResultType = true
	case OpSubgroupAvcSicGetInterRawSadsINTEL:
		hasResult = true
		hasResultType = true
	case OpVariableLengthArrayINTEL:
		hasResult = true
		hasResultType = true
	case OpSaveMemoryINTEL:
		hasResult = true
		hasResultType = true
	case OpRestoreMemoryINTEL:
		hasResult = false
		hasResultType = false
	case OpArbitraryFloatSinCosPiINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatCastINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatCastFromIntINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatCastToIntINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatAddINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatSubINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatMulINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatDivINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatGTINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatGEINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatLTINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatLEINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatEQINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatRecipINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatRSqrtINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatCbrtINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatHypotINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatSqrtINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatLogINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatLog2INTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatLog10INTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatLog1pINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatExpINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatExp2INTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatExp10INTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatExpm1INTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatSinINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatCosINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatSinCosINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatSinPiINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatCosPiINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatASinINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatASinPiINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatACosINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatACosPiINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatATanINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatATanPiINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatATan2INTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatPowINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatPowRINTEL:
		hasResult = true
		hasResultType = true
	case OpArbitraryFloatPowNINTEL:
		hasResult = true
		hasResultType = true
	case OpLoopControlINTEL:
		hasResult = false
		hasResultType = false
	case OpAliasDomainDeclINTEL:
		hasResult = true
		hasResultType = false
	case OpAliasScopeDeclINTEL:
		hasResult = true
		hasResultType = false
	case OpAliasScopeListDeclINTEL:
		hasResult = true
		hasResultType = false
	case OpFixedSqrtINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedRecipINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedRsqrtINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedSinINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedCosINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedSinCosINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedSinPiINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedCosPiINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedSinCosPiINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedLogINTEL:
		hasResult = true
		hasResultType = true
	case OpFixedExpINTEL:
		hasResult = true
		hasResultType = true
	case OpPtrCastToCrossWorkgroupINTEL:
		hasResult = true
		hasResultType = true
	case OpCrossWorkgroupCastToPtrINTEL:
		hasResult = true
		hasResultType = true
	case OpReadPipeBlockingINTEL:
		hasResult = true
		hasResultType = true
	case OpWritePipeBlockingINTEL:
		hasResult = true
		hasResultType = true
	case OpFPGARegINTEL:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetRayTMinKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetRayFlagsKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionTKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionInstanceCustomIndexKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionInstanceIdKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionInstanceShaderBindingTableRecordOffsetKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionGeometryIndexKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionPrimitiveIndexKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionBarycentricsKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionFrontFaceKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionCandidateAABBOpaqueKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionObjectRayDirectionKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionObjectRayOriginKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetWorldRayDirectionKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetWorldRayOriginKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionObjectToWorldKHR:
		hasResult = true
		hasResultType = true
	case OpRayQueryGetIntersectionWorldToObjectKHR:
		hasResult = true
		hasResultType = true
	case OpAtomicFAddEXT:
		hasResult = true
		hasResultType = true
	case OpTypeBufferSurfaceINTEL:
		hasResult = true
		hasResultType = false
	case OpTypeStructContinuedINTEL:
		hasResult = false
		hasResultType = false
	case OpConstantCompositeContinuedINTEL:
		hasResult = false
		hasResultType = false
	case OpSpecConstantCompositeContinuedINTEL:
		hasResult = false
		hasResultType = false
	case OpControlBarrierArriveINTEL:
		hasResult = false
		hasResultType = false
	case OpControlBarrierWaitINTEL:
		hasResult = false
		hasResultType = false
	case OpGroupIMulKHR:
		hasResult = true
		hasResultType = true
	case OpGroupFMulKHR:
		hasResult = true
		hasResultType = true
	case OpGroupBitwiseAndKHR:
		hasResult = true
		hasResultType = true
	case OpGroupBitwiseOrKHR:
		hasResult = true
		hasResultType = true
	case OpGroupBitwiseXorKHR:
		hasResult = true
		hasResultType = true
	case OpGroupLogicalAndKHR:
		hasResult = true
		hasResultType = true
	case OpGroupLogicalOrKHR:
		hasResult = true
		hasResultType = true
	case OpGroupLogicalXorKHR:
		hasResult = true
		hasResultType = true
	}
	return
}
