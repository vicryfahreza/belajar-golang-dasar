package main

import "fmt"

func main() {
	type NoKtp string

	var ktpVicry NoKtp = "55555555"
	var contohKtp string = "66666666"
	var contohKtpVicry NoKtp = NoKtp(contohKtp)

	fmt.Println(ktpVicry)
	fmt.Println(contohKtpVicry)

}
