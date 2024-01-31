package main

import "fmt"

func factorialWithLoop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialWithRecursive(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorialWithRecursive(num-1)
	}
}

func main() {
	fmt.Println(factorialWithLoop(10))
	result := factorialWithRecursive(10)
	fmt.Println(result)
}
