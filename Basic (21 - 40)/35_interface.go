package main

import "fmt"

type Person struct {
	Name string
}

type hasname interface {
	getName() string
}

type Animal struct {
	Name string
}

func sayHi(hasName hasname) {
	fmt.Println("Hello", hasName.getName())
}

func (person Person) getName() string {
	return person.Name
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	orang := Person{Name: "Raung"}
	sayHi(orang)

	animal := Animal{Name: "Harimau"}
	sayHi(animal)
}
