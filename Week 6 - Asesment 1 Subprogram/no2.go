package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var maxHari int
	if tujuan == "domestik" {
		maxHari = 3
	} else if tujuan == "mancanegara" {
		maxHari = 8
	}

	if jumlahHari > maxHari {
		return maxHari
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	biayaSatuan := 70000 + 250000 + 300000
	return biayaSatuan * jumlahMhs
}

func perhitunganBiaya(jumlahMhs int, lamaPerjalanan int, tujuan string, totalBiaya *float64) {

	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaDomestik := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDomestik)
	} else if tujuan == "mancanegara" {
		
		*totalBiaya = float64(hariDitanggung) * (float64(biayaDomestik) * 1.5)
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}