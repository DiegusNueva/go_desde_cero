package main

import "fmt"

func main() {
	// Operadores comparación: >, <, ==, !=, >=, <=
	fmt.Println(4 > 6)
	fmt.Println(4 < 6)
	fmt.Println(4 != 4)

	//Operadores Lógicos &&, ||
	var age uint = 33
	fmt.Println("Es adulto ?", age >= 18 && age <= 60)
	fmt.Println("Es niño o anciano ?", age <= 18 || age >= 60)

	// Operador lógico Unitario: !8
	fmt.Println(!(4 == 4))

}
