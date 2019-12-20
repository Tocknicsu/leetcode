package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i = i - 1 {
		if s[i] == ' ' {
			if ans != 0 {
				return ans
			}
		} else {
			ans = ans + 1
		}
	}
	return ans
}

func main() {
	input := "  "
	fmt.Println(lengthOfLastWord(input))
}
