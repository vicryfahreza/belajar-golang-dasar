package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result1 := sumAll(10, 10, 10, 10, 10)
	fmt.Println(result1)

	// dengan slice cukup ribet
	// cara ini kurang efektif
	fmt.Println(sumSlice([]int{10, 20, 30, 40}))
	fmt.Println(sumSlice([]int{10, 20}))

	// kirim saja sebagai argument
	result2 := []int{10, 20, 30, 40, 50, 50}
	fmt.Println(sumAll(result2...))
}
