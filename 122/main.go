func maxProfit(prices []int) int {
	ans := 0
	i := 0
	for i < len(prices) {
			j := i
			for ; j + 1 < len(prices) && prices[j] <= prices[j+1] ; j = j + 1{
			}
			if prices[j] > prices[i] {
					ans += prices[j] - prices[i]
			}
			i = j + 1
	}
	return ans
}
