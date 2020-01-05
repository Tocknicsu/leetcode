package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	x := n / 3
	if n%3 == 0 {
		return int(math.Pow(3, float64(x)))
	} else if n%3 == 2 {
		return 2 * int(math.Pow(3, float64(x)))
	} else {
		return 2 * 2 * int(math.Pow(3, float64(x-1)))
	}
}

func main() {
	fmt.Println(integerBreak(8))
}
