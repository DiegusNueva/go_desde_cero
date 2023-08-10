package main

import "fmt"

func main() {
	// bool, string, numeric

	var a bool = true

	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)

	var b string = "EDteam"

	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)

	/* tipos de datos numéricos

	uint -> números positivos de 0 a un número muy alto: edades.
	int -> números positivos negativos y positivos
	byte -> alias for uint8
	rune -> alias de int32
	float32 -> números decimales (float64) */

	var c uint8 = 33

	fmt.Printf("Tipo: %T, Valor: %v\n", c, c)

	var d byte = 33

	fmt.Printf("Tipo: %T, Valor: %v\n", d, d)

	var e rune = 'a'

	fmt.Printf("Tipo: %T, Valor: %v\n", e, e)

	var f float32 = 3234.434

	fmt.Printf("Tipo: %T, Valor: %v\n", f, f)

	var a2 uint8 = 255
	var b2 uint16 = 2550

	c2 := uint16(a2) + b2

	_ = uint16(a2) + b2

	fmt.Printf("Tipo: %T, Valor: %v\n", c2, c2)
	fmt.Printf("Tipo: %T, Valor: %v\n", a2, a2)

	var a64 string

	fmt.Printf("Tipo: %T, Valor: %q\n", a64, a64)

	var a65 uint

	fmt.Printf("Tipo: %T, Valor: %v\n", a65, a65)

	var a66 bool

	fmt.Printf("Tipo: %T, Valor: %v\n", a66, a66)

}
