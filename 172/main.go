package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	ans := 0
	for n != 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}

func main() {
	input := 3
	fmt.Println(trailingZeroes(input))
}
