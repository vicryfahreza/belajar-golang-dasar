package main

import "fmt"

func main() {
	name := "Raiden"

	switch name {
	case "Evo":
		fmt.Println("Hello Evo")
	case "Oro":
		fmt.Println("Hello Oro")
	default:
		fmt.Println("Boleh Kenalan?")
	}

	// short statement pada switch yang mirip dengan if pada golang
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah pas")
	}

	// apabila ingin membuat statement pada switch case mirip seperti if
	name2 := "vicryfahreza"
	length := len(name2)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah pas")
	}
}
