package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface ke struct yang benar
func (person Person) GetName() string {
	return person.Name
}

// interface dapat digunakan berkali kali untuk struct

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Vicry"}
	SayHello(person)
	SayHello(Animal{"GajahSaurus"})
}
