package spv

type FunctionParameterAttribute uint32

//goland:noinspection GoUnusedConst
const (
	FunctionParameterAttributeZext        FunctionParameterAttribute = 0
	FunctionParameterAttributeSext        FunctionParameterAttribute = 1
	FunctionParameterAttributeByVal       FunctionParameterAttribute = 2
	FunctionParameterAttributeSret        FunctionParameterAttribute = 3
	FunctionParameterAttributeNoAlias     FunctionParameterAttribute = 4
	FunctionParameterAttributeNoCapture   FunctionParameterAttribute = 5
	FunctionParameterAttributeNoWrite     FunctionParameterAttribute = 6
	FunctionParameterAttributeNoReadWrite FunctionParameterAttribute = 7
	FunctionParameterAttributeMax         FunctionParameterAttribute = 0x7fffffff
)
