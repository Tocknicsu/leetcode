package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	flag := 1
	if numerator < 0 {
		flag *= -1
		numerator *= -1
	}
	if denominator < 0 {
		flag *= -1
		numerator *= -1
	}
	t := numerator / denominator
	res := fmt.Sprintf("%d", t)
	numerator %= denominator
	if numerator != 0 {
		res += "."
	}
	cache := map[int]int{}
	res2 := ""
	for numerator != 0 {
		numerator *= 10
		if idx, ok := cache[numerator]; ok {
			res2 = res2[:idx] + "(" + res2[idx:len(res2)] + ")"
			break
		}
		cache[numerator] = len(res2)
		t = numerator / denominator
		numerator %= denominator
		res2 += strconv.Itoa(t)
	}
	res += res2

	if (res[0] != '0' || len(res) > 1) && flag == -1 {
		res = "-" + res
	}

	return res
}

func main() {
	a := -50
	b := 8
	fmt.Println(fractionToDecimal(a, b))
}
