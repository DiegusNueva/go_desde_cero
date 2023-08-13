package main

import (
	"fmt"

	"github.com/DiegusNueva/go_desde_cero/tree/main/POO/encapsulation/course"
)

func main() {
	Go := course.New("Go desde Cero", 12.34, false)
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.PrintClasses()
	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
}
