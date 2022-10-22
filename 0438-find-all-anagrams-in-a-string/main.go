package main

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	var pCount, sCount [26]int

	for i := range p {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	var res []int
	for l := 0; l < len(s)-len(p)+1; l++ {
		if pCount == sCount {
			res = append(res, l)
		}

		if l+len(p) < len(s) {
			sCount[s[l]-'a']--
			sCount[s[l+len(p)]-'a']++
		}
	}

	return res
}
