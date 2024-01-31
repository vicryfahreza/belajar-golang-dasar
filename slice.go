package main

import "fmt"

func main() {

	name := [...]string{"Vicry", "Fahreza", "Mizella", "Giovana", "Mercy", "Gerry", "Thomson", "Jack"}

	fmt.Println(name[4:6])
	fmt.Println(name[:4])
	fmt.Println(name[4:])
	fmt.Println(name[:])

	month := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	// index 4 sebanyak 7
	// slice yang terbuat adalah index 4, 5, 6 = May June July
	// pointer 4
	// length 3
	// capacity 8 karena dihitung panjang dair index 4 sampai ujung
	slice := month[4:7]
	fmt.Println(slice)

	arrayOldGod := [...]string{
		"Gro-goroth",
		"Sylvian",
		"God-of-the-depths",
		"Rher",
		"Vinushka",
	}

	slice1 := arrayOldGod[3:]
	// slice1[0] = "Fake God #1"
	// slice1[1] = "Fake God #2"
	// cara diatas akan merubah bukan hanya slice tapi isi dari array utama
	fmt.Println(slice1)

	// append membutuhkan slice dan akan dibuatkan slice baru yang tidak terhubung ke slice yang dituju
	slice2 := append(slice1, "God-Of-Fear-And-Hungers")
	fmt.Println(slice2)
	fmt.Println(arrayOldGod)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Gro-goroth"
	newSlice[1] = "Sylvian"
	// newSlice[2] = "Gro-goroth" error harus menggunakan append
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var newSlice2 = append(newSlice, "Nashr")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// kasus slice yang terhubung
	newSlice2[0] = "Gro-grrr"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice
	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Hati-hati array dengan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
