package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(pangkat(n))
}

func pangkat(n int) int {
	if n == 0 {
		return 1
	} else {
		return 2 * pangkat(n-1)
	}
}
