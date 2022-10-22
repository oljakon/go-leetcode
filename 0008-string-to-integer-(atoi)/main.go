package main

import "math"

func myAtoi(s string) int {
	res := 0
	sign := 1
	var isSigned bool

	for _, ch := range s {
		if ch == ' ' && res == 0 && !isSigned {
			continue
		} else if ch == '+' && !isSigned {
			isSigned = true
		} else if ch == '-' && !isSigned {
			sign = -1
			isSigned = true
		} else if ch >= '0' && ch <= '9' {
			res = res*10 + sign*int(ch-'0')
			isSigned = true
		} else {
			break
		}

		if res > math.MaxInt32 {
			return math.MaxInt32
		} else if res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return res
}
