package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung"

	// Apabila ingin pointer yang baru tanpa mengubah pointer yang dituju
	// address2 = &Address{"Jakarta", "Jakarta Barat", "Indonesia"}
	*address2 = Address{"Jakarta", "Jakarta Barat", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}
