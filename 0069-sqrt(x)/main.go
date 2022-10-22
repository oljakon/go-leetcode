package main

func mySqrt(x int) int {
	l, r := 0, x

	for (l + 1) < r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid
		} else {
			r = mid
		}
	}

	if r*r == x {
		return r
	} else {
		return l
	}
}
