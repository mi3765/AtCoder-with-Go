package main

import "fmt"

func rotate(var n int) {
	for i := 0; i < n; i++ {

	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	rotate(n)
	return
}