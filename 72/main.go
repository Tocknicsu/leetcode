package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	mat := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i = i + 1 {
		mat[i] = make([]int, len(word2)+1)
		mat[i][0] = i
	}

	for i := 0; i <= len(word2); i = i + 1 {
		mat[0][i] = i
	}

	mat[0][0] = 0

	for i := 1; i <= len(word1); i = i + 1 {
		for j := 1; j <= len(word2); j = j + 1 {
			if word1[i-1] == word2[j-1] {
				mat[i][j] = mat[i-1][j-1]
			} else {
				mat[i][j] = Min(Min(mat[i-1][j], mat[i][j-1]), mat[i-1][j-1]) + 1
			}
		}
	}

	return mat[len(word1)][len(word2)]
}

func main() {
	x := "intention"
	y := "execution"
	x = "plasma"
	y = "altruism"
	// x = "horse"
	// y = "ros"
	// x := "int"
	// y := "fat"
	fmt.Println(minDistance(x, y))
}
