package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	x := map[int]int{}
	for _, y := range nums1 {
		x[y] |= 1
	}
	for _, y := range nums2 {
		x[y] |= 2
	}
	ans := []int{}
	for k, y := range x {
		if y == 3 {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
