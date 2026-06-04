package main

import (
	"fmt"
)

type Pemain struct {
	NamaDepan    string
	NamaBelakang string
	Gol          int
	Assist       int
}

func main() {
	var n int

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	dataPemain := make([]Pemain, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&dataPemain[i].NamaDepan, &dataPemain[i].NamaBelakang, &dataPemain[i].Gol, &dataPemain[i].Assist)
	}

	for i := 1; i < n; i++ {
		key := dataPemain[i]
		j := i - 1

		for j >= 0 && (dataPemain[j].Gol < key.Gol || (dataPemain[j].Gol == key.Gol && dataPemain[j].Assist < key.Assist)) {
			dataPemain[j+1] = dataPemain[j]
			j--
		}
		dataPemain[j+1] = key
	}

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", dataPemain[i].NamaDepan, dataPemain[i].NamaBelakang, dataPemain[i].Gol, dataPemain[i].Assist)
	}
}
