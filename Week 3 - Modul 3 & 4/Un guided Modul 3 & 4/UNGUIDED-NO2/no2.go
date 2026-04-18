package main

import (
    "fmt"
    "math"
)

func hitungPersegi(sisi int) {
    luas := sisi * sisi
    keliling := 4 * sisi

    fmt.Println("Luas persegi :", luas)
    fmt.Println("Keliling persegi :", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
    luas := panjang * lebar
    keliling := 2 * (panjang + lebar)

    fmt.Println("Luas persegi panjang :", luas)
    fmt.Println("Keliling persegi panjang :", keliling)
}

func hitungLingkaran(jarijari float64) {
    luas := math.Pi * jarijari * jarijari
    keliling := 2 * math.Pi * jarijari

    fmt.Println("Luas lingkaran :", luas)
    fmt.Println("Keliling lingkaran :", keliling)
}

func main() {
    var pilihan int

    fmt.Println("--- PROGRAM BANGUN DATAR ---")
    fmt.Println("1. Hitung luas & keliling persegi")
    fmt.Println("2. Hitung luas & keliling persegi panjang")
    fmt.Println("3. Hitung luas & keliling lingkaran")
    fmt.Print("Pilihan : ")
    fmt.Scan(&pilihan)

    switch pilihan {

    case 1:
        var sisi int
        fmt.Print("\nMasukkan sisi : ")
        fmt.Scan(&sisi)
        hitungPersegi(sisi)

    case 2:
        var panjang, lebar int
        fmt.Print("\nMasukkan panjang : ")
        fmt.Scan(&panjang)
        fmt.Print("Masukkan lebar : ")
        fmt.Scan(&lebar)
        hitungPersegiPanjang(panjang, lebar)

    case 3:
        var r float64
        fmt.Print("\nMasukkan jari-jari : ")
        fmt.Scan(&r)
        hitungLingkaran(r)

    default:
        fmt.Println("Pilihan tidak tersedia")
    }
}