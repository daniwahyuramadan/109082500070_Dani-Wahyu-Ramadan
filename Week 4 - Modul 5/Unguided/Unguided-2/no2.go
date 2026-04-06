package main

import "fmt"

func printBintang(n int, current int) {
	if current > n {
		return
	}

	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	printBintang(n, current+1)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printBintang(n, 1)
}
