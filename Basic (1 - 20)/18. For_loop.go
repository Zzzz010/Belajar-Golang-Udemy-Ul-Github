package main

import "fmt"

func main() {
	counter := 0

	for counter <= 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}

	for countD := 10; countD >= 0; countD-- {
		fmt.Println("Nilainya ke ", countD)
	}

	name := []string{"Raung", "Raul", "Aung", "Ung"}

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	for index, names := range name {
		fmt.Println("Index", index, "=", names)
	}

	buku := make(map[string]string)

	buku["Judul"] = "Menjadi Orang Kaya"
	buku["Author"] = "Raul"

	for key, books := range buku {
		fmt.Println(key, "=", books)

	}

}
