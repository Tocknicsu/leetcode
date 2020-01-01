package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProduct(nums []int) int {
	ans := -2147483648
	nowMin := 0
	nowMax := 0

	for _, x := range nums {
		if x != 0 {
			if nowMin == 0 {
				nowMin = x
				nowMax = x
			} else {
				nextMin := Min(x, Min(nowMin*x, nowMax*x))
				nextMax := Max(x, Max(nowMin*x, nowMax*x))
				nowMax = nextMax
				nowMin = nextMin
			}
		} else {
			nowMin = 0
			nowMax = 0
		}
		ans = Max(ans, nowMax)
	}
	return ans
}

func main() {
	nums := []int{3, -1, 4}
	fmt.Println(maxProduct(nums))
}
