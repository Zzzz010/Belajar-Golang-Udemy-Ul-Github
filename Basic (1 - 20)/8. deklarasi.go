package main

import "fmt"

func main() {

	type Binatang string
	type Laki bool

	var Darat Binatang = "Macan"
	var Kelamin Laki = true

	fmt.Println(Darat)
	fmt.Println(Kelamin)
}
