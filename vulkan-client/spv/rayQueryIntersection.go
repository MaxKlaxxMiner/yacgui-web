package spv

type RayQueryIntersection uint32

const (
	RayQueryIntersectionRayQueryCandidateIntersectionKHR RayQueryIntersection = 0
	RayQueryIntersectionRayQueryCommittedIntersectionKHR RayQueryIntersection = 1
	RayQueryIntersectionMax                              RayQueryIntersection = 0x7fffffff
)
