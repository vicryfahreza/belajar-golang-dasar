package main

import "fmt"

type Employee struct {
	Name, Job string
	Age       int
}

func (employee Employee) sayHello(name string) {
	fmt.Println("Hello, My name is", employee.Name, "As Admin What i can help", name, "?")
}

func main() {
	budi := Employee{"Budi", "Product Manager", 25}
	budi.sayHello("Agus")
}
