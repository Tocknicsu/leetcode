package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i, x := range nums {
		var pre, next int
		if i == 0 {
			pre = -2147483648
		} else {
			pre = nums[i-1]
		}
		if i == len(nums)-1 {
			next = -2147483648
		} else {
			next = nums[i+1]
		}
		if pre < x && x > next {
			return i
		}
	}
	return -1
}
func main() {
	input := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(input))
}
