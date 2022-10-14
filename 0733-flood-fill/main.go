package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, columns := len(image), len(image[0])
	oldColor := image[sr][sc]
	if color == oldColor {
		return image
	}

	var dfs func(r, c int)

	dfs = func(r, c int) {
		if image[r][c] == oldColor {
			image[r][c] = color
			if r >= 1 {
				dfs(r-1, c)
			}
			if r+1 < rows {
				dfs(r+1, c)
			}
			if c >= 1 {
				dfs(r, c-1)
			}
			if c+1 < columns {
				dfs(r, c+1)
			}
		}
	}

	dfs(sr, sc)

	return image
}
