package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	res := []string{}

	buttons := make(map[string][]string)
	buttons = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	var backtrack func(i int, currStr string)
	backtrack = func(i int, currStr string) {
		if len(currStr) == len(digits) {
			res = append(res, currStr)
			return
		}

		val := buttons[string(digits[i])]
		for _, c := range val {
			backtrack(i+1, currStr+string(c))
		}
	}

	backtrack(0, "")

	return res
}
