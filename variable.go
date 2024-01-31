package main

import "fmt"

func main() {
	var name string

	name = "Vicry Fahreza"
	fmt.Println(name)

	name = "Vicry Cristata"
	fmt.Println(name)

	var address = "Jl Mawar 8"
	age := 22
	fmt.Println(address)
	fmt.Println(age)

	// multiple variable
	var (
		firstName = "Vicry"
		lastName  = "Fahreza"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
