package main

import (
	"fmt"
)

func reverse(x []int) {
	for i := 0; i < len(x)/2; i = i + 1 {
		x[i], x[len(x)-1-i] = x[len(x)-1-i], x[i]
	}
}

func plusOne(digits []int) []int {
	reverse(digits)
	digits[0] = digits[0] + 1
	flag := 0

	for i, x := range digits {
		x = x + flag
		flag = x / 10
		digits[i] = x % 10
	}
	if flag != 0 {
		digits = append(digits, 1)
	}
	reverse(digits)
	return digits
}
func main() {
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}
