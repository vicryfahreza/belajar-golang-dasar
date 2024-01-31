package main

import "fmt"

func main() {
	// map adalah tipe data yang menampung key value dengan tipe data tanpa batas
	person := map[string]string{
		"Name":    "Vicry",
		"Address": "Jakarta Barat",
		"Age":     "18",
	}
	fmt.Println(person["Name"])
	fmt.Println(person["Address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["judul"] = "Buku Golang"
	book["penerbit"] = "PT Angkasa Sejahtera"
	book["salah"] = "Ups"

	delete(book, "salah")
	fmt.Println(book)

}
