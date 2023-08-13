package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	//fmt.Printf("valor: %v, Tipo: %T\n", i, i)

	// v, ok := i.(string)

	// if ok {
	// 	fmt.Println(strings.ToUpper(v))
	// }

	switch v := i.(type) {

	case string:
		fmt.Println(strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Println("valor: %v, Tipo: %T\n", i, i)
	}
}

func main() {
	// var e exampler
	// e.x()
	wrapper(12.34)
	wrapper(false)
	wrapper("Diego")
	wrapper("Álvaro")

}
