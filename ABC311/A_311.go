package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	var cntA, cntB, cntC int
	for i := range s {
		switch s[i] {
		case 'A':
			cntA++
		case 'B':
			cntB++
		case 'C':
			cntC++
		}
		if cntA*cntB*cntC > 0 {
			fmt.Println(i + 1)
			return
		}
	}
}
