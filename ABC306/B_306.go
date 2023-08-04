package main

import (
	"fmt"
)

func main() {
	a := make([]int, 64)
	var sum uint64
	for i := 0; i < 64; i++ {
		fmt.Scan(&a[i])
		pow := uint64(1<<i) * uint64(a[i])
		sum += pow
	}
	fmt.Println(sum)
}
