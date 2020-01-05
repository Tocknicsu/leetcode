func Min(x, y int) int {
	if x < y {
			return x
	}
	return y
}

func coinChange(coins []int, amount int) int {
	const max = 1029384756
	dp := make([]int, amount + 1)

	for x := range dp {
			dp[x] = max
	}
	 dp[0] = 0
	for _, x := range coins {
			for i := x ; i <= amount ; i = i + 1 {
							dp[i] = Min(dp[i], dp[i-x] + 1)
			}
	}
	if dp[amount] >= max {
			return -1
	} else {
			return dp[amount]
	}
}
