package spv

type Dim uint32

const (
	Dim1D          Dim = 0
	Dim2D          Dim = 1
	Dim3D          Dim = 2
	DimCube        Dim = 3
	DimRect        Dim = 4
	DimBuffer      Dim = 5
	DimSubpassData Dim = 6
	DimMax         Dim = 0x7fffffff
)
