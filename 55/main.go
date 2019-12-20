package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	now := nums[0]
	for i := 0; i <= now && i < len(nums) && now < len(nums); i = i + 1 {
		nextNow := i + nums[i]
		if nextNow > now {
			now = nextNow
		}
	}
	return now >= len(nums)-1
}
func main() {
	input := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(input))
}
