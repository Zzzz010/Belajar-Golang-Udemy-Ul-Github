package main

import "fmt"

func byebye(name string) string {
	return "Selamat Tinggal " + name
}

func main() {
	sj := byebye

	last := sj("Kamu")

	fmt.Println(last)
	fmt.Println(sj("Dirimu"))
}
