package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	l1 := 0
	l2 := 0

	res := []int{}

	for l1 != m && l2 != n {
		if nums1[l1] < nums2[l2] {
			res = append(res, nums1[l1])
			l1 = l1 + 1
		} else {
			res = append(res, nums2[l2])
			l2 = l2 + 1
		}
	}

	for l1 != m {
		res = append(res, nums1[l1])
		l1 = l1 + 1
	}
	for l2 != n {
		res = append(res, nums2[l2])
		l2 = l2 + 1
	}
	copy(nums1, res)
}
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
