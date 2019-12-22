package main

import (
	"container/list"
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	dag := make([]int, numCourses)
	edges := make([][]int, numCourses)
	for _, edge := range prerequisites {
		x, y := edge[0], edge[1]
		dag[x] += 1
		edges[y] = append(edges[y], x)
	}
	q := list.New()
	for id, d := range dag {
		if d == 0 {
			q.PushBack(id)
		}
	}
	for q.Len() != 0 {
		now := q.Front().Value.(int)
		q.Remove(q.Front())
		for _, x := range edges[now] {
			dag[x] -= 1
			if dag[x] == 0 {
				q.PushBack(x)
			}
		}
	}
	for _, d := range dag {
		if d != 0 {
			return false
		}
	}
	return true
}

func main() {
	input1 := 2
	input2 := [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(canFinish(input1, input2))
}
