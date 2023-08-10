package main

import "fmt"

func main() {

	character := "🧟"
	canSearch := true

	// switch character {

	// case "🦸", "🧞":
	// 	fmt.Println("Es un superheroe")
	// case "🦹":
	// 	fmt.Println("Es un supervillano")
	// case "🧟":
	// 	fmt.Println("Es un zombie (supervillano)")
	// default:
	// 	fmt.Println("Es un personaje normal")
	// }

	switch {
	case !canSearch:
		fmt.Println("No está permitida la búsqueda")
	case character == "🦸" || character == "🧞":
		fmt.Println("Es un superheroe")
	case character == "🦹":
		fmt.Println("Es un supervillano")
	case character == "🧟":
		fmt.Println("Es un zombie (supervillano)")
	default:
		fmt.Println("Es un personaje normal")
	}
}
