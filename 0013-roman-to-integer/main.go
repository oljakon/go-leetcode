package main

func romanToInt(s string) int {
	res := 0

	roman := make(map[string]int, 7)
	roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := range s {
		if i+1 < len(s) {
			if roman[string(s[i+1])] > roman[string(s[i])] {
				res -= roman[string(s[i])]
			} else {
				res += roman[string(s[i])]
			}
		} else {
			res += roman[string(s[i])]
		}
	}

	return res
}
