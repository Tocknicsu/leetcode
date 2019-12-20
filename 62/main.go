package main

import (
	"fmt"
)

func C(m int, n int) int {
	if m-n < n {
		n = m - n
	}
	re := 1
	for i := 1; i <= n; i = i + 1 {
		re = re * (m - i + 1) / i
	}
	return re
}

func uniquePaths(m int, n int) int {
	// mat := make([][]int, m+1)
	// for i := 0; i <= m; i = i + 1 {
	// 	mat[i] = make([]int, n+1)
	// }
	// mat[0][1] = 1

	// for i := 1; i <= m; i = i + 1 {
	// 	for j := 1; j <= n; j = j + 1 {
	// 		mat[i][j] = mat[i-1][j] + mat[i][j-1]
	// 	}
	// }

	return C(m+n-2, m-1)
}

func main() {
	m := 3
	n := 2
	fmt.Println(uniquePaths(m, n))
}
