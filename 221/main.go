
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	st := list.New()
	ans := 0
	for i, x := range heights {
		l := i
		for st.Len() != 0 && heights[st.Back().Value.(int)] >= x {
			l = st.Back().Value.(int)
			st.Remove(st.Back())
            t := Min(heights[l], (i-l))
			ans = Max(ans, t * t)
		}
		heights[l] = x
		st.PushBack(l)
	}
	return ans
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	count := make([][]int, len(matrix))
	ans := 0
	for i := 0; i < len(matrix); i = i + 1 {
		count[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j = j + 1 {
			if matrix[i][j] == '1' {
				count[i][j] = 1
				if i != 0 {
					count[i][j] += count[i-1][j]
				}
			}
		}
		ans = Max(ans, largestRectangleArea(count[i]))
	}
	return ans
}
