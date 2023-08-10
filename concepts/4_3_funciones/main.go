package main

import (
	"fmt"
	"strings"
)

func main() {
	//greet("Diego", "Alonso")
	name := "diego"
	toUpperCase(&name)

	fmt.Println(name)
}

// func greet(firstName, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

func toUpperCase(text *string) {
	*text = strings.ToUpper(*text)
}
