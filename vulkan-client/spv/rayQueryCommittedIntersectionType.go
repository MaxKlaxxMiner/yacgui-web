package spv

type RayQueryCommittedIntersectionType uint32

const (
	RayQueryCommittedIntersectionTypeRayQueryCommittedIntersectionNoneKHR      RayQueryCommittedIntersectionType = 0
	RayQueryCommittedIntersectionTypeRayQueryCommittedIntersectionTriangleKHR  RayQueryCommittedIntersectionType = 1
	RayQueryCommittedIntersectionTypeRayQueryCommittedIntersectionGeneratedKHR RayQueryCommittedIntersectionType = 2
	RayQueryCommittedIntersectionTypeMax                                       RayQueryCommittedIntersectionType = 0x7fffffff
)
