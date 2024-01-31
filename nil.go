package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Redi")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println("Data map telah terisi")
		fmt.Println(data["name"])
	}
}
