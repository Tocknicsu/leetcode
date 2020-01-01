package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	li := []string{}

	now := ""
	for _, x := range s {
		if x != ' ' {
			now = now + string(x)
		} else {
			if len(now) != 0 {
				li = append(li, now)
			}
			now = ""
		}
	}
	if len(now) != 0 {
		li = append(li, now)
	}
	mid := len(li) / 2
	for i := 0; i < mid; i = i + 1 {
		li[i], li[len(li)-1-i] = li[len(li)-1-i], li[i]
	}

	return strings.Join(li, " ")

}

func main() {
	input := "     hello world!     "
	fmt.Println(reverseWords(input))
}
