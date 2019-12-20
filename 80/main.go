package main

import "fmt"

func removeDuplicates(nums []int) int {
	now := 0
	for i, _ := range nums {
		if now >= 2 && nums[i] == nums[now-2] {
			continue
		}
		nums[now] = nums[i]
		now = now + 1
	}
	return now
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
