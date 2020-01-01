package main

import (
	"fmt"
	"strings"
)

func strcmp(s1 string, s2 string) int {
	for len(s1) != 0 && s1[0] == '0' {
		s1 = s1[1:]
	}
	for len(s2) != 0 && s2[0] == '0' {
		s2 = s2[1:]
	}
	if len(s1) != len(s2) {
		if len(s1) < len(s2) {
			return -1
		} else {
			return 1
		}
	}
	for i := 0; i < len(s1); i = i + 1 {
		if s1[i] != s2[i] {
			if s1[i] < s2[i] {
				return -1
			} else {
				return 1
			}
		}
	}
	return 0
}

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")

	for len(s1) < len(s2) {
		s1 = append(s1, "0")
	}
	for len(s2) < len(s1) {
		s2 = append(s2, "0")
	}
	for i := 0; i < len(s1); i = i + 1 {
		now := strcmp(s1[i], s2[i])
		if now != 0 {
			return now
		}
	}
	return 0
}

func main() {
	v1 := "1.0.1"
	v2 := "1"
	fmt.Println(compareVersion(v1, v2))
}
