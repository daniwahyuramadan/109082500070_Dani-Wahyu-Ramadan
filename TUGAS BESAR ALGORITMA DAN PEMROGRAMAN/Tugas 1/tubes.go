package main

import (
	"fmt"
)

const MAX int = 100

type Task struct {
	NamaPekerjaan   string
	KategoriRuangan string
	SkalaKesulitan  int
	Durasi          int
	SudahSelesai    bool
}

var DaftarTugas [MAX]Task
var JumlahTugas int = 0

func tambahTugas(nama string, kategori string, kesulitan int, durasi int) {
	if JumlahTugas < MAX {
		DaftarTugas[JumlahTugas].NamaPekerjaan = nama
		DaftarTugas[JumlahTugas].KategoriRuangan = kategori
		DaftarTugas[JumlahTugas].SkalaKesulitan = kesulitan
		DaftarTugas[JumlahTugas].Durasi = durasi
		DaftarTugas[JumlahTugas].SudahSelesai = false
		JumlahTugas++
		fmt.Println("Tugas berhasil ditambahkan!")
	} else {
		fmt.Println("Memori penuh! Tidak bisa menambahkan tugas lagi.")
	}
}

func ubahTugas(indeks int, nama string, kategori string, kesulitan int, durasi int, status bool) {
	if indeks >= 0 && indeks < JumlahTugas {
		DaftarTugas[indeks].NamaPekerjaan = nama
		DaftarTugas[indeks].KategoriRuangan = kategori
		DaftarTugas[indeks].SkalaKesulitan = kesulitan
		DaftarTugas[indeks].Durasi = durasi
		DaftarTugas[indeks].SudahSelesai = status
		fmt.Println("Data tugas berhasil diperbarui!")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusTugas(indeks int) {
	if indeks >= 0 && indeks < JumlahTugas {
		i := indeks
		for i < JumlahTugas-1 {
			DaftarTugas[i] = DaftarTugas[i+1]
			i++
		}
		JumlahTugas--
		fmt.Println("Tugas berhasil dihapus!")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func sequentialSearchNama(target string) int {
	indeks := -1
	i := 0
	for i < JumlahTugas && indeks == -1 {
		if DaftarTugas[i].NamaPekerjaan == target {
			indeks = i
		}
		i++
	}
	return indeks
}

func binarySearchKategori(target string) int {
	urutKategoriInternal()
	low := 0
	high := JumlahTugas - 1
	indeks := -1

	for low <= high && indeks == -1 {
		mid := (low + high) / 2
		if DaftarTugas[mid].KategoriRuangan == target {
			indeks = mid
		} else if DaftarTugas[mid].KategoriRuangan < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return indeks
}

func urutKategoriInternal() {
	for i := 0; i < JumlahTugas-1; i++ {
		minIdx := i
		for j := i + 1; j < JumlahTugas; j++ {
			if DaftarTugas[j].KategoriRuangan < DaftarTugas[minIdx].KategoriRuangan {
				minIdx = j
			}
		}
		temp := DaftarTugas[i]
		DaftarTugas[i] = DaftarTugas[minIdx]
		DaftarTugas[minIdx] = temp
	}
}

func selectionSortKesulitan(urutan string) {
	for i := 0; i < JumlahTugas-1; i++ {
		targetIdx := i
		for j := i + 1; j < JumlahTugas; j++ {
			if urutan == "asc" {
				if DaftarTugas[j].SkalaKesulitan < DaftarTugas[targetIdx].SkalaKesulitan {
					targetIdx = j
				}
			} else if urutan == "desc" {
				if DaftarTugas[j].SkalaKesulitan > DaftarTugas[targetIdx].SkalaKesulitan {
					targetIdx = j
				}
			}
		}
		temp := DaftarTugas[i]
		DaftarTugas[i] = DaftarTugas[targetIdx]
		DaftarTugas[targetIdx] = temp
	}
	fmt.Println("Data berhasil diurutkan dengan Selection Sort!")
}

func insertionSortDurasi(urutan string) {
	for i := 1; i < JumlahTugas; i++ {
		key := DaftarTugas[i]
		j := i - 1

		if urutan == "asc" {
			for j >= 0 && DaftarTugas[j].Durasi > key.Durasi {
				DaftarTugas[j+1] = DaftarTugas[j]
				j--
			}
		} else if urutan == "desc" {
			for j >= 0 && DaftarTugas[j].Durasi < key.Durasi {
				DaftarTugas[j+1] = DaftarTugas[j]
				j--
			}
		}
		DaftarTugas[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan dengan Insertion Sort!")
}

func tampilkanSemuaTugas() {
	if JumlahTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}
	fmt.Println("\n=== DAFTAR TUGAS RUMAH TANGGA ===")
	for i := 0; i < JumlahTugas; i++ {
		status := "Belum Selesai"
		if DaftarTugas[i].SudahSelesai {
			status = "Selesai"
		}
		fmt.Printf("[%d] Nama: %s | Ruang: %s | Kesulitan: %d | Durasi: %d mnt | Status: %s\n",
			i, DaftarTugas[i].NamaPekerjaan, DaftarTugas[i].KategoriRuangan,
			DaftarTugas[i].SkalaKesulitan, DaftarTugas[i].Durasi, status)
	}
}

func tampilkanStatistik() {
	tugasSelesai := 0
	totalWaktu := 0

	for i := 0; i < JumlahTugas; i++ {
		if DaftarTugas[i].SudahSelesai {
			tugasSelesai++
			totalWaktu += DaftarTugas[i].Durasi
		}
	}

	fmt.Println("\n=== STATISTIK TASKMATE ===")
	fmt.Printf("Jumlah tugas yang sudah selesai      : %d\n", tugasSelesai)

	if tugasSelesai > 0 {
		rataRata := float64(totalWaktu) / float64(tugasSelesai)
		fmt.Printf("Rata-rata waktu tugas yang selesai : %.2f menit\n", rataRata)
	} else {
		fmt.Println("Rata-rata waktu tugas yang selesai : 0 menit (Belum ada tugas selesai)")
	}
}

func prosesUbahTugas() {
	tampilkanSemuaTugas()
	if JumlahTugas == 0 {
		return
	}

	var target string
	fmt.Print("Masukkan Nama Pekerjaan yang ingin diubah: ")
	fmt.Scan(&target)

	indeks := sequentialSearchNama(target)
	if indeks == -1 {
		fmt.Println("Tugas dengan nama tersebut tidak ditemukan.")
		return
	}

	fmt.Printf("Tugas ditemukan di indeks ke-%d. Masukkan data baru:\n", indeks)

	var nama, kategori string
	var kesulitan, durasi int
	var selesai bool

	fmt.Print("Nama Pekerjaan Baru : ")
	fmt.Scan(&nama)
	fmt.Print("Kategori Ruang Baru : ")
	fmt.Scan(&kategori)
	fmt.Print("Skala Kesulitan Baru: ")
	fmt.Scan(&kesulitan)
	fmt.Print("Estimasi Waktu Baru : ")
	fmt.Scan(&durasi)
	fmt.Print("Sudah Selesai? (true/false): ")
	fmt.Scan(&selesai)

	ubahTugas(indeks, nama, kategori, kesulitan, durasi, selesai)
}

func prosesHapusTugas() {
	tampilkanSemuaTugas()
	if JumlahTugas == 0 {
		return
	}

	var target string
	fmt.Print("Masukkan Nama Pekerjaan yang ingin dihapus: ")
	fmt.Scan(&target)

	indeks := sequentialSearchNama(target)
	if indeks == -1 {
		fmt.Println("Tugas dengan nama tersebut tidak ditemukan.")
		return
	}

	fmt.Printf("Tugas ditemukan di indeks ke-%d.\n", indeks)
	hapusTugas(indeks)
}

func main() {
	var pilihan int = -1

	for pilihan != 0 {
		fmt.Println("\n==============================")
		fmt.Println("      TASKMATE MANAGEMENT     ")
		fmt.Println("==============================")
		fmt.Println("1. Tambah Tugas Baru")
		fmt.Println("2. Lihat Semua Tugas")
		fmt.Println("3. Ubah Data Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Cari Tugas (Sequential - Nama)")
		fmt.Println("6. Cari Tugas (Binary - Kategori)")
		fmt.Println("7. Urutkan Tugas (Selection - Kesulitan)")
		fmt.Println("8. Urutkan Tugas (Insertion - Durasi)")
		fmt.Println("9. Lihat Statistik Aplikasi")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih menu (0-9): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var nama, kategori string
			var kesulitan, durasi int
			fmt.Print("Masukkan Nama Pekerjaan : ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan Kategori Ruang : ")
			fmt.Scan(&kategori)
			fmt.Print("Masukkan Skala Kesulitan: ")
			fmt.Scan(&kesulitan)
			fmt.Print("Masukkan Estimasi Waktu : ")
			fmt.Scan(&durasi)
			tambahTugas(nama, kategori, kesulitan, durasi)

		} else if pilihan == 2 {
			tampilkanSemuaTugas()

		} else if pilihan == 3 {
			prosesUbahTugas()

		} else if pilihan == 4 {
			prosesHapusTugas()

		} else if pilihan == 5 {
			var target string
			fmt.Print("Masukkan Nama Pekerjaan yang dicari: ")
			fmt.Scan(&target)
			hasil := sequentialSearchNama(target)
			if hasil != -1 {
				fmt.Printf("Ketemu! Tugas ada di indeks ke-%d\n", hasil)
			} else {
				fmt.Println("Tugas tidak ditemukan.")
			}

		} else if pilihan == 6 {
			var target string
			fmt.Print("Masukkan Kategori Ruangan yang dicari: ")
			fmt.Scan(&target)
			hasil := binarySearchKategori(target)
			if hasil != -1 {
				fmt.Printf("Ketemu! Tugas ada di indeks ke-%d (Data diurutkan berdasarkan Kategori)\n", hasil)
			} else {
				fmt.Println("Kategori tidak ditemukan.")
			}

		} else if pilihan == 7 {
			var urutan string
			fmt.Print("Pilih urutan (asc/desc): ")
			fmt.Scan(&urutan)
			selectionSortKesulitan(urutan)
			tampilkanSemuaTugas()

		} else if pilihan == 8 {
			var urutan string
			fmt.Print("Pilih urutan (asc/desc): ")
			fmt.Scan(&urutan)
			insertionSortDurasi(urutan)
			tampilkanSemuaTugas()

		} else if pilihan == 9 {
			tampilkanStatistik()

		} else if pilihan != 0 {
			fmt.Println("Pilihan tidak tersedia. Silakan ulangi.")
		}
	}
	fmt.Println("Terima kasih telah menggunakan TaskMate!")
}
