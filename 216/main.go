package main

import "fmt"

func dfs(k int, n int, total int, next int, now *[]int, ans *[][]int) {
	if len(*now) == k {
		if total == n {
			res := make([]int, k)
			copy(res, *now)
			*ans = append(*ans, res)
		}
		return
	}
	for i := next; i <= 9; i = i + 1 {
		if total+i <= n {
			*now = append(*now, i)
			dfs(k, n, total+i, i+1, now, ans)
			*now = (*now)[:len(*now)-1]
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	now := []int{}
	dfs(k, n, 0, 1, &now, &ans)
	return ans
}

func main() {
	k := 3
	n := 9
	fmt.Println(combinationSum3(k, n))
}
