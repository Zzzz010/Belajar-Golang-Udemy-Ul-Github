package main

import "fmt"

func main() {
	name := "Raul"
	counter := 0

	increment := func() {
		name := "Jayan"
		fmt.Println("increment")
		counter++
		fmt.Println(name)

	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
