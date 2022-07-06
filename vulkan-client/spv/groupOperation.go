package spv

type GroupOperation uint32

const (
	GroupOperationReduce                     GroupOperation = 0
	GroupOperationInclusiveScan              GroupOperation = 1
	GroupOperationExclusiveScan              GroupOperation = 2
	GroupOperationClusteredReduce            GroupOperation = 3
	GroupOperationPartitionedReduceNV        GroupOperation = 6
	GroupOperationPartitionedInclusiveScanNV GroupOperation = 7
	GroupOperationPartitionedExclusiveScanNV GroupOperation = 8
	GroupOperationMax                        GroupOperation = 0x7fffffff
)
