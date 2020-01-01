package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	now := len(nums) - 1
	step := len(nums)
	for step != 0 {
		if now-step >= 0 && nums[now] >= nums[now-step] {
			now = now - step
		} else {
			step = step >> 1
		}
	}
	return nums[now]
}

func main() {
	input := []int{1, 3, 3}
	fmt.Println(findMin(input))
}
