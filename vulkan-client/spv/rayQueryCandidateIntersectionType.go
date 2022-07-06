package spv

type RayQueryCandidateIntersectionType uint32

//goland:noinspection GoUnusedConst
const (
	RayQueryCandidateIntersectionTypeRayQueryCandidateIntersectionTriangleKHR RayQueryCandidateIntersectionType = 0
	RayQueryCandidateIntersectionTypeRayQueryCandidateIntersectionAABBKHR     RayQueryCandidateIntersectionType = 1
	RayQueryCandidateIntersectionTypeMax                                      RayQueryCandidateIntersectionType = 0x7fffffff
)
