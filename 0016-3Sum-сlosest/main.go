package main

import (
	"math"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	minDif := math.MaxInt
	minSum := math.MaxInt

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := a + nums[l] + nums[r]
			if threeSum > target {
				r--
			} else {
				l++
			}

			dif := abs(threeSum - target)
			if dif < minDif {
				minDif = dif
				minSum = threeSum
			}
		}
	}

	return minSum
}
