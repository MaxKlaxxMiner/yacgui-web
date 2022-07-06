package spv

type AddressingModel uint32

const (
	AddressingModelLogical                    AddressingModel = 0
	AddressingModelPhysical32                 AddressingModel = 1
	AddressingModelPhysical64                 AddressingModel = 2
	AddressingModelPhysicalStorageBuffer64    AddressingModel = 5348
	AddressingModelPhysicalStorageBuffer64EXT AddressingModel = 5348
	AddressingModelMax                        AddressingModel = 0x7fffffff
)
