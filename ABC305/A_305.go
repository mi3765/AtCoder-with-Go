package main

import "fmt"

func main() {
	var n int
	ans := 0
	fmt.Scan(&n)
	if n%5 >= 3 {
		ans = (n/5 + 1) * 5
	} else {
		ans = (n / 5) * 5
	}
	fmt.Println(ans)
	return
}
