package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i = i + 1 {
		ans[i] = make([]int, n)
	}
	d := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	k := 0
	now := 1
	end := n * n
	x := 0
	y := 0
	for {
		ans[x][y] = now
		now = now + 1
		if now > end {
			break
		}
		for {
			nx := x + d[k][0]
			ny := y + d[k][1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && ans[nx][ny] == 0 {
				x = nx
				y = ny
				break
			}
			k = (k + 1) % 4
		}

	}
	return ans
}

func main() {
	input := 3
	fmt.Println(generateMatrix(input))
}
