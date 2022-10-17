package main

func climbStairs(n int) int {
	one, two := 1, 0

	for i := 0; i < n; i++ {
		tmp := one
		one = one + two
		two = tmp
	}

	return one
}
