package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	res := make([]int, 2)
	count := make([]int, 2)

	for _, x := range nums {
		same := false
		for i := 0; i < 2; i = i + 1 {
			if count[i] == 0 {
				same = true
				res[i] = x
				count[i] = 1
				break
			} else if res[i] == x {
				same = true
				count[i] += 1
				break
			}
		}
		if !same {
			count[0] -= 1
			count[1] -= 1
		}
	}

	if count[0] == count[1] {
		if count[0] == 0 {
			return []int{}
		} else {
			return []int{res[0], res[1]}

		}
	} else if count[0] > count[1] {
		return []int{res[0]}
	} else {
		return []int{res[1]}
	}
}

func main() {
	input := []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Println(majorityElement(input))
}
