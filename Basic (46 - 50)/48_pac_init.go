package main

import (
	database "data-base"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
