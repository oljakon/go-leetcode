package main

func uniquePaths(m int, n int) int {
	row := make([]int, n) // the lowest row
	row[n-1] = 1

	for i := 0; i < m; i++ {
		for j := n - 2; j >= 0; j-- { // except the right column
			row[j] = row[j] + row[j+1]
		}
	}

	return row[0]
}
