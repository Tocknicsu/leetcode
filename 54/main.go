package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	d := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	row := len(matrix)
	col := len(matrix[0])
	border := []int{col, row, 0, 0}
	borderOp := []int{-1, -1, 1, 1}
	ans := make([]int, row*col)
	now := 0
	x := 0
	y := 0
	k := 0

	for {
		ans[now] = matrix[x][y]
		now = now + 1
		if now == len(ans) {
			break
		}
		for {
			nx := x + d[k][0]
			ny := y + d[k][1]
			if k == 0 || k == 2 {
				if ny < border[0] && ny >= border[2] {
					x = nx
					y = ny
					break
				}
			} else {
				if nx < border[1] && nx >= border[3] {
					x = nx
					y = ny
					break
				}
			}
			border[(k+3)%4] = border[(k+3)%4] + borderOp[(k+3)%4]
			k = (k + 1) % 4
		}
		// for i := 0
	}

	return ans
}

func main() {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(input))
}
