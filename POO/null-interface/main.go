package main

import "fmt"

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("valor: %v, Tipo: %T\n", i, i)

}

func main() {
	// var e exampler
	// e.x()
	wrapper(12.34)
	wrapper(false)

}
