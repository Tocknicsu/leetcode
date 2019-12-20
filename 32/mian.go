package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	ans := 0
	left := 0
	for _, x := range s {
		if x == '(' {
			left = left + 1
		} else {
			if left > 0 {
				ans = ans + 1
				left = left - 1
			}
		}
	}
	return ans * 2
}

func main() {
	input := ")()())"
	fmt.Println(longestValidParentheses(input))
}
