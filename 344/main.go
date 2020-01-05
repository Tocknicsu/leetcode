package main

import (
	"fmt"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i = i + 1 {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	x := []byte{'H', 'e', 'l', 'l', 'o'}
	reverseString(x)
	fmt.Println(x)
}
