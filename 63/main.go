package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	for i := 0; i < m; i = i + 1 {
		for j := 0; j < n; j = j + 1 {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				if i != 0 {
					obstacleGrid[i][j] = obstacleGrid[i][j] + obstacleGrid[i-1][j]
				}
				if j != 0 {
					obstacleGrid[i][j] = obstacleGrid[i][j] + obstacleGrid[i][j-1]
				}
				if i == 0 && j == 0 {
					obstacleGrid[i][j] = 1
				}
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}

func main() {
	input := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(input))
}
