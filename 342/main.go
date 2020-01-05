package main

import "fmt"

func isPowerOfFour(num int) bool {
	return num > 0 && (num&(num-1)) == 0 && (num&0x55555555) != 0
}

func main() {
	fmt.Println(isPowerOfFour(4))
}
