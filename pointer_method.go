package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	vicry := Man{"Vicry"}
	vicry.Married()

	fmt.Println(vicry.Name)

}
