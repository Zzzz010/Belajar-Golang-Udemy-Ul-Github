package main

import "fmt"

func nama() (firstname, lastname string) {
	firstname = "Raung"
	lastname = "Kawijayan"

	return
}

func main() {
	firstname, lastname := nama()
	fmt.Println(firstname, lastname)
}
