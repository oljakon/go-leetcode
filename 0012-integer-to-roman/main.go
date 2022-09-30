package main

import "strings"

type Roman struct {
	Symbol string
	Value  int
}

func intToRoman(num int) string {
	var res string

	roman := [13]Roman{{"M", 1000}, {"CM", 900}, {"D", 500}, {"CD", 400},
		{"C", 100}, {"XC", 90}, {"L", 50}, {"XL", 40},
		{"X", 10}, {"IX", 9}, {"V", 5}, {"IV", 4}, {"I", 1}}

	for _, val := range roman {
		if num/val.Value != 0 {
			res += strings.Repeat(val.Symbol, num/val.Value)
			num = num % val.Value
		}
	}

	return res
}
