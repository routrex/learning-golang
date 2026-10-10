package main

import "fmt"

// type deklarasi
type Email string

type Passsword string


func main() {

	const (
		email Email = "JhonDoe@gmail.com"
		pass Passsword = "JhonDoe123"
	)

	fmt.Println(email)
	fmt.Println(pass)

	fmt.Println("==========================")
// Konversi tipe data
	var total int16 = 3000

	fmt.Println("Ini tipe data asli :",total)

	var convertion int8 = int8(total)

	fmt.Println("Ini tipe data konversi", convertion)

	fmt.Println("==========================")

	var amount int16 = 3000

	fmt.Println("Ini tipe data asli :",amount)

	var konversi int64 = int64(amount)

	fmt.Println("Ini tipe data konversi", konversi)

	
}
