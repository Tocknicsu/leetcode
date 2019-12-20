package main

import "fmt"

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	x := 1
	y := 2
	for i := 3; i <= n; i = i + 1 {
		y, x = x+y, y
	}
	return y
}

func main() {
	input := 4
	fmt.Println(climbStairs(input))
}
