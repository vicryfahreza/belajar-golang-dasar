package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Konsep Pass By Value
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1
	address2.City = "Dadap"

	// address 2 akan menduplikat struct baru
	fmt.Println(address2)

	// pointer agar kita dapat Pass By Reference
	// address3 sebagai pointer menunjuk address1
	var address3 *Address = &address1
	address3.City = "Dadap"
	fmt.Println(address3)
	fmt.Println(address1)

}
