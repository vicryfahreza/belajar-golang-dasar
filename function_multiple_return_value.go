package main

import "fmt"

func getFullName() (string, string, string) {
	return "Vicry", "Nugraha", "Fahreza"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
