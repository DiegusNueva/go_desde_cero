package main

import "fmt"

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func main() {
	product1 := Product[uint]{ID: 1, Description: "Zapatos", Price: 30}
	product2 := Product[string]{ID: "serial-123432441435864", Description: "Zapatos", Price: 30}
	fmt.Println(product2)
	fmt.Println(product1)
}
