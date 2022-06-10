package main

import "fmt"

func hello(nama string) string {
	if nama == "Unk" {
		return "Anda siapa?"
	} else {
		return "Hallo " + nama
	}
}

func main() {
	hasil := hello("Raung")
	fmt.Println(hasil)

	fmt.Println(hello("Unk"))

}
