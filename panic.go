package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic:", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Upss Error")
	}

	// cara salah memanggil recover
	// message := recover()
	// fmt.Println("Terjadi Panic:", message)
}

func main() {
	runApp(true)
	fmt.Println("Hello Vicry Fahreza")
}
