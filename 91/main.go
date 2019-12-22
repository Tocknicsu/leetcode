package main

import "fmt"

func isAlpha(s string) int {
	if len(s) == 1 && s[0] == '0' {
		return 0
	}
	if len(s) == 2 && (s[0] == '0' || s[0] > '2' || (s[0] == '2' && s[1] > '6')) {
		return 0
	}
	return 1
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := isAlpha(s[:1])
	tmp := 1
	for i := 1; i < len(s); i = i + 1 {
		nextAns := 0
		if isAlpha(s[i:i+1]) == 1 {
			nextAns = nextAns + ans
		}
		if isAlpha(s[i-1:i+1]) == 1 {
			nextAns = nextAns + tmp
		}
		ans, tmp = nextAns, ans
	}
	return ans
}

func main() {
	input := "226"
	fmt.Println(numDecodings(input))
}
