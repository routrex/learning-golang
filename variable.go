package main

import "fmt"

// deklarasi multiple variable
var (
	fristname = "Routrex" 
	lastname = "Gerzma"
)


// deklarasi multiple const
const (

	port = 3306
	host = "localhost"
	database = "Mysql"
)


func main() {

	fmt.Println(fristname)
	fmt.Println(lastname)

	fmt.Println(port)
	fmt.Println(host)
	fmt.Println(database)

	role:= "ADMIN"

	fmt.Println(role)

	role = "Gudang"

	fmt.Println(role)

	var cek = "Customer"

	fmt.Println(cek)

	fmt.Println(role)
	// Menulis tipe data dalam variable secara eksplisit
	var addrress string

	addrress = "New York"

	fmt.Println(addrress)

	var age int = 22

	fmt.Println(age)

	// Menulis var tanpa tipe data dalam variable
	var city = "France"

	fmt.Println(city)

	
}
