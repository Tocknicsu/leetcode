package main

import (
	"fmt"
)

func mySqrt(x int) int {
	ans := 0
	for (ans+1)*(ans+1) <= x {
		ans = ans + 1
	}
	return ans
}

func main() {
	input := 10
	fmt.Println(mySqrt(input))
}
