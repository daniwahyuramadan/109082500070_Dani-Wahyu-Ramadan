package main

import "fmt"

type mahasiswa struct {
	nama, nim string
	IPK       float64
}

func SelectionSortArray(sMHS *[5]mahasiswa) {
	var idx_min, i, j int
	for i = 0; i < len(sMHS)-1; i++ {
		idx_min = i
		for j = i + 1; j < len(sMHS); j++ {
			if sMHS[j].IPK < sMHS[idx_min].IPK {
				idx_min = j
			}
		}
		if idx_min != i {
			sMHS[i], sMHS[idx_min] = sMHS[idx_min], sMHS[i]
		}
	}
}

func main() {
	var structMHS [5]mahasiswa

	for i := 0; i < len(structMHS); i++ {
		fmt.Printf("--- Masukkan Data Mahasiswa ke-%d ---\n", i)
		fmt.Print("Nama : ")
		fmt.Scan(&structMHS[i].nama)
		fmt.Print("NIM : ")
		fmt.Scan(&structMHS[i].nim)
		fmt.Print("IPK : ")
		fmt.Scan(&structMHS[i].IPK)
	}
	fmt.Println()

	fmt.Printf(" === SEBELUM SORTING === ")
	for i := 0; i < len(structMHS); i++ {
		fmt.Printf("--- Data Mahasiswa ke-%d ---\n", i)
		fmt.Printf("Nama : %s\n", structMHS[i].nama)
		fmt.Printf("NIM : %s\n", structMHS[i].nim)
		fmt.Printf("IPK : %.2f\n", structMHS[i].IPK)
	}
	fmt.Println("===============")
	fmt.Println()

	SelectionSortArray(&structMHS)
	fmt.Println(" === SETELAH SRORTING === ")
	for i := 0; i < len(structMHS); i++ {
		fmt.Printf("--- Data Mahasiswa ke-%d ---\n", i)
		fmt.Printf("Nama : %s\n", structMHS[i].nama)
		fmt.Printf("NIM : %s\n", structMHS[i].nim)
		fmt.Printf("IPK : %.2f\n", structMHS[i].IPK)
	}
	fmt.Println("===============")
	fmt.Println()
}
