package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func main() {
	var data arrayMahasiswa
	var N int
	var searchNIM string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	firstScore := -1
	maxScore := -1
	found := false

	for i := 0; i < N; i++ {
		if data[i].NIM == searchNIM {
			if !found {
				firstScore = data[i].nilai
				maxScore = data[i].nilai
				found = true
			}
			if data[i].nilai > maxScore {
				maxScore = data[i].nilai
			}
		}
	}

	if found {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, firstScore)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, maxScore)
	}
}