package main

import (
	"fmt"
	"sort"
)

type PII [][]int

func (list PII) Len() int           { return len(list) }
func (list PII) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }
func (list PII) Less(i, j int) bool { return list[i][0] < list[j][0] }

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(PII(intervals))
	ans := [][]int{}
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i = i + 1 {
		if intervals[i][0] > right {
			ans = append(ans, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
		if intervals[i][1] > right {
			right = intervals[i][1]
		}
	}
	ans = append(ans, []int{left, right})

	return ans
}
func main() {
	input := [][]int{
		{-1, 1},
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(input))
}
