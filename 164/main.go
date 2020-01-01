package main

import (
	"fmt"
	"math"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var min, max int
	min = nums[0]
	max = nums[0]

	for _, x := range nums {
		min = Min(min, x)
		max = Max(max, x)
	}

	bucketGap := int(math.Ceil(float64((max - min)) / float64(len(nums)-1)))
	bucketMin := make([]int, len(nums)-1)
	bucketMax := make([]int, len(nums)-1)

	for i := 0; i < len(nums)-1; i = i + 1 {
		bucketMin[i] = 2147483647
		bucketMax[i] = -2147483648
	}

	for _, x := range nums {
		if x == min || x == max {
			continue
		}
		bucket := (x - min) / bucketGap
		bucketMin[bucket] = Min(bucketMin[bucket], x)
		bucketMax[bucket] = Max(bucketMax[bucket], x)
	}

	res := -2147483648
	pre := min

	for i := 0; i < len(nums)-1; i = i + 1 {
		if bucketMin[i] == 2147483647 {
			continue
		}
		res = Max(res, bucketMin[i]-pre)
		pre = bucketMax[i]
	}

	res = Max(res, max-pre)
	return res
}

func main() {
	input := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(input))
}
