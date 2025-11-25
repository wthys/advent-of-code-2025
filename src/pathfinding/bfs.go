package pathfinding

type (
	BreadthFirst[T comparable] interface {
		AllPathsTo(end T) [][]T
		AllPathsFunc(end T, complete PathConsumer[T])
		AllPathsFuncVisited(end T, complete PathConsumer[T], visited VisitedFunc[T])
	}

	SimpleBFS[T comparable] struct {
		Start T
		neejberFunc NeejberFunc[T]
	}
)

func NeverVisited[T comparable](_ []T, _ T) bool {
	return false
}

func ConstructBreadthFirst[T comparable](start T, neejbers NeejberFunc[T]) BreadthFirst[T] {
	return SimpleBFS[T]{start, neejbers}
}


func (bfs SimpleBFS[T]) AllPathsTo(end T) [][]T {
	paths := [][]T{}
	bfs.AllPathsFunc(end, func(path []T) {
		paths = append(paths, path)
	})
	return paths
}

func (bfs SimpleBFS[T]) AllPathsFunc(end T, complete PathConsumer[T]) {
	bfs.seek([]T{bfs.Start}, end, complete, NeverVisited)
}

func (bfs SimpleBFS[T]) AllPathsFuncVisited(end T, complete PathConsumer[T], visited VisitedFunc[T]) {
	bfs.seek([]T{bfs.Start}, end, complete, visited)
}

func (bfs SimpleBFS[T]) seek(path []T, end T, complete PathConsumer[T], visited VisitedFunc[T]) {
	last := path[len(path)-1]
	for _, neejber := range bfs.neejberFunc(last) {
		if neejber == end {
			complete(append(path, neejber))
			continue
		}

		if visited(path, neejber) {
			continue
		}

		bfs.seek(append(path, neejber), end, complete, visited)
	}
}