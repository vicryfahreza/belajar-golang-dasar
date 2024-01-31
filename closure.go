package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		counter++
		fmt.Println("Increment")
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
}
