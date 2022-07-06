package spv

type Scope uint32

//goland:noinspection GoUnusedConst
const (
	ScopeCrossDevice    Scope = 0
	ScopeDevice         Scope = 1
	ScopeWorkgroup      Scope = 2
	ScopeSubgroup       Scope = 3
	ScopeInvocation     Scope = 4
	ScopeQueueFamily    Scope = 5
	ScopeQueueFamilyKHR Scope = 5
	ScopeShaderCallKHR  Scope = 6
	ScopeMax            Scope = 0x7fffffff
)
