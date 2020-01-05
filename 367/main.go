func isPerfectSquare(num int) bool {
	now := 0
	x := 1
	for now < num {
			now += x
			x += 2
	}
	return now == num
}