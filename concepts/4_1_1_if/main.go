package main

import "fmt"

func main() {

	if character := "🦹"; character == "🦸" {
		fmt.Println("Es un superheroe")
	} else if character == "🦹" {
		fmt.Println("Es un supervillano")
	} else {
		fmt.Println("Es un personaje normal")
	}

	//fmt.Println(character)
}
