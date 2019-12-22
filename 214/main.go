package main

import "fmt"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Reverse(s string) string {
	t := []rune(s)
	mid := len(s) / 2
	for i := 0; i < mid; i = i + 1 {
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	return string(t)
}

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	l := 0
	z := make([]int, len(s))
	t := Reverse(s)
	ans := 0
	for i := 0; i < len(s); i = i + 1 {
		if z[l]+l < i {
			z[i] = 0
		} else {
			z[i] = Min(z[l]+l-i, z[i-l])
		}
		for i+z[i] < len(s) && s[z[i]] == t[i+z[i]] {
			z[i]++
		}
		if z[i]+i > z[l]+l {
			l = i
		}
		if i+z[i] == len(s) {
			ans = i
			break
		}
	}
	return t[:len(s)-z[ans]] + s
}

func main() {
	input := ""
	fmt.Println(shortestPalindrome(input))
}
