package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	res := 0
	count := 0
	for _, x := range nums {
		if count == 0 {
			res = x

			count = 1
		} else {
			if res == x {
				count += 1
			} else {
				count -= 1
			}
		}
	}
	return res
}
func main() {
	input := []int{6, 5, 5}
	fmt.Println(majorityElement((input)))
}
