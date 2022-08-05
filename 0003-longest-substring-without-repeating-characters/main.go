package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	res := 0

	// stores the current index of a character + 1
	charset := make(map[byte]int, n)

	// try to extend the range [i, j]
	for i, j := 0, 0; j < n; j++ {
		if charset[s[j]] > 0 {
			i = max(i, charset[s[j]])
		}
		res = max(res, j-i+1)
		charset[s[j]] = j + 1
	}

	return res
}
