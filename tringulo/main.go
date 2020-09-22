package main

import "fmt"

func main() {
	var b, h float64

	fmt.Println("CALCULAR EL ÁREA DE UN TRIÁNGULO")
	fmt.Println("Ingrese la base")
	fmt.Scan(&b)
	fmt.Println("Ingrese la altura")
	fmt.Scan(&h)

	output := b * h / 2
	fmt.Println("el area es de: ", output)
}
