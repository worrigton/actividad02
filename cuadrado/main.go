package main

import "fmt"

func main() {
	var lado float64

	fmt.Println("CALCULAR EL ÁREA DE UN CUADRADO")
	fmt.Println("Ingrese el lado")
	fmt.Scan(&lado)

	output := lado * lado
	fmt.Println("el area es de: ", output)

}
