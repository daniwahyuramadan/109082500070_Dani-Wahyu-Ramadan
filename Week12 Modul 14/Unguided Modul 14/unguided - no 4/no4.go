package main

import (
	"fmt"
)

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

const nMax = 7919

type DaftarBuku [nMax + 1]Buku

func main() {
	var Pustaka DaftarBuku
	var nPustaka int
	var ratingDicari int

	fmt.Scanln(&nPustaka)

	DaftarkanBuku(&Pustaka, nPustaka)

	CetakTerfavorit(Pustaka, nPustaka)

	UrutBuku(&Pustaka, nPustaka)

	Cetak5Terbaru(Pustaka, nPustaka)

	fmt.Scanln(&ratingDicari)

	CariBuku(Pustaka, nPustaka, ratingDicari)
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scanln(&pustaka[i].id)
		fmt.Scanln(&pustaka[i].judul)
		fmt.Scanln(&pustaka[i].penulis)
		fmt.Scanln(&pustaka[i].penerbit)
		fmt.Scanln(&pustaka[i].eksemplar)
		fmt.Scanln(&pustaka[i].tahun)
		fmt.Scanln(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	idxMax := 1
	for i := 2; i <= n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	b := pustaka[idxMax]
	fmt.Printf("%s, %s, %s, %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < 5 {
		batas = n
	}
	for i := 1; i <= batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low := 1
	high := n
	foundIdx := -1

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			foundIdx = mid
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}
