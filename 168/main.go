package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	res := ""
	for n != 0 {
		n -= 1
		t := n % 26
		n /= 26
		res = string(byte(65+t)) + res
	}
	return res
}

func main() {
	input := 701
	fmt.Println(convertToTitle((input)))
}
