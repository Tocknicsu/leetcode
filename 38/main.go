package main

import (
	"fmt"
)

func countAndSay(n int) string {
	now := "1"
	for i := 1; i < n; i = i + 1 {
		next := ""
		count := 1
		value := int(now[0])
		for _, x := range now[1:] {
			if int(x) == int(value) {
				count = count + 1
			} else {
				next = next + fmt.Sprintf("%d%c", count, value)
				count = 1
				value = int(x)
			}
		}
		next = next + fmt.Sprintf("%d%c", count, value)
		now = next
	}
	return now
}

func main() {
	input := 10
	fmt.Println(countAndSay(input))
}
