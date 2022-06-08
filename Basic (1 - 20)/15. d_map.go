package main

import (
	"fmt"
)

func main() {
	kendaraan := map[string]string{
		"darat": "kereta",
		"air":   "kapal",
	}

	kendaraan["udara"] = "pesawat"

	fmt.Println(kendaraan["darat"])
	fmt.Println(kendaraan["air"])
	fmt.Println(kendaraan)
	fmt.Println(len(kendaraan))

	book := make(map[string]string)
	book["title"] = "Cara mencekek DIDI"
	book["author"] = "Raul"
	book["dele"] = "no"
	delete(book, "dele")
	fmt.Println(book)

}
