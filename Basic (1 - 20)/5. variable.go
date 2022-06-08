package main

import "fmt"

func main() {

	// Shortcut No 2

	var (
		myname  = "Raung Kawijayan"
		hername = "Zahratun Nisa"
	)

	fmt.Println(myname)
	fmt.Println(hername)

	// default

	var namaku string

	namaku = "Raung"
	fmt.Println(namaku)

	namaku = "Raung K"
	fmt.Println(namaku)

	// Shortcut No 1

	var namanya = "Ayang"
	fmt.Println(namanya)

	// Declare variable int8

	var age int8 = 21
	fmt.Println(age)

	// Declare without var ( instead using := )

	backname := "Nisa"
	fmt.Println(backname)

	backname = "Kawijayan"
	fmt.Println(backname)

}
