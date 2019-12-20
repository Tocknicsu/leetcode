package main

import "fmt"

func dfs(x int, y int, board [][]byte, word string, now int, d [][]int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	if board[x][y] != word[now] {
		return false
	}
	if now == len(word)-1 {
		return true
	}

	tmp := board[x][y]
	board[x][y] = 0

	for k := 0; k < 4; k = k + 1 {
		nx := x + d[k][0]
		ny := y + d[k][1]
		if dfs(nx, ny, board, word, now+1, d) {
			return true
		}
	}
	board[x][y] = tmp
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(word) == 0 {
		return false
	}
	d := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for i := 0; i < len(board); i = i + 1 {
		for j := 0; j < len(board[i]); j = j + 1 {
			if dfs(i, j, board, word, 0, d) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	board = [][]byte{
		{'a'},
	}
	word := "a"
	fmt.Println(exist(board, word))
}
