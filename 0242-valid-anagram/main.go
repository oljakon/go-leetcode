package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[byte]int, len(s))
	countT := make(map[byte]int, len(t))
	for i := 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++
	}

	for val := range countS {
		if _, ok := countT[val]; !ok {
			return false
		} else if countS[val] != countT[val] {
			return false
		}
	}

	return true
}
