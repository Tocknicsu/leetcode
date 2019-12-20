package main

import "fmt"

func dfs(target int, now int, stack []int, col []int, plus []int, minus []int, ans *int) {
	if target == now {
		*ans = *ans + 1
		return
	}
	for i := 0; i < target; i = i + 1 {
		if col[i] == 0 && plus[now+i] == 0 && minus[now-i+target-1] == 0 {
			col[i] = 1
			plus[now+i] = 1
			minus[now-i+target-1] = 1

			stack[now] = i
			dfs(target, now+1, stack, col, plus, minus, ans)
			col[i] = 0
			plus[now+i] = 0
			minus[now-i+target-1] = 0
		}
	}
}

func totalNQueens(n int) int {
	col := make([]int, n)
	plus := make([]int, n*2+1)
	minus := make([]int, n*2+1)
	stack := make([]int, n)

	ans := 0
	dfs(n, 0, stack, col, plus, minus, &ans)
	return ans
}

func main() {
	input := 4
	fmt.Println(totalNQueens(input))
}
