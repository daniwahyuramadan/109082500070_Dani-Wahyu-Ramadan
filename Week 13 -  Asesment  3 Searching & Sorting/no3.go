package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int = 0
	var inputNama int

	fmt.Println("Masukkan proses input suara :")
	for {
		fmt.Scan(&inputNama)
		if inputNama == -1 {
			break
		}

		idx := posisi(p, n, inputNama)
		if idx != -1 {
			p[idx].suara++
		} else {
			p[n].nama = inputNama
			p[n].suara = 1
			n++
		}
	}

	for i := 1; i < n; i++ {
		key := p[i]
		j := i - 1

		for j >= 0 && (p[j].suara < key.suara || (p[j].suara == key.suara && p[j].nama > key.nama)) {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}

	fmt.Println("\nHasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
