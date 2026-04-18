package main

import "fmt"

func tampil(n int) {
	if n == 1 {
		fmt.Print(1, " ")
		return
	}

	fmt.Print(n, " ")

	tampil(n - 1)

	fmt.Print(n, " ")
}

func main() {
	var N int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&N)

	tampil(N)
}
