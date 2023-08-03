package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var cnt int
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if s == "For" {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt > 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
