package main

import "fmt"

const pi float64 = 3.1416

func main() {
	var r float64

	fmt.Println("CALCULAR EL ÁREA DE UN CÍRCULO")
	fmt.Println("Ingrese el radio")
	fmt.Scan(&r)

	output := pi * (r * r)
	fmt.Println("el area es de: ", output)
}
