package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHelloFromHelper("Wandah")
	fmt.Println(result)
	fmt.Println(helper.Application)
	helper.Contoh()
}
