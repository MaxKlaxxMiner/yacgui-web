package spv

type StorageClass uint32

const (
	StorageClassUniformConstant          StorageClass = 0
	StorageClassInput                    StorageClass = 1
	StorageClassUniform                  StorageClass = 2
	StorageClassOutput                   StorageClass = 3
	StorageClassWorkgroup                StorageClass = 4
	StorageClassCrossWorkgroup           StorageClass = 5
	StorageClassPrivate                  StorageClass = 6
	StorageClassFunction                 StorageClass = 7
	StorageClassGeneric                  StorageClass = 8
	StorageClassPushConstant             StorageClass = 9
	StorageClassAtomicCounter            StorageClass = 10
	StorageClassImage                    StorageClass = 11
	StorageClassStorageBuffer            StorageClass = 12
	StorageClassCallableDataKHR          StorageClass = 5328
	StorageClassCallableDataNV           StorageClass = 5328
	StorageClassIncomingCallableDataKHR  StorageClass = 5329
	StorageClassIncomingCallableDataNV   StorageClass = 5329
	StorageClassRayPayloadKHR            StorageClass = 5338
	StorageClassRayPayloadNV             StorageClass = 5338
	StorageClassHitAttributeKHR          StorageClass = 5339
	StorageClassHitAttributeNV           StorageClass = 5339
	StorageClassIncomingRayPayloadKHR    StorageClass = 5342
	StorageClassIncomingRayPayloadNV     StorageClass = 5342
	StorageClassShaderRecordBufferKHR    StorageClass = 5343
	StorageClassShaderRecordBufferNV     StorageClass = 5343
	StorageClassPhysicalStorageBuffer    StorageClass = 5349
	StorageClassPhysicalStorageBufferEXT StorageClass = 5349
	StorageClassCodeSectionINTEL         StorageClass = 5605
	StorageClassDeviceOnlyINTEL          StorageClass = 5936
	StorageClassHostOnlyINTEL            StorageClass = 5937
	StorageClassMax                      StorageClass = 0x7fffffff
)
