package main

import "fmt"

type exampler interface {
	x()
}

func main() {
	var e exampler
	fmt.Printf("valor: %v, Tipo: %T\n", e, e)
	e.x()
}
