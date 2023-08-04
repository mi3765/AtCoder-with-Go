package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1000 {
		fmt.Println(n)
		return
	} else if n < 10000 {
		n = n / 10 * 10
		fmt.Println(n)
		return
	} else if n < 100000 {
		n = n / 100 * 100
		fmt.Println(n)
		return
	} else if n < 1000000 {
		n = n / 1000 * 1000
		fmt.Println(n)
		return
	} else if n < 10000000 {
		n = n / 10000 * 10000
		fmt.Println(n)
		return
	} else if n < 100000000 {
		n = n / 100000 * 100000
		fmt.Println(n)
		return
	} else if n < 1000000000 {
		n = n / 1000000 * 1000000
		fmt.Println(n)
		return
	}
	return
}
