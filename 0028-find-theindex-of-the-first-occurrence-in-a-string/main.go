package main

func removeElement(nums []int, val int) int {
	l := 0

	for i := range nums {
		if nums[i] != val {
			nums[l] = nums[i]
			l++
		}
	}

	return l
}
