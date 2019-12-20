package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := [][]int{}
	for i := int(1<<uint(len(nums))) - 1; i >= 0; i = i - 1 {
		res := []int{}
		for j := 0; j < len(nums); j = j + 1 {
			if i&(1<<uint(j)) != 0 {
				res = append(res, nums[j])
			}
		}
		ans = append(ans, res)
	}
	return ans
}

func main() {
	nums := []int{}
	fmt.Println(subsets(nums))
}
