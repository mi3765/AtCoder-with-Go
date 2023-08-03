package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b-a == 1 && b != 4 && b != 7 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
