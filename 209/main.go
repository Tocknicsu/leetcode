package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	ans := len(nums) + 1
	now := 0
	l := 0
	for r, x := range nums {
		now = now + x
		for now >= s {
			if r-l+1 < ans {
				ans = r - l + 1
			}
			now = now - nums[l]
			l = l + 1
		}
	}
	if ans > len(nums) {
		return 0
	} else {
		return ans
	}
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}
