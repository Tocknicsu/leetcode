package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	um := uint(m)
	un := uint(n)
	udiff := un - um
	now := uint(1)
	for now < udiff {
		now = now << 1
	}
	return int(um & un & (^(now - 1)))
}

func main() {
	m := 0
	n := 1
	fmt.Println(rangeBitwiseAnd(m, n))
}
