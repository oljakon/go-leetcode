package main

import "math"

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 || len(t) < len(t) {
		return ""
	}

	countT := make(map[rune]int)
	window := make(map[rune]int)

	for _, ch := range t {
		countT[rune(ch)]++

	}

	res := []int{-1, -1}
	resLen := math.MaxInt
	have, need := 0, len(countT)
	l := 0

	for r, c := range s {
		window[c]++

		if window[c] == countT[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}
			window[rune(s[l])]--
			if window[rune(s[l])] < countT[rune(s[l])] {
				have--
			}
			l++
		}
	}

	l, r := res[0], res[1]

	if resLen != math.MaxInt {
		return s[l : r+1]
	}

	return ""
}
