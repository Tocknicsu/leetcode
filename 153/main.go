package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	now := len(nums) - 1
	step := len(nums)
	for step != 0 {
		if now-step >= 0 && nums[now] > nums[now-step] {
			now = now - step
		} else {
			step = step >> 1
		}
	}
	return nums[now]
}

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(input))
}
