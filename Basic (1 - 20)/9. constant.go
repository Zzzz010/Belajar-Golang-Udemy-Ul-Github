package main

import "fmt"

func main() {
	const myname string = "Raul"
	const hername = "Rosa"

	// Error
	//myname = "Jayan"
	//hername = "Zerochan"

	const (
		mycat  = "Blacky"
		hercat = "Sudimoro"
	)

	//fmt.Println(mycat)
	fmt.Println(hercat)
}
