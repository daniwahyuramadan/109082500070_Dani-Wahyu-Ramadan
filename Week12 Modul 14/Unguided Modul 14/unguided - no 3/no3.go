package main

import (
	"fmt"
)

func main() {
	var input int
	var data []int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		return
	}

	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}

	for i, val := range data {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()

	if len(data) < 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak := data[1] - data[0]
	tetap := true

	for i := 1; i < len(data)-1; i++ {
		if data[i+1]-data[i] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}