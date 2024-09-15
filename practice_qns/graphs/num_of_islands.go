package graphs

func NumIslands(grid [][]byte) int {
	ROWS, COLS := len(grid), len(grid[0])
	res := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= ROWS {
			return
		}
		if j < 0 || j >= COLS {
			return
		}
		if grid[i][j] == '0' || grid[i][j] == 'X' {
			return
		}

		// mark as visited
		grid[i][j] = 'X'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i, row := range grid {
		for j, point := range row {
			if point == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
