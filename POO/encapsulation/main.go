package main

import (
	"fmt"
)

func main() {

	Go := &Course{
		"Go desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.ChangePrice(67.01)
	fmt.Println(Go.Price)

	Go.PrintClasses()

	// css := Course{Name: "CSS desde cero", IsFree: true}
	// js := Course{}
	// js.Name = "Curso JS"
	// js.UserIDs = []uint{12, 67}

	// fmt.Println(Go.Name)
	// fmt.Printf("%+v\n", css)
	// fmt.Printf("%+v\n", js)

}
