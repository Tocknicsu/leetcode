package main

import (
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func multiply(num1 string, num2 string) string {
	num1 = Reverse(num1)
	num2 = Reverse(num2)

	ans := ""

	for i, x := range num1 {
		for j, y := range num2 {
			if i+j+1 > len(ans) {
				ans = ans + ""
			}
			ans[i+j] = char((x - '0') * (y - '0'))

		}
	}

	return ans
}

func main() {
	input1 := "123"
	input2 := "456"
	fmt.Println(multiply(input1, input2))
}
