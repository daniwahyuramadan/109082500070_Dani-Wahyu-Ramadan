package main

import (
	"fmt"
)

func main() {
	var hitungSuara [21]int

	var suaraMasuk int = 0
	var suaraSah int = 0

	for {
		var angka int
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}

		suaraMasuk++

		if angka >= 1 && angka <= 20 {
			suaraSah++
			hitungSuara[angka]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitungSuara[i])
		}
	}
}
