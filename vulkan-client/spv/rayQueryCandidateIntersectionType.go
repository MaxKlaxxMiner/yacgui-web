package spv

type RayQueryCandidateIntersectionType uint32

const (
	RayQueryCandidateIntersectionTypeRayQueryCandidateIntersectionTriangleKHR RayQueryCandidateIntersectionType = 0
	RayQueryCandidateIntersectionTypeRayQueryCandidateIntersectionAABBKHR     RayQueryCandidateIntersectionType = 1
	RayQueryCandidateIntersectionTypeMax                                      RayQueryCandidateIntersectionType = 0x7fffffff
)
