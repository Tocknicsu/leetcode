package main

import (
	"fmt"
)

var base []string

func dfs(digits *string, index int, str string, result *[]string) {
	// fmt.Println(str, *digits, len(*digits))
	if index == len(*digits) {
		*result = append(*result, str)
		return
	}
	for _, x := range base[int((*digits)[index])-int('0')] {
		dfs(digits, index+1, str+string(x), result)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	base = []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	res := []string{}
	dfs(&digits, 0, "", &res)
	return res
}

func main() {
	input := "23"
	fmt.Println(letterCombinations(input))
}
