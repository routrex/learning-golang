package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello Golang!")

	// fmt.Println("Ini adalah angka", 1)
	// fmt.Println("Ini adalah angka", 2)
	// fmt.Println("Ini adalah angka", 5.5)

	// fmt.Println("apakah bola itu bulat ?", true)
	// fmt.Println("apakah sepak bola itu sama dengan voli", false)
	// fmt.Println(len("Hello golang!"))
	// fmt.Println("apakah sepak bola itu sama dengan voli"[5])

	// var name = "Hello Golang!"
	// fmt.Println(name)
	// name := "Hello Golang!"
	// fmt.Println(name)
	//  var (
	// 	name = "Jhon Doe"
	// 	gender = "Man"
	// 	age = 30
	// 	address = "New York US"
	//  )

	//  fmt.Println(name)
	//  fmt.Println(gender)
	//  fmt.Println(age)
	//  fmt.Println(address)

	// type NoKtp string
	// const (
	// 	angka1 = 12345
	// 	conversi = int32(angka1)
	// 	name = "Rifaldi"
	// 	ktp NoKtp = "2929292929292"
	// )

	// var address = "Tangerang"
	// var konversi = NoKtp(ktp)

	// fmt.Println(konversi)
	
	// fmt.Println(ktp)
	//  const konversiKtp string = NoKtp(ktp)
	//  fmt.Println(konversiKtp)


	// fmt.Println(name)
	// var angka = name[2]
	// fmt.Println(angka)
	// var konversi = string(angka)
	// fmt.Println(konversiKtp)
	// fmt.Println(conversi)


	// var a = 1

	// var b = 2

	// var join = a + b

	// var result = join == 3

	// var name = "Budi"

	// var result = name != "Budi"

	var a = 20

	var b = 25

	var result = a * b

	var ceks1 = result < 100
	var ceks2 = result > 100

	fmt.Println("Hasilnya salah !", ceks1)
	fmt.Println("Hasilnya benar !", ceks2)
}