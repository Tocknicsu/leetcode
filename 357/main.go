func Min(x, y int) int {
	if x < y {
			return x
	}
	return y
}

func count(n int) int {
	if n == 0 {
			return 1
	}
	res := 9
	for i := 2 ; i <= n ; i = i + 1 {
			res *= 11 - i
	}
	return res
}

func countNumbersWithUniqueDigits(n int) int {
	res := 0
	for i := 0 ; i <= Min(10, n) ; i = i + 1 {
			res += count(i)
	}
	return res
}