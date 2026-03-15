package main

import "fmt"

func hitungSkor(jumlahSoal *int, totalWaktu *int) {
    var waktu int
    *jumlahSoal = 0
    *totalWaktu = 0

    for i := 0; i < 8; i++ {
        fmt.Scan(&waktu)

        if waktu <= 300 {
            *jumlahSoal++
            *totalWaktu += waktu
        }
    }
}

func main() {
    var nama string
    var pemenang string
    var soal, waktu int
    var maxSoal = -1
    var minWaktu int

    for {
        fmt.Scan(&nama)

        if nama == "Selesai" {
            break
        }

        hitungSkor(&soal, &waktu)

        if soal > maxSoal || (soal == maxSoal && waktu < minWaktu) {
            maxSoal = soal
            minWaktu = waktu
            pemenang = nama
        }
    }

    fmt.Println(pemenang, maxSoal, minWaktu)
}
