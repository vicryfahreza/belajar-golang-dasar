package main

import "fmt"

func getFullNameWithNamedReturnValue() (firstName, middleName, lastName string) {
	firstName = "Vicry"
	middleName = "Nugraha"
	lastName = "Fahreza"
	return firstName, middleName, lastName
}

func main() {
	first, middle, last := getFullNameWithNamedReturnValue()
	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)
}
