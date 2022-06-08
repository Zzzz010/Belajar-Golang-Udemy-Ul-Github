package main

import "fmt"

func main() {
	name := "Jayan"

	if name == "Raul" {
		fmt.Println("Halo Raul")
	} else if name == "Jayan" {
		fmt.Println("Bagaimana kabarmu ?")
	} else {
		fmt.Println("Bye - bye")
	}

	if length := len(name); length > 4 {
		fmt.Println("Nama anda terlalu panjang")
	} else {
		fmt.Println("Nama anda tidak panjang")
	}
}
