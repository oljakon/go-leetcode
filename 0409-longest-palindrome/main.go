package main

func longestPalindrome(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	charCount := make(map[byte]int)

	for i := range s {
		charCount[s[i]]++
	}

	res := 0

	for _, count := range charCount {
		res += (count / 2) * 2
		if (res%2 == 0) && (count%2 == 1) {
			res++
		}
	}

	return res
}
