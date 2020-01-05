package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	x := map[int]int{}
	for _, y := range nums1 {
		x[y] += 1
	}
	ans := []int{}
	for _, y := range nums2 {
		if x[y] > 0 {
			x[y] -= 1
			ans = append(ans, y)
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{1, 2, 2, 3}

	fmt.Println(intersect(nums1, nums2))
}
