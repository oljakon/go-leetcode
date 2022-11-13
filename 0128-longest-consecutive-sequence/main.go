package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	longestSeq := 0

	for _, num := range nums {
		numMap[num] = true
	}

	for _, num := range nums {
		if _, ok := numMap[num-1]; ok {
			continue
		}

		curr := num
		for numMap[curr+1] {
			curr++
		}

		longestSeq = max(longestSeq, curr-num+1)
	}

	return longestSeq
}
