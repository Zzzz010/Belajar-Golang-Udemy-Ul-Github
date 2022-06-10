package main

import "fmt"

func endapp() {
	message := recover()
	if message != nil {
		fmt.Println("Aplikasi mengalami ", message)
	}
	fmt.Println("End app")

}

func runapp(error bool) {
	defer endapp()
	if error {
		panic("ERROR")
	}

	fmt.Println("Run app")

}

func main() {
	runapp(true)
}
