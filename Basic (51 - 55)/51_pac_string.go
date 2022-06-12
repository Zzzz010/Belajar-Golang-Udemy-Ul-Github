package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Raung Kawijayan", "Raung"))
	fmt.Println(strings.Split("Raung Kawijayan", " "))
	fmt.Println(strings.ToLower("Raung Kawijayan"))
	fmt.Println(strings.ToUpper("Raung Kawijayan"))
	fmt.Println(strings.Trim("   Raung Kawijayan   ", " "))
	fmt.Println(strings.ReplaceAll("Raul Kawijayan", "Raul", "Raung"))
}
