package main

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	inc := 2 * (numRows - 1)
	// var res string

	var res bytes.Buffer

	for row := 0; row < numRows; row++ {
		for i := row; i < n; i += inc {
			res.WriteByte(s[i])
			// res += string(s[i])
			if row != 0 && row != (numRows-1) && i+inc-2*row < n {
				// res += string(s[i + inc - 2 * row])
				res.WriteByte(s[i+inc-2*row])
			}
		}
	}

	return res.String()
}
