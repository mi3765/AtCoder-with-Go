package main

import (
	"fmt"
	"math"
)

func main() {
	var p, q string
	fmt.Scan(&p, &q)
	m := map[string]int{"A": 0, "B": 3, "C": 4, "D": 8, "E": 9, "F": 14, "G": 23}
	c := m[q] - m[p]
	fmt.Println(math.Abs(float64(c)))
	return
}
