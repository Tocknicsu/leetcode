package main

import "fmt"

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1

	for i := 0; i < r+1; i = i + 1 {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l = l + 1
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r = r - 1
			i = i - 1
		}
	}
}

func main() {
	input := []int{2, 2, 2, 2}
	sortColors(input)
	fmt.Println(input)
}
