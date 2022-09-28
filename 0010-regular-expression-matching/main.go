package main

type coord struct {
	i int
	j int
}

// i - s index
// j - p index
func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	memo := make(map[coord]bool)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if !memo[coord{i, j}] {
			if i >= sLen && j >= pLen {
				return true
			} else if j >= pLen {
				return false
			}

			match := i < sLen && (s[i] == p[j] || string(p[j]) == ".")

			if (j+1) < pLen && string(p[j+1]) == "*" {
				memo[coord{i, j}] = dp(i, j+2) || (match && dp(i+1, j))
			} else {
				memo[coord{i, j}] = match && dp(i+1, j+1)
			}
		}
		return memo[coord{i, j}]
	}
	return dp(0, 0)
}
