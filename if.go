package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Vicry" {
		fmt.Println("Hello Vicry")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Boleh Kenalan?")
	}

	// if pada go dapat ditaruh inisiasinya di baris if
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah pas")
	}
}
