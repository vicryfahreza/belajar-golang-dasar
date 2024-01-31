package main

import "fmt"

func main() {
	var arrayUmur [2]int
	arrayUmur[0] = 10
	arrayUmur[1] = 23

	arrayNama := [...]string{
		"Vicry",
		"Fahreza",
		"Cristata",
	}

	fmt.Println(arrayNama)
	fmt.Println(len(arrayNama))
	fmt.Println(arrayNama[0])
	fmt.Println(arrayUmur)
}
