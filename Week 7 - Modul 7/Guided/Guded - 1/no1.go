package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	return (Celcius * 4) / 5
}
func CelciusToFahrenheit(Celcius suhu) suhu {
	return (Celcius * 9 / 5) + 32
}
func CelciusToKelvin(Celcius suhu) suhu {
	return Celcius + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&c)
	fmt.Println()

	reamur := CelciusToReamur(c)
	fahrenheit := CelciusToFahrenheit(c)
	kelvin := CelciusToKelvin(c)

	fmt.Printf("%v celcius = %v reamur\n", c, reamur)
	fmt.Printf("%v celcius = %v fahrenheit\n", c, fahrenheit)
	fmt.Printf("%v celcius = %v kelvin\n", c, kelvin)
}