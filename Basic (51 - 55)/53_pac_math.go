package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(15, 21))
	fmt.Println(math.Min(12, 27))
}