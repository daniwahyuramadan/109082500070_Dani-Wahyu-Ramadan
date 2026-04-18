package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku : ")
	fmt.Scan(&b.judul)

	fmt.Print("Masukkan nama penulis : ")
	fmt.Scan(&b.penulis)

	fmt.Print("Masukkan penerbit : ")
	fmt.Scan(&b.penerbit)

	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scan(&b.tahunTerbit)

	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&b.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Printf("Judul Buku : %s\n", b.judul)
	fmt.Printf("Penulis : %s\n", b.penulis)
	fmt.Printf("Penerbit : %s\n", b.penerbit)
	fmt.Printf("Tahun Terbit : %d\n", b.tahunTerbit)
	fmt.Printf("Jumlah Halaman : %d\n", b.jumlahHalaman)
}