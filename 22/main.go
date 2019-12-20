package main

import (
	"fmt"
)

func dfs(l, r int, now string, ans *[]string) {
	if l == 0 && r == 0 {
		*ans = append(*ans, now)
		return
	}
	if l > 0 {
		dfs(l-1, r, now+"(", ans)
	}
	if r > l {
		dfs(l, r-1, now+")", ans)
	}
	// *ans = append(*ans, "13234")
}

func generateParenthesis(n int) []string {
	ans := []string{}
	dfs(n, n, "", &ans)
	return ans
}

func main() {
	input := 3
	fmt.Println(generateParenthesis(input))
}
