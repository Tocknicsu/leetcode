package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	needCnt := 0
	need := make([]int, 256)
	has := make([]int, 256)

	for _, x := range t {
		if need[int(x)] == 0 {
			needCnt = needCnt + 1
		}
		need[int(x)] = need[int(x)] + 1
	}

	now := 0

	ansl := -1
	ansr := len(s) + 1
	anslength := len(s) + 2

	// fmt.Println(ansl, ansr)

	l := 0
	for r, y := range s {
		iy := int(y)
		has[iy] = has[iy] + 1
		if has[iy] == need[iy] {
			now = now + 1
		}
		for now == needCnt {
			il := int(s[l])
			has[il] = has[il] - 1
			if has[il] < need[il] {
				length := r - l + 1
				if length < anslength {
					ansl = l
					ansr = r
					anslength = length
				}
				now = now - 1
			}
			l = l + 1
		}
	}

	if anslength > len(s) {
		return ""
	}

	return s[ansl : ansr+1]
}

func main() {
	s := "aa"
	t := "aa"
	fmt.Println(minWindow(s, t))
}
