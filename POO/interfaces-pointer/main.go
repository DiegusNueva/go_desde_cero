package main

import "fmt"

type Storager interface {
	Get() string
	Set() string
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) Get() string {
	return p.name
}

func (p *Person) Set() string {
	// Cambio en la implementación para cumplir con la firma correcta
	return p.name
}

func Exec(s Storager, name string) {
	s.Set() // Llamada al método Set sin argumentos
	fmt.Println(s.Get())
}

func main() {
	p := NewPerson("Alejandro")
	Exec(p, "Álvaro") // Aquí también debes quitar el argumento "Álvaro"
}
