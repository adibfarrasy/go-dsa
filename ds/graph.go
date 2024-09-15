package ds

type Graph[T comparable] struct {
	adjList    map[T][]T
	isDirected bool
}

func NewGraph[T comparable](isDirected bool) Graph[T] {
	return Graph[T]{
		adjList:    map[T][]T{},
		isDirected: isDirected,
	}
}

func (g *Graph[T]) AddEdge(u, v T) {
	g.adjList[u] = append(g.adjList[u], v)

	if !g.isDirected {
		g.adjList[v] = append(g.adjList[v], u)
	}
}

func DFS[T comparable](graph *Graph[T], start T, visit func(T)) {
	visited := map[T]bool{}

	var dfsHelper func(v T)

	dfsHelper = func(v T) {
		if visited[v] {
			return
		}
		visit(v)
		visited[v] = true

		for _, neighbor := range graph.adjList[v] {
			dfsHelper(neighbor)
		}
	}

	dfsHelper(start)
}

func BFS[T comparable](graph *Graph[T], start T, visit func(T)) {
	visited := map[T]bool{}
	queue := []T{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visit(node)

		for _, neighbor := range graph.adjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
