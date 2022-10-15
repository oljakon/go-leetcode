package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	res := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int, grid *[][]byte)

	dfs = func(r, c int, grid *[][]byte) {
		if (*grid)[r][c] == '1' {
			(*grid)[r][c] = '9'
			if r >= 1 {
				dfs(r-1, c, grid)
			}
			if r+1 < rows {
				dfs(r+1, c, grid)
			}
			if c >= 1 {
				dfs(r, c-1, grid)
			}
			if c+1 < cols {
				dfs(r, c+1, grid)
			}
		}
	}

	for i, row := range grid {
		for j, col := range row {
			if col == '0' || col == '9' {
				continue
			}
			dfs(i, j, &grid)
			res++
		}
	}

	return res
}
