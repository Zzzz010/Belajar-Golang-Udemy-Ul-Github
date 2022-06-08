package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "Raung"
	nama[1] = "Kawijayan"
	nama[2] = "Raul"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int8{
		50,
		60,
		80,
	}

	fmt.Println(values)
	values[0] = 100
	fmt.Println(values)

	fmt.Println(len(nama))
	fmt.Println(len(values))
}
