package main

import "fmt"

func main() {
	const nama string = "Vicry Fahreza"
	const npm = "51420265"

	const (
		firstName = "Vicry"
		lastName  = "Fahreza"
	)

	// error
	// nama = "Mera"
	// npm = "144"

	fmt.Println("Nama Mahasiswa = ", nama)
	fmt.Println("NPM Mahasiswa = ", npm)
}
