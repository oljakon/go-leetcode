package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	res := 0
	count := make(map[byte]int, len(s))

	l := 0
	maxF := 0
	for r := 0; r < len(s); r++ {
		count[s[r]]++
		maxF = max(maxF, count[s[r]])

		for ((r - l + 1) - maxF) > k {
			count[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
