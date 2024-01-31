package main

import "fmt"

func main() {
	result := sayHelloWithReturn("MellowDie")
	fmt.Println(result)
}

func sayHelloWithReturn(name string) string {
	return "Hello" + name
}
