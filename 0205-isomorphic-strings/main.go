package main

func isIsomorphic(s string, t string) bool {
	sT := make(map[byte]byte)
	tS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := sT[s[i]]; ok && sT[s[i]] != t[i] {
			return false
		} else if _, ok := tS[t[i]]; ok && tS[t[i]] != s[i] {
			return false
		}
		sT[s[i]] = t[i]
		tS[t[i]] = s[i]
	}

	return true
}
