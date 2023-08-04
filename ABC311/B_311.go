package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	s := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	ans := 0
	comp := 0

	for j := 0; j < d; j++ {
		var flag int
		flag = 1
		for i := 0; i < n; i++ {
			flag && (s[i][j] == 'o')
		}
		if flag {
			comp++
		} else {
			comp = 0
		}
		if ans <= comp {
			ans = comp
		}
	}
	fmt.Println(ans)
	return
}
