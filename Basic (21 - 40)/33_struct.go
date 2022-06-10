package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int8
}

func main() {

	// cara pertama
	var raul Customer

	raul.Name = "Raung"
	raul.Address = "Blitar"
	raul.Age = 21

	fmt.Println(raul)
	fmt.Println(raul.Name)
	fmt.Println(raul.Address)
	fmt.Println(raul.Age)

	// cara kedua
	khayan := Customer{
		Name:    "Khayan",
		Address: "Jakarta",
		Age:     22,
	}

	fmt.Println(khayan)

	// cara ketiga
	Excel := Customer{"Excel", "Bekasi", 20}
	fmt.Println(Excel)
}
