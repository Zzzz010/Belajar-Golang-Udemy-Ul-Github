package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello() {
	fmt.Println("Hallo..", customer.Name)
}

func main() {
	rudy := Customer{Name: "Rudy"}
	rudy.sayHello()
}
