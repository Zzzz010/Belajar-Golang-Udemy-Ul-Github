package main

import "fmt"

type Filter func(string) string

func kalimatfilter(name string, filter Filter) {
	namefilter := filter(name)
	fmt.Println("Hello", namefilter)
}

func spamfilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	kalimatfilter("Raul", spamfilter)

	kalimatburuk := spamfilter

	kalimatfilter("Anjing", kalimatburuk)
}
