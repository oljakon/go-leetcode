package main

func sum(arr []int) int {
	res := 0
	for i := range arr {
		res += arr[i]
	}
	return res
}

func pivotIndex(nums []int) int {
	s := sum(nums)
	leftsum := 0
	for i, num := range nums {
		if leftsum == s-leftsum-num {
			return i
		}
		leftsum += num
	}
	return -1
}
