package graphs

func FindRedundantConnection(edges [][]int) []int {
	parent := []int{}
	parent = append(parent, 0) // dummy

	// edge start is 1-indexed
	for i := 1; i < len(edges)+1; i++ {
		parent = append(parent, i)
	}

	for _, edge := range edges {
		if find(parent, edge[0]) == find(parent, edge[1]) {
			return edge
		}

		unify(parent, edge[0], edge[1])
	}

	return []int{}
}

func find(parent []int, num int) int {
	if parent[num] == num {
		return num
	}

	return find(parent, parent[num])
}

func unify(parent []int, x, y int) {
	parent[find(parent, y)] = find(parent, x)
}
