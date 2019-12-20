package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	now := 0
	for _, x := range nums {
		now = now + x
		if now > ans {
			ans = now
		}
		if now < 0 {
			now = 0
		}
	}
	return ans
}

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(input))
}
