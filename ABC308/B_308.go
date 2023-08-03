package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	c := make([]string, n)
	d := make([]string, m)
	p := make([]int, m+1)
	var ans int
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&d[i])
	}
	for i := 0; i < m+1; i++ {
		fmt.Scan(&p[i])
	}
	for i := 0; i < n; i++ {
		var flag bool
		for j := 0; j < m; j++ {
			if c[i] == d[j] {
				ans += p[j+1]
				flag = true
			}
		}
		if !flag {
			ans += p[0]
		}
	}
	fmt.Println(ans)
	return
}
