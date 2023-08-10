package main

import "fmt"

func main() {
	// Slice: son  apuntadores a un Array, no poseen datos.
	// things := [7]string{"🎢", "🚗", "😂", "📃", "🎢", "🚗", "😂"}
	// cars := things[:5]
	// red := things[3:]

	// red[1] = "🇸🇿"
	// all := things[:]

	// fmt.Println("Things:", things)
	// fmt.Println("Cars:", cars)
	// fmt.Println("Cars:", red)

	// fmt.Println("Cars[0]:", cars[0])
	// fmt.Println(all)

	//len(): # de elementos en el slice
	//cap(): # de elementos del array origen, a partir del índice donde se creo el slice

	// animals := [5]string{"🦍", "🐶", "😺", "🐦", "🐘"}
	// pets := animals[1:3] // "🐶", "😺"

	// pets = append(pets, "🦁", "🦷", "😆")

	// // Array[4] "🐶", "😺", "🐦", "🐘"
	// // new Array[8]{"🐶", "😺", "🦁", "🦷", "😆"}

	// fmt.Println("animals:", animals)

	// pets := []string{"🐶", "😺"}
	// pets := make([]string, 0, 3)
	// pets = append(pets, "🦁", "🦷", "😆", "🐶")

	var pets []string

	fmt.Println("pets:", pets)
	fmt.Println("tamaño pets:", len(pets))
	fmt.Println("capacidad pets:", cap(pets))
	fmt.Println("valor 0:", pets == nil)
}
