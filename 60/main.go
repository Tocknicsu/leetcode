package main

import (
	"bytes"
	"fmt"
)

func dfs(cnt int, use []bool, fra []int, remain int, buf *bytes.Buffer) {
	if remain == 0 {
		return
	}
	now := 0
	for ; (now+1)*fra[remain-1] <= cnt; now = now + 1 {
	}
	nxtCnt := cnt - now*fra[remain-1]
	for i := 0; i < len(use); i = i + 1 {
		if use[i] {
			continue
		}
		if now != 0 {
			now = now - 1
			continue
		}
		use[i] = true
		buf.WriteRune(rune(int('0') + i + 1))
		dfs(nxtCnt, use, fra, remain-1, buf)
		break
	}
}

func getPermutation(n int, k int) string {
	use := make([]bool, n)
	fra := make([]int, n)
	fra[0] = 1
	k = k - 1
	for i := 1; i < n; i = i + 1 {
		fra[i] = i * fra[i-1]
	}
	buff := bytes.NewBuffer([]byte{})

	dfs(k, use, fra, n, buff)
	return buff.String()
}

func main() {
	n := 3
	k := 6
	fmt.Println(getPermutation(n, k))
}
