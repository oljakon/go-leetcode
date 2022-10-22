package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0

	sLen := len(s)

	for i := 0; i < sLen; i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		lenMax := max(len1, len2)

		if lenMax > (end - start) {
			start = i - (lenMax-1)/2
			end = i + lenMax/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
