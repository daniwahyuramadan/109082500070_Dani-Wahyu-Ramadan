package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	*n = 0

	fmt.Println("Masukkan karakter satu per satu, akhiri dengan titik (.)")

	for *n < NMAX {
		fmt.Scan(&input)

		if input == "." {
			break
		}

		t[*n] = rune(input[0])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	if palindrome(tab, m) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}

	balikanArray(&tab, m)

	fmt.Print("Hasil dibalik: ")
	cetakArray(tab, m)
}
