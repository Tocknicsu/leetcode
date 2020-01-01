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

func maxProfit(prices []int) int {
	if len(prices) == 0 {
			return 0
	}
	ans := 0
	pre := prices[0]
	for _, x := range prices {
			ans = Max(ans, x - pre)
			pre = Min(pre, x)
	}
	return ans
}