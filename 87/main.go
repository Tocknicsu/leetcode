package main

import "fmt"

type III struct {
	sz  int
	s1l int
	s2l int
}

func dc(sz int, s1l, s2l int, s1 *string, s2 *string, cache map[III]bool) bool {
	if ret, ok := cache[III{sz, s1l, s2l}]; ok {
		return ret
	}
	if (*s1)[s1l:s1l+sz] == (*s2)[s2l:s2l+sz] {
		cache[III{sz, s1l, s2l}] = true
		return true
	}

	res := false

	for i := 1; i < sz; i = i + 1 {
		if (dc(i, s1l, s2l, s1, s2, cache) && dc(sz-i, s1l+i, s2l+i, s1, s2, cache)) || (dc(i, s1l, s2l+(sz-i), s1, s2, cache) && dc(sz-i, s1l+i, s2l, s1, s2, cache)) {
			res = true
			break
		}
	}
	cache[III{sz, s1l, s2l}] = res
	return res
}

func isScramble(s1 string, s2 string) bool {
	cache := make(map[III]bool)

	return dc(len(s1), 0, 0, &s1, &s2, cache)
}

func main() {
	s1 := "bcbbcccbcbcaaacbb"
	s2 := "acbcabbbbacccbbcc"
	fmt.Println(isScramble(s1, s2))
}
