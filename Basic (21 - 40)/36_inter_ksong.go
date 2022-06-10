package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 5
	} else if i == 2 {
		return true
	} else {
		return "OK"
	}

}

func main() {
	var nilai interface{} = Ups(1)
	fmt.Println(nilai)
}
