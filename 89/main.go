package main

import "fmt"

func grayCode(n int) []int {
	res := []int{0}

	for i := 0; i < n; i = i + 1 {
		add := 1 << uint(i)
		for j := len(res) - 1; j >= 0; j = j - 1 {
			res = append(res, res[j]|add)
		}
	}

	return res
}

func main() {
	input := 2
	fmt.Println(grayCode(input))
}
