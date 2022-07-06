package spv

type LinkageType uint32

const (
	LinkageTypeExport      LinkageType = 0
	LinkageTypeImport      LinkageType = 1
	LinkageTypeLinkOnceODR LinkageType = 2
	LinkageTypeMax         LinkageType = 0x7fffffff
)
