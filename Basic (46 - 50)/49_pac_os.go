package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argumen : ")
	fmt.Println(args)

	fmt.Println("Nama depan : ", args[1])
	fmt.Println("Nama belakang : ", args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	// username := os.Getenv("APP_USERNAME")
	// password := os.Getenv("APP_PASSWORD")

	// fmt.Println(username, password)

}
