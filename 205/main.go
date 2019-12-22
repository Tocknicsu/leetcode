package main

import "fmt"

func oneSide(s string, t string) bool {
	m := map[rune]rune{}

	for i, x := range s {
		if y, ok := m[x]; ok {
			if y != rune(t[i]) {
				return false
			}
		} else {
			m[x] = rune(t[i])
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	return oneSide(s, t) && oneSide(t, s)
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
