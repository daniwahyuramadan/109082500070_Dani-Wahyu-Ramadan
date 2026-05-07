package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Print("Masukkan jumlah elemen array: ")
    fmt.Scan(&n)

    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Printf("Elemen ke-%d: ", i)
        fmt.Scan(&arr[i])
    }

    fmt.Println("\na. Isi array:")
    fmt.Println(arr)

    fmt.Println("\nb. Elemen dengan indeks ganjil:")
    for i := 1; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()

    fmt.Println("\nc. Elemen dengan indeks genap:")
    for i := 0; i < len(arr); i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()

    var x int
    fmt.Print("\nd. Masukkan nilai x: ")
    fmt.Scan(&x)

    fmt.Println("Elemen pada indeks kelipatan", x, ":")
    for i := 0; i < len(arr); i++ {
        if i%x == 0 {
            fmt.Print(arr[i], " ")
        }
    }
    fmt.Println()

    var idx int
    fmt.Print("\ne. Masukkan indeks yang akan dihapus: ")
    fmt.Scan(&idx)

    arr = append(arr[:idx], arr[idx+1:]...)

    fmt.Println("Array setelah dihapus:")
    fmt.Println(arr)

    total := 0
    for i := 0; i < len(arr); i++ {
        total += arr[i]
    }
    rata := float64(total) / float64(len(arr))
    fmt.Println("\nf. Rata-rata =", rata)

    var jumlah float64
    for i := 0; i < len(arr); i++ {
        selisih := float64(arr[i]) - rata
        jumlah += selisih * selisih
    }
    stdDev := math.Sqrt(jumlah / float64(len(arr)))
    fmt.Println("g. Standar deviasi =", stdDev)

    var cari, frek int
    fmt.Print("\nh. Masukkan bilangan yang dicari frekuensinya: ")
    fmt.Scan(&cari)

    for i := 0; i < len(arr); i++ {
        if arr[i] == cari {
            frek++
        }
    }
    fmt.Println("Frekuensi", cari, "=", frek)
}


