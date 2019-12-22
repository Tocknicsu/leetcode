package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := map[int]bool{}
	for _, x := range nums {
		if _, ok := set[x]; ok {
			return true
		}
		set[x] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(containsDuplicate(nums))
}
