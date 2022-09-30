package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}

	quand := []int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var kSum func(k, start, target int)
	kSum = func(k, start, target int) {
		if k != 2 {
			for i := start; i < len(nums)-k+1; i++ {
				if i > start && nums[i] == nums[i-1] {
					continue
				}
				quand = append(quand, nums[i])
				kSum(k-1, i+1, target-nums[i])
				quand = quand[:len(quand)-1]
			}
			return
		}

		l, r := start, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				res1 := quand
				res1 = append(res1, nums[l])
				res1 = append(res1, nums[r])
				res = append(res, res1)
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	kSum(4, 0, target)

	return res
}
