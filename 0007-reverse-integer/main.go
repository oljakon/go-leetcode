package main

import "math"

func reverse(x int) int {
	if x >= -9 && x <= 9 {
		return x
	}

	rev := 0

	for x != 0 {
		pop := x % 10
		x /= 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		} else if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}

		rev = rev*10 + pop
	}

	return rev
}
