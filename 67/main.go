package main

import (
	"fmt"
)

func Reverse(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func Minus(x []rune) {
	for i := 0; i < len(x); i = i + 1 {
		x[i] = x[i] - rune('0')
	}
}

func Plus(x []rune) {
	for i := 0; i < len(x); i = i + 1 {
		x[i] = x[i] + rune('0')
	}
}

func addBinary(a string, b string) string {
	ra := Reverse([]rune(a))
	rb := Reverse([]rune(b))

	Minus(ra)
	Minus(rb)

	if len(ra) < len(rb) {
		ra, rb = rb, ra
	}

	flag := rune(0)
	for i := 0; i < len(ra); i = i + 1 {
		ra[i] = ra[i] + flag
		if len(rb) > i {
			ra[i] = ra[i] + rb[i]
		}
		flag = ra[i] / 2
		ra[i] = ra[i] % 2
		if flag == 1 && i == len(ra)-1 {
			ra = append(ra, 0)
		}
	}
	Plus(ra)

	return string(Reverse(ra))
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
