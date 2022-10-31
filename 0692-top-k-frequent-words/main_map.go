package main

import "sort"

func topKFrequent(words []string, k int) []string {
	count := make(map[string]int)
	var highest []string

	for _, w := range words {
		if _, ok := count[w]; !ok {
			highest = append(highest, w)
		}
		count[w]++
	}

	sort.SliceStable(highest, func(i, j int) bool {
		if count[highest[i]] == count[highest[j]] {
			return highest[i] < highest[j]
		}
		return count[highest[i]] > count[highest[j]]
	})

	return highest[:k]
}
