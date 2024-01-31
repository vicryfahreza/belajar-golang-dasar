package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var vicry Customer
	vicry.Name = "Vicry"
	vicry.Address = "Jakarta"
	vicry.Age = 21

	fmt.Println(vicry)

	// cara literal
	yugi := Customer{
		Name:    "Yugi",
		Address: "Japan",
		Age:     18,
	}

	fmt.Println(yugi)
	fmt.Println(Customer{"Joey", "Japan", 19})
}
