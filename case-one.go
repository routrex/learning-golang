package main

import "fmt"


func main() {

	var noId [3]uint8

	noId[0] = 001
	noId[1] = 002
	noId[2] = 003

	// var check = noId[0] > noId[2]

	var (
		check1 = noId[2] == noId[2] && noId[2] >= noId[2]
		check2 = noId[0] < noId[1] || noId[0] == noId[1]
	)

	fmt.Println(check1)
	fmt.Println(check2)

	var status = [3]string{
		"pending",
		"waiting",
		"success",
	}

	fmt.Println("Masih",status[0])
	fmt.Println("Dalam",status[1])
	fmt.Println("Sudah",status[2])



	var value = [...]uint8{
		80,
		90,
		100,
	}

	fmt.Println(len(value))


}
