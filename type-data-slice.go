package main

import "fmt"

func main() {
	// definisi 1 untuk membuat slice
	nilai:= []uint8{50, 60, 70, 80, 90, 100}

	fmt.Println(nilai)
	fmt.Println(len(nilai))
	fmt.Println(cap(nilai))

	// tertinggi:= nilai[2:6]
	// tertinggi:= nilai[0:]
	// tertinggi:= nilai[:5]
	tertinggi:= nilai[:]

	fmt.Println(tertinggi)

	fmt.Println("=====================")
	// Deklarasi 2 untuk membuat slice
	var buah = []string{"Apel", "Jeruk", "Melon", "Semangka", "Pisang", "Anggur", "Duren"}
	
	fmt.Println(buah)
	
	var result = buah[1:5]

	fmt.Println(result)

	
}