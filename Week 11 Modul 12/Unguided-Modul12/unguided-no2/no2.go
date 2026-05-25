package main

import (
	"fmt"
)

func main() {
	var hitungSuara [21]int
	var suaraMasuk int = 0
	var suaraSah int = 0

	// 1. Membaca dan menghitung input suara
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

	// 2. Mencari Ketua RT (Suara Terbanyak Pertama)
	idKetua := -1
	maksSuaraKetua := -1

	for i := 1; i <= 20; i++ {
		// Menggunakan '>' memastikan jika ada suara yang sama,
		// nomor peserta yang lebih kecil (yang diproses duluan) yang terpilih.
		if hitungSuara[i] > maksSuaraKetua {
			maksSuaraKetua = hitungSuara[i]
			idKetua = i
		}
	}

	// 3. Mencari Wakil Ketua RT (Suara Terbanyak Kedua)
	idWakil := -1
	maksSuaraWakil := -1

	for i := 1; i <= 20; i++ {
		// Jangan hitung calon yang sudah jadi Ketua RT
		if i == idKetua {
			continue
		}

		if hitungSuara[i] > maksSuaraWakil {
			maksSuaraWakil = hitungSuara[i]
			idWakil = i
		}
	}

	// 4. Cetak Hasil Sesuai Format Keluaran
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", idKetua)
	fmt.Printf("Wakil ketua: %d\n", idWakil)
}
