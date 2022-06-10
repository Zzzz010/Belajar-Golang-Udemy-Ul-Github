package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	a := sumAll(10, 10, 10, 10)
	fmt.Println(a)

	slice := []int{30, 20, 40, 50, 70}
	a = sumAll(slice...)
	fmt.Println(a)
}
