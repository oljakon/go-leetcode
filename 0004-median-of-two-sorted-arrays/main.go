package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2

	if len(a) > len(b) {
		a, b = b, a
	}

	aLen, bLen := len(a), len(b)
	total := aLen + bLen
	half := (total + 1) / 2

	l, r := 0, aLen-1

	for {
		h1 := (l + r) >> 1  //nums1
		h2 := half - h1 - 2 //nums2

		var aLeft, aRight, bLeft, bRight int
		if h1 >= 0 {
			aLeft = a[h1]
		} else {
			aLeft = -math.MaxInt
		}
		if h1+1 < aLen {
			aRight = a[h1+1]
		} else {
			aRight = math.MaxInt
		}
		if h2 >= 0 {
			bLeft = b[h2]
		} else {
			bLeft = -math.MaxInt
		}
		if h2+1 < bLen {
			bRight = b[h2+1]
		} else {
			bRight = math.MaxInt
		}

		if aLeft <= bRight && bLeft <= aRight {
			if total%2 == 0 {
				return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2
			}
			return float64(max(aLeft, bLeft))
		}

		if aLeft > bRight {
			r = h1 - 1
		} else {
			l = h1 + 1
		}
	}
}
