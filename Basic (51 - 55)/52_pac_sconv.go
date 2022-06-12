package main

import (
	"fmt"
	"strconv"
)

func main() {
	int, err := strconv.ParseInt("1000", 2, 8)
	if err == nil {
		fmt.Println(int)
	} else {
		fmt.Println("Error", err.Error())
	}

	nilai := strconv.FormatInt(1000000, 10)
	fmt.Println(nilai)

	nilaiInt, _ := strconv.Atoi("300")
	fmt.Println(nilaiInt)

}
