package main

import (
	"fmt"
)

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i = i + 1 {
		if i&1 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}

func main() {
	input := 5
	fmt.Println(countBits(input))
}
