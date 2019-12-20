package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	ans := []string{}
	for _, x := range strings.Split(path, "/") {
		if len(x) == 0 || x == "." {
			continue
		}
		if x == ".." {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans = append(ans, x)
		}
	}
	return "/" + strings.Join(ans, "/")
}

func main() {
	input := "/a/../../c7"
	fmt.Println(simplifyPath(input))
}
