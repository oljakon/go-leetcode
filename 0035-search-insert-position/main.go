package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		val := nums[mid]
		if val == target {
			return mid
		} else if val < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
