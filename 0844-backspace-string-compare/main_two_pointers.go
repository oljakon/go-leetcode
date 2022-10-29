package main

func backspaceCompare2(s string, t string) bool {
	ptr1, ptr2 := len(s)-1, len(t)-1
	for ptr1 >= 0 || ptr2 >= 0 {
		for ptr1 >= 0 && s[ptr1] == '#' {
			toDel := 1
			ptr1--
			for ptr1 >= 0 && toDel > 0 {
				if s[ptr1] == '#' {
					toDel++
				} else {
					toDel--
				}
				ptr1--
			}
		}

		for ptr2 >= 0 && t[ptr2] == '#' {
			toDel := 1
			ptr2--
			for ptr2 >= 0 && toDel > 0 {
				if t[ptr2] == '#' {
					toDel++
				} else {
					toDel--
				}
				ptr2--
			}
		}

		if (ptr1 >= 0 && ptr2 < 0) ||
			(ptr2 >= 0 && ptr1 < 0) ||
			(ptr1 >= 0 && ptr2 >= 0 && s[ptr1] != t[ptr2]) {
			return false
		}
		ptr1--
		ptr2--
	}

	return ptr1 <= 0 && ptr2 <= 0
}
