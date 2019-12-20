package main

import "fmt"

func Print(matrix [][]int) {
	for i := 0; i < len(matrix); i = i + 1 {
		fmt.Println(matrix[i])
	}
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	x := 1
	y := 1

	for i := 0; i < len(matrix); i = i + 1 {
		if matrix[i][0] == 0 {
			y = 0
		}
	}
	for i := 0; i < len(matrix[0]); i = i + 1 {
		if matrix[0][i] == 0 {
			x = 0
		}
	}
	for i := 1; i < len(matrix); i = i + 1 {
		for j := 1; j < len(matrix[0]); j = j + 1 {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i = i + 1 {
		for j := 1; j < len(matrix[0]); j = j + 1 {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if x == 0 {
		for i := 0; i < len(matrix[0]); i = i + 1 {
			matrix[0][i] = 0
		}
	}
	if y == 0 {
		for i := 0; i < len(matrix); i = i + 1 {
			matrix[i][0] = 0
		}
	}
}

func main() {
	input := [][]int{
		{1, 1, 1},
		{0, 1, 1},
		{1, 1, 1},
	}

	setZeroes(input)
	fmt.Println(input)

}
