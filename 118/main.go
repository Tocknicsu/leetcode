func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0 ; i < numRows ; i = i + 1 {
			res[i] = make([]int, i+1)
			res[i][0] = 1
			res[i][i] = 1
			for j := 1 ; j < i ; j = j + 1 {
					res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
	}
	return res
}