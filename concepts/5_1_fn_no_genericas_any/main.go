package main

import "fmt"

func main() {
	PrintList("EDteam", "gophers", "Go desde cero")
	PrintList(1, 2, 3)
}

// Función NO genérica
func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
