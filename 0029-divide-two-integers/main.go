package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	quotient := 0

	dividendAbs := abs(dividend)
	divisorAbs := abs(divisor)
	for dividendAbs >= divisorAbs {
		sub := divisorAbs
		add := 1
		for dividendAbs >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		dividendAbs -= sub
		quotient += add
	}

	neg := (dividend < 0) != (divisor < 0)
	if neg {
		return -quotient
	}
	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
