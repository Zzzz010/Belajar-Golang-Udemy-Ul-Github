package main

import "fmt"

func main() {
	var nilai16 int16 = 120
	var nilai32 int32 = int32(nilai16)
	var nilai8 int8 = int8(nilai16)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai8)

	var nama = "Raung"
	var e byte = nama[3]
	var Estring string = string(e)

	fmt.Println(nama)
	fmt.Println(Estring)

}
