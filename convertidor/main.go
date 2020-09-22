package main

import "fmt"

func main() {
	var f float64

	fmt.Println("Convertidor de grados Fahrenheit a Celcius")
	fmt.Println("Ingrese los °F")
	fmt.Scan(&f)

	output := (f - 32) * 5 / 9
	fmt.Println("grados Celcius: ", output)
}
