package main

import "fmt"

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int

	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	hasil := pangkat(x, y)
	fmt.Println("Hasil:", hasil)
}
