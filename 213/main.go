package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := [2][]int{}
	for i := 0; i < 2; i = i + 1 {
		dp[i] = make([]int, len(nums)+1)
	}
	dp[1][1] = nums[0]
	for i := 2; i <= len(nums); i = i + 1 {
		val := nums[i-1]
		dp[0][i] = max(dp[0][i-2]+val, dp[0][i-1])
		dp[1][i] = max(dp[1][i-2]+val, dp[1][i-1])
	}
	return max(dp[0][len(nums)], dp[1][len(nums)-1])
}

func main() {
	input := []int{3}
	fmt.Println(rob(input))
}
