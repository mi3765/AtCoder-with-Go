package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	var t string
	var flag bool
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			if i != j {
				t = s[i] + s[j]
				flag = true
				l := len(t)
				for k := 0; k < l; k++ {
					if t[k] != t[l-k-1] {
						flag = false
					}
				}
				if flag {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
	return
}
