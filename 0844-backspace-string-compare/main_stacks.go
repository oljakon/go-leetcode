package main

func backspaceCompare(s string, t string) bool {
	var sStack, tStack []byte

	for i := range s {
		if s[i] != '#' {
			sStack = append(sStack, s[i])
		} else {
			if len(sStack) > 0 {
				sStack = sStack[:len(sStack)-1]
			}
		}
	}

	for i := range t {
		if t[i] != '#' {
			tStack = append(tStack, t[i])
		} else {
			if len(tStack) > 0 {
				tStack = tStack[:len(tStack)-1]
			}
		}
	}

	// return reflect.DeepEqual(sStack, tStack)

	if len(sStack) != len(tStack) {
		return false
	}
	for i := range sStack {
		if sStack[i] != tStack[i] {
			return false
		}
	}
	return true
}
