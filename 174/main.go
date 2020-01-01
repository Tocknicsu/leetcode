package main

import (
	"fmt"
)

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func check(m [][]int, k int64) bool {
	if k+int64(m[0][0]) <= 0 {
		return false
	}
	minValue := (int64(-2147483648) - 2) * int64(len(m)) * int64(len(m[0]))
	dp := make([][]int64, len(m))
	for i := range dp {
		dp[i] = make([]int64, len(m[0]))
	}
	dp[0][0] = k + int64(m[0][0])
	for i := range dp {
		for j := range dp[0] {
			if i == 0 && j == 0 {
				continue
			}
			dp[i][j] = minValue
			now := int64(m[i][j])
			if i > 0 && dp[i-1][j]+now > 0 {
				dp[i][j] = Max(dp[i][j], dp[i-1][j]+now)
			}
			if j > 0 && dp[i][j-1]+now > 0 {
				dp[i][j] = Max(dp[i][j], dp[i][j-1]+now)
			}
		}
	}
	return dp[len(m)-1][len(m[0])-1] > 0
}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	ans := int64(0)
	step := int64(1 << 31)
	for step != 0 {
		res := check(dungeon, ans+step)
		if res {
			step >>= 1
		} else {
			ans += step
		}
	}
	return int(ans + 1)
}

func main() {
	input := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(input))
}
