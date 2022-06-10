package main

import "fmt"

func Hello(namaW string, namaK string) {
	fmt.Println("Hello", namaW, namaK)
}

func main() {
	namaW := "Raung"
	Hello(namaW, "Kawijayan")
	Hello("Kamu", "Siapa?")
}
