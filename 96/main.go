package main

import "fmt"

func C(n int, m int) int {
	if n-m < m {
		m = n - m
	}
	re := 1
	for i := 0; i < m; i = i + 1 {
		re = re * (n - i) / (i + 1)
	}
	return re
}

func numTrees(n int) int {
	return C(2*n, n) - C(2*n, n-1)
}
func main() {
	input := 0
	fmt.Println(numTrees(input))
}
