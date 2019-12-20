package main

import (
	"fmt"

	"container/list"
)

func Max(x, y int) int {
	if x > y {
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
			ans = Max(ans, heights[l]*(i-l))
		}
		heights[l] = x
		st.PushBack(l)
	}
	return ans
}

func main() {
	input := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(input))
}
