package twoSum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{i, j}
		}
		seen[num] = i
	}

	return []int{}
}
