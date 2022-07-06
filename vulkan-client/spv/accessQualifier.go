package spv

type AccessQualifier uint32

const (
	AccessQualifierReadOnly  AccessQualifier = 0
	AccessQualifierWriteOnly AccessQualifier = 1
	AccessQualifierReadWrite AccessQualifier = 2
	AccessQualifierMax       AccessQualifier = 0x7fffffff
)
