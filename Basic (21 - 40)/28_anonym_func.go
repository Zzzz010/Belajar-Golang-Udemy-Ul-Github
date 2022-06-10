package main

import "fmt"

type Blacklist func(string) bool

func register_user(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}

}

func main() {
	blacklist := func(name string) bool {
		return name == "Kecoa"
	}

	register_user("Raul", blacklist)
	register_user("Kecoa", blacklist)
	register_user("Anjing", func(name string) bool {
		return name == "Anjing"
	})
	register_user("Meow", func(name string) bool {
		return name == "Guk"
	})

}
