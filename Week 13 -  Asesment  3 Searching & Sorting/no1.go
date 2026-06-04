package main

import "fmt"

const NMAX = 100

type arrInt [NMAX]int

func SelectionSort(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

func median(T arrInt, n int) float64 {
	if n%2 != 0 {
		return float64(T[n/2])
	}
	return float64(T[n/2-1]+T[n/2]) / 2.0
}

func main() {
	var A arrInt
	var x int
	var n int = 0

	fmt.Println("Input data masukan :")
	fmt.Scan(&x)

	for x != -5313541 {
		if x == 0 {
			SelectionSort(&A, n)
			fmt.Println("Median :")
			fmt.Println(median(A, n))
		} else {
			if n < NMAX {
				A[n] = x
				n++
			}
		}
		fmt.Scan(&x)
	}
}
