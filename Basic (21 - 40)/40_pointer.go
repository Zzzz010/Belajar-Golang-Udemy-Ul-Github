package main

import "fmt"

type Address struct {
	Name, City, Country string
}

func main() {
	address1 := Address{"Raul", "Blitar", "Indonesia"}
	address2 := &address1

	address2.City = "Bekasi"

	*address2 = Address{"Khayan", "Jakarta", "Indonesia"}

	fmt.Println(address1) // address1 berubah
	fmt.Println(address2)

	address4 := new(Address)
	address4.City = "Malang"
	fmt.Println(address4)
}
