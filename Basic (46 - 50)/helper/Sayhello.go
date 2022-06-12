package helper

import "fmt"

var version = 1 // tidak bisa diakses dari luar
var Application = "Belajar Golang"

// Tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Hello" + name
}

func Sayhello(name string) {
	fmt.Println("Hello", name)
}
