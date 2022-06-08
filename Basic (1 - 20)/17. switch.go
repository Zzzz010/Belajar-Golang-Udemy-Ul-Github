package main

import "fmt"

func main() {
	name := "Zzzzzz"
	nama_lengkap := "Raung Kamanijen"

	switch name {
	case "Raul":
		fmt.Println("Hallo Raul")
	case "Jayan":
		fmt.Println("Hallo Jayan")
	default:
		fmt.Println("Anda siapa ya ?")
	}
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama anda panjang ya")
	case false:
		fmt.Println("Nama anda pendek ya")
	}
	length_NL := len(nama_lengkap)
	switch {
	case length_NL > 10:
		fmt.Println("Nama anda terlalu panjang")
	case length_NL > 5:
		fmt.Println("Nama anda mayan panjang")
	default:
		fmt.Println("Nama saya cukup pas bagi saya")
	}

}
