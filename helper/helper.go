package helper

import "fmt"

// tidak bisa diakses
var version = "1.0.0"
var Application = "Helper"

// fungsi ini tidak bisa diimport karena huruf kecil diawal
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHelloFromHelper(name string) string {
	return "Hello " + name + "From Helper"
}

func Contoh() {
	fmt.Println(sayGoodBye("Haryanto"))
	fmt.Println(version)
}
