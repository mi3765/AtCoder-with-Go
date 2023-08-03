package main

import "fmt"

func main() {
	var n, p, q int
	var ans int
	fmt.Scan(&n, &p, &q)
	ans = p
	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		d += q
		if ans > d {
			ans = d
		}
	}
	fmt.Println(ans)
	return
}
