package main

import "fmt"

type Blacklist func(string) bool
type LegalAge func(int) int

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked!", name)
	} else {
		fmt.Println("You are welcome", name)
	}
}

func loginUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Maaf tidak boleh kasar")
	} else {
		fmt.Println("Ok Nama sudah sopan")
	}
}

func dateApp(name string, age int, legalAge LegalAge) {
	if legalAge(age) <= 18 {
		fmt.Println("Maaf masih dibawah umur")
	} else {
		fmt.Println("Iya silahkan menggunakan aplikasi")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("vicry", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	sensorBangsatWord := func(name string) bool {
		return name == "bangsat"
	}

	loginUser("bangsat", sensorBangsatWord)

	dateApp("Toni", 17, func(age int) int {
		return 17
	})

}
