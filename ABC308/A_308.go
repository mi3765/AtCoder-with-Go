package main

import "fmt"

func main() {
	a := make([]int, 8)
	var compare int = 0
	for i := 0; i < 8; i++ {
		fmt.Scan(&a[i])
		if a[i] < 100 || a[i] > 675 || a[i] < compare || a[i]%25 != 0 {
			fmt.Println("No")
			return
		}
		compare = a[i]
	}

	fmt.Println("Yes")
	return
}
