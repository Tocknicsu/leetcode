func Min(x, y int) int {
	if x < y {
			return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
			return 0
	}

	for i := 1 ; i < len(triangle) ; i = i + 1 {
			for j := range triangle[i] {
					last := 2147483647
					if j != i {
							last = Min(last, triangle[i-1][j])
					}
					if j != 0 {
							last = Min(last, triangle[i-1][j-1])
					}
					triangle[i][j] += last
			}
	}

	ans := triangle[len(triangle)-1][0]

	for _, x := range triangle[len(triangle)-1] {
			ans = Min(ans, x)
	}
	return ans
}