package main

import "fmt"

func getHint(secret string, guess string) string {
	var bulls, cows int
	var bucket [10]int

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			bucket[int(secret[i]-'0')]++
			bucket[int(guess[i]-'0')]--
		}
	}

	k := 0
	for i := range bucket {
		if bucket[i] > 0 {
			k += bucket[i]
		}
	}

	cows = len(secret) - bulls - k

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
