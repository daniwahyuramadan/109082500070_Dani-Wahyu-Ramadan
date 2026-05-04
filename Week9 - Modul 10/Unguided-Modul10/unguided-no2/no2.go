package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var totalWadah []float64

	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += ikan[j]
		}
		totalWadah = append(totalWadah, total)
	}

	fmt.Println("Total berat tiap wadah:")
	for _, t := range totalWadah {
		fmt.Printf("%.2f ", t)
	}

	sum := 0.0
	for _, t := range totalWadah {
		sum += t
	}

	rata := sum / float64(len(totalWadah))

	fmt.Printf("\nRata-rata per wadah: %.2f\n", rata)
}
