package main

func isValid(s string) bool {
	parentheses := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	tmp := []byte{}
	for _, ch := range s {
		if _, ok := parentheses[byte(ch)]; ok {
			tmp = append(tmp, byte(ch))
		} else {
			if len(tmp) == 0 {
				return false
			}
			last := tmp[len(tmp)-1]
			if parentheses[last] != byte(ch) {
				return false
			}
			tmp = tmp[:len(tmp)-1]
		}
	}

	if len(tmp) > 0 {
		return false
	}

	return true
}
