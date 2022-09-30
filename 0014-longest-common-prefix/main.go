package main

func getShortestString(strs []string) string {
	minLen := 201
	var shortest string

	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			shortest = str
		}
	}

	return shortest
}

func isCommonPrefix(target string, strs []string) bool {
	targetLen := len(target)
	for _, str := range strs {
		if str[:targetLen] != target {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	var res string

	if len(strs) == 1 {
		return strs[0]
	}

	shortest := getShortestString(strs)

	first := 1
	last := len(shortest)

	for first <= last {
		mid := (first + last) / 2
		target := shortest[:mid]
		if isCommonPrefix(target, strs) {
			first++
		} else {
			last--
		}
	}
	res = shortest[:last]
	return res
}
