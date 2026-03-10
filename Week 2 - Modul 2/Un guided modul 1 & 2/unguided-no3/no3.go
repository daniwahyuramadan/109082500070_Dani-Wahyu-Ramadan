package main

import "fmt"

func main() {
	var berat int
	var kg, sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("masukkan total berat (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total = biayaKg + biayaSisa

	fmt.Println("\n=== Detail Perhitungan ===")
	fmt.Printf("detail berat : %d kg + %d gram\n", kg, sisa)
	fmt.Printf("detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("total biaya : Rp. %d\n", total)
}