package main

import (
	"strconv"
	"strings"
)

func isDigit(s string) bool {
	var _, e = strconv.Atoi(s)
	if e != nil {
		return false
	}

	return true
}

func decodeString(s string) string {
	var stack []string

	for _, ch := range s {
		if string(ch) != "]" {
			stack = append(stack, string(ch))
		} else {
			substr := ""
			for stack[len(stack)-1] != "[" {
				substr = stack[len(stack)-1] + substr
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]

			var k string
			for len(stack) > 0 && isDigit(stack[len(stack)-1]) {
				k = stack[len(stack)-1] + k
				stack = stack[:len(stack)-1]
			}

			if k == "" {
				stack = append(stack, substr)
			} else {
				t, _ := strconv.Atoi(k)
				for t > 0 {
					stack = append(stack, substr)
					t--
				}
			}
		}
	}

	return strings.Join(stack, "")
}
