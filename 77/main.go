package main

import "fmt"

func dfs(n int, k int, now int, nowlen int, res []int, ans *[][]int) {
	if nowlen == k {
		co := make([]int, k)
		copy(co, res)
		*ans = append(*ans, co)
		return
	}
	if now > n {
		return
	}
	res[nowlen] = now
	dfs(n, k, now+1, nowlen+1, res, ans)
	dfs(n, k, now+1, nowlen, res, ans)
}

func combine(n int, k int) [][]int {
	ans := &[][]int{}
	res := make([]int, k)
	dfs(n, k, 1, 0, res, ans)
	return *ans
}

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}
