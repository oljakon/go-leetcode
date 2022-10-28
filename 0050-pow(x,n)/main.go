package main

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if x == 0 {
		return 0
	}

	if n == 1 {
		return x
	}

	var helper func(x float64, n int) float64

	helper = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		if x == 0 {
			return 0
		}

		res := helper(x, n/2)
		res = res * res
		if n%2 == 1 {
			return x * res
		} else {
			return res
		}
	}

	res := helper(x, abs(n))
	if n > 0 {
		return res
	} else {
		return 1 / res
	}
}
