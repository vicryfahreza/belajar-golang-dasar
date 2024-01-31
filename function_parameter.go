package main

import "fmt"

func sayHelloWithParam(name string) {
	fmt.Println("Hello,", name)
}

func sayHelloWithAge(name string, age int) {
	fmt.Println("Hello", name, "umur", age)
}

func main() {
	sayHelloWithParam("Marwan")
	sayHelloWithParam("Skibidi")
	sayHelloWithParam("Narwal")

	sayHelloWithAge("Mijah", 23)
	sayHelloWithAge("Jeki", 24)
}
