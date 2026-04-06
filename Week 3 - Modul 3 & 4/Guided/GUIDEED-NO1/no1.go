package main

import "fmt"

func factorial(n int, hasil *int) {
    *hasil = 1
    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

func permutation(n, r int, hasil *int) {
    var fn, fnr int

    factorial(n, &fn)
    factorial(n-r, &fnr)

    *hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
    var fn, fr, fnr int

    factorial(n, &fn)
    factorial(r, &fr)
    factorial(n-r, &fnr)

    *hasil = fn / (fr * fnr)
}

func main() {
    var a, b, c, d int
    var p1, c1, p2, c2 int

    fmt.Scan(&a, &b, &c, &d)

    permutation(a, c, &p1)
    combination(a, c, &c1)

    permutation(b, d, &p2)
    combination(b, d, &c2)

    fmt.Println(p1, c1)
    fmt.Println(p2, c2)
}
