package main

import (
    "fmt"
    "sort"
)

func main() {
    var n int
    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)

        var ganjil []int
        var genap []int

        for j := 0; j < m; j++ {
            var nomor int
            fmt.Scan(&nomor)

            if nomor%2 != 0 {
                ganjil = append(ganjil, nomor)
            } else {
                genap = append(genap, nomor)
            }
        }

        sort.Ints(ganjil)

        sort.Slice(genap, func(x, y int) bool {
            return genap[x] > genap[y]
        })

        var hasil []int
        hasil = append(hasil, ganjil...)
        hasil = append(hasil, genap...)

        for k, val := range hasil {
            if k > 0 {
                fmt.Print(" ")
            }
            fmt.Print(val)
        }
        fmt.Println()
    }
}
