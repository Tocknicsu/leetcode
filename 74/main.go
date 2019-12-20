package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	step := len(matrix)
	x := 0
	for step != 0 {
		if x+step < len(matrix) && matrix[x+step][0] <= target {
			x = x + step
		} else {
			step = step >> 1
		}
	}
	step = len(matrix[0])
	y := 0
	for step != 0 {
		if y+step < len(matrix[0]) && matrix[x][y+step] <= target {
			y = y + step
		} else {
			step = step >> 1
		}
	}
	return matrix[x][y] == target
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 13
	fmt.Println(searchMatrix(matrix, target))
}
