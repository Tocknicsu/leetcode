package main

import "fmt"

func countPrimes(n int) int {
	res := 0
	p := make([]bool, n)
	for i := 2; i < n; i = i + 1 {
		if p[i] == false {
			res = res + 1
			for j := i * i; j < n; j = j + i {
				p[j] = true
			}
		}
	}
	return res

}

func main() {
	input := 10
	fmt.Println(countPrimes(input))
}
