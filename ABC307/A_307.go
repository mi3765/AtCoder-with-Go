package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 7*n)
	var cnt int
	for i := 0; i < 7*n; i++ {
		fmt.Scan(&a[i])
		cnt += a[i]
		if i != 0 && i%7 == 6 {
			fmt.Printf("%d"+" ", cnt)
			cnt = 0
		}
	}
	return
}
