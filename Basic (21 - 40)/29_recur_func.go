package main

import "fmt"

//func factorialloop(value int) int {
//	result := 1
//	for i := value; i > 0; i-- {
//		result *= i
//	}
//
//	return result
//}

func factorrecur(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorrecur(value-1)
	}
}

func main() {
	fmt.Println(factorrecur(10))
	fmt.Println(10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1)
}
