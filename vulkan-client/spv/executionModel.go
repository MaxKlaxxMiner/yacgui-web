package spv

type ExecutionModel uint32

//goland:noinspection GoUnusedConst
const (
	ExecutionModelVertex                 ExecutionModel = 0
	ExecutionModelTessellationControl    ExecutionModel = 1
	ExecutionModelTessellationEvaluation ExecutionModel = 2
	ExecutionModelGeometry               ExecutionModel = 3
	ExecutionModelFragment               ExecutionModel = 4
	ExecutionModelGLCompute              ExecutionModel = 5
	ExecutionModelKernel                 ExecutionModel = 6
	ExecutionModelTaskNV                 ExecutionModel = 5267
	ExecutionModelMeshNV                 ExecutionModel = 5268
	ExecutionModelRayGenerationKHR       ExecutionModel = 5313
	ExecutionModelRayGenerationNV        ExecutionModel = 5313
	ExecutionModelIntersectionKHR        ExecutionModel = 5314
	ExecutionModelIntersectionNV         ExecutionModel = 5314
	ExecutionModelAnyHitKHR              ExecutionModel = 5315
	ExecutionModelAnyHitNV               ExecutionModel = 5315
	ExecutionModelClosestHitKHR          ExecutionModel = 5316
	ExecutionModelClosestHitNV           ExecutionModel = 5316
	ExecutionModelMissKHR                ExecutionModel = 5317
	ExecutionModelMissNV                 ExecutionModel = 5317
	ExecutionModelCallableKHR            ExecutionModel = 5318
	ExecutionModelCallableNV             ExecutionModel = 5318
	ExecutionModelMax                    ExecutionModel = 0x7fffffff
)
