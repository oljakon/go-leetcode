package main

func sortedSquares(nums []int) []int {
	first := 0
	last := len(nums) - 1
	squares := make([]int, len(nums))

	for i := range nums {
		if nums[first]*nums[first] > nums[last]*nums[last] {
			squares[len(squares)-i-1] = nums[first] * nums[first]
			first++
		} else {
			squares[len(squares)-i-1] = nums[last] * nums[last]
			last--
		}
	}

	return squares
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}

	squares := sortedSquares(arr)

	_ = squares
}
