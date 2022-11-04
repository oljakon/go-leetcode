package main

func containsDuplicate(nums []int) bool {
	dups := make(map[int]bool)

	for _, num := range nums {
		if _, ok := dups[num]; ok {
			return true
		}
		dups[num] = true
	}

	return false
}
