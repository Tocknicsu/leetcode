package main

import "fmt"

func isHappy(n int) bool {
	cache := map[int]bool{}
	now := n
	for now != 1 {
		num := now
		nextNow := 0
		for num != 0 {
			c := num % 10
			nextNow = nextNow + c*c
			num = num / 10
		}
		now = nextNow
		if _, ok := cache[now]; ok {
			return false
		}
		cache[now] = true

	}
	return true
}

func main() {
	input := 3
	fmt.Println(isHappy(input))
}
