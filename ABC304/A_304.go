package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	a := make([]int, n)
	small := 1e9 + 1
	stock := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
		fmt.Scan(&a[i])
		if small > float64(a[i]) {
			small = float64(a[i])
			stock = i
		}
	}
	for i := stock; i < n; i++ {
		fmt.Println(s[i])
	}
	for i := 0; i < stock; i++ {
		fmt.Println(s[i])
	}
	return
}
