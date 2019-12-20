package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	for i := 0; i < m; i = i + 1 {
		for j := 0; j < n; j = j + 1 {
			cost := 0
			if i != 0 {
				if cost == 0 {
					cost = grid[i-1][j]
				}
			}
			if j != 0 {
				if cost == 0 {
					cost = grid[i][j-1]
				} else {
					cost = min(cost, grid[i][j-1])
				}
			}
			grid[i][j] += cost
		}
	}

	return grid[m-1][n-1]
}

func main() {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(input))
}
