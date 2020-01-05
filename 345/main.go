package main

import (
	"fmt"
)

func isVowel(x byte) bool {
	s := "aeiouAEIOU"
	for _, y := range s {
		if rune(x) == y {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1
	b := []byte(s)
	for i < j {
		for i < j && !isVowel(b[i]) {
			i = i + 1
		}
		for i < j && !isVowel(b[j]) {
			j = j - 1
		}
		if i < j {
			b[i], b[j] = b[j], b[i]
			i = i + 1
			j = j - 1
		}
	}
	return string(b)
}

func main() {
	x := "leetcode"
	fmt.Println(reverseVowels(x))
}
