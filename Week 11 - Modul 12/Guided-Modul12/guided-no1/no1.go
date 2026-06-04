package main

import "fmt"

type mahasiswa struct {
	nama string
	nim  int
	IPK  float64
}

type arrMahasiswa [5]mahasiswa

func binSearchStruct(arrData arrMahasiswa, nimDicari int) int {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = len(arrData) - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if nimDicari < arrData[med].nim { //kurang dari median, pencarian ke kiri
			kn = med - 1
		} else if nimDicari > arrData[med].nim { //lebih dari median, pencarian ke kanan
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func seqSearchStruct(arrData arrMahasiswa, nimDicari int) int {
	var found int = -1
	for i := 0; i < len(arrData); i++ {
		if arrData[i].nim == nimDicari {
			found = i
			break
		}
	}
	return found
}

func main() {
	var mahasiswa06 arrMahasiswa
	var nimDicari int

	fmt.Println("--- MASUKKAN DATA NIM MAHASISWA SECARA TERURUT ---")
	for i := 0; i < len(mahasiswa06); i++ {
		fmt.Printf("=== Masukkan data mahasiswa 06 pada indeks ke-%d ===\n", i)
		fmt.Print("Masukkan nama : ")
		fmt.Scan(&mahasiswa06[i].nama)
		fmt.Print("Masukkan NIM : ")
		fmt.Scan(&mahasiswa06[i].nim)
		fmt.Print("Masukkan IPK : ")
		fmt.Scan(&mahasiswa06[i].IPK)
		fmt.Println("-----------------------------------------------------")
	}
	fmt.Println()

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari : ")
	fmt.Scan(&nimDicari)
	fmt.Println()

	var idxHasilCariSeqSearch int
	idxHasilCariSeqSearch = seqSearchStruct(mahasiswa06, nimDicari)

	fmt.Println("== HASIL SEQUENTIAL SEARCH ==")
	if idxHasilCariSeqSearch > -1 {
		fmt.Printf("Data NIM mahasiswa %d ditemukan pada array di indeks ke-%d!\n", nimDicari, idxHasilCariSeqSearch)
		fmt.Println("Berikut Detail Datanya : ")
		fmt.Printf("Nama : %s\n", mahasiswa06[idxHasilCariSeqSearch].nama)
		fmt.Printf("NIM : %d\n", mahasiswa06[idxHasilCariSeqSearch].nim)
		fmt.Printf("IPK : %.2f\n", mahasiswa06[idxHasilCariSeqSearch].IPK)
	} else if idxHasilCariSeqSearch == -1 {
		fmt.Printf("Data NIM mahasiswa %d tidak ditemukan pada array!", nimDicari)
	}
	fmt.Println()

	var idxHasilCariBinSearch int
	idxHasilCariBinSearch = binSearchStruct(mahasiswa06, nimDicari)

	fmt.Println("== HASIL BINARY SEARCH ==")
	if idxHasilCariBinSearch > -1 {
		fmt.Printf("Data NIM mahasiswa %d ditemukan pada array di indeks ke-%d!\n", nimDicari, idxHasilCariBinSearch)
		fmt.Println("Berikut Detail Datanya : ")
		fmt.Printf("Nama : %s\n", mahasiswa06[idxHasilCariBinSearch].nama)
		fmt.Printf("NIM : %d\n", mahasiswa06[idxHasilCariBinSearch].nim)
		fmt.Printf("IPK : %.2f\n", mahasiswa06[idxHasilCariBinSearch].IPK)
	} else if idxHasilCariBinSearch == -1 {
		fmt.Printf("Data NIM mahasiswa %d tidak ditemukan pada array!", nimDicari)
	}
}