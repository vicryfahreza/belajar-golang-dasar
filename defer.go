package main

import "fmt"

func loogin() {
	fmt.Println("Sedang menjalankan Aplikasi")
}

func runApplication() {
	defer loogin()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
