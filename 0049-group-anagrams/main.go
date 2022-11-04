package main

type CharCount [26]int

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	hash := make(map[CharCount][]string)
	var res [][]string

	for _, w := range strs {
		charCount := CharCount{}
		for i := range w {
			charCount[w[i]-'a']++
		}

		hash[charCount] = append(hash[charCount], w)
	}

	for _, val := range hash {
		res = append(res, val)
	}

	return res
}
