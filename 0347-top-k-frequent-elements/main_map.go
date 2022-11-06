package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	var freq []int

	for _, n := range nums {
		if _, ok := count[n]; !ok {
			freq = append(freq, n)
		}
		count[n]++
	}

	sort.SliceStable(freq, func(i, j int) bool {
		if count[freq[i]] == count[freq[j]] {
			return freq[i] < freq[j]
		}
		return count[freq[i]] > count[freq[j]]
	})

	return freq[:k]
}
