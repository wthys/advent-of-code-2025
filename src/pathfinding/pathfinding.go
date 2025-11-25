package pathfinding

type (
	PathConsumer[T comparable]   func(path []T)
	NeejberFunc[T comparable]    func(node T) []T
	ExitFunc[T comparable]       func(node T) bool
	EdgeWeightFunc[T comparable] func(in T, out T) int
	VisitedFunc[T comparable]    func(path []T, node T) bool
)