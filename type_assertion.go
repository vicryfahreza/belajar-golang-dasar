package main

import (
	"fmt"
)

func GetOk() any {
	return 12
}

func main() {
	result := GetOk()
	// fmt.Println(result.(string))

	switch value := result.(type) {
	case string:
		fmt.Println("string,", value)
	case int:
		fmt.Println("int,", value)
	default:
		fmt.Println("Unknown")
	}

}
