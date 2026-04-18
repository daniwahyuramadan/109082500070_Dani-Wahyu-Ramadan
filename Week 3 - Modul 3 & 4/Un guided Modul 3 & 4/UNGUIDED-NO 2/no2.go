package main

import "fmt"

func hitungBiaya(jenis string, masuk int, keluar int) int {
    biaya := 0

    for jam := masuk; jam < keluar; jam++ {

        if jenis == "motor" {

            if jam < 17 {
                biaya += 4000
            } else {
                biaya += 5000
            }

        } else if jenis == "mobil" {

            if jam < 17 {
                biaya += 6000
            } else {
                biaya += 7000
            }

        }
    }

    return biaya
}

func main() {

    var jenis string
    var masuk, keluar int
    total := 0
    no := 1

    fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

    for {

        fmt.Printf("\n*Kendaraan %d\n", no)
        fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
        fmt.Scan(&jenis)

        if jenis == "-" {
            break
        }

        fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
        fmt.Scan(&masuk)

        fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
        fmt.Scan(&keluar)

        biaya := hitungBiaya(jenis, masuk, keluar)

        fmt.Printf("Biaya parkir %s %d: %d\n", jenis, no, biaya)
        fmt.Println("================================")

        total += biaya
        no++
    }

    fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", total)
}
