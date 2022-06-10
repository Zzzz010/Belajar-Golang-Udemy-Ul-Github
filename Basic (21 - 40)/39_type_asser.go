package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	hasil := random()

	switch value := hasil.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)

	default:
		fmt.Println("Unknown")
	}

	hasilString := hasil.(string)
	fmt.Println(hasilString)

	//hasilInt := hasil.(int) // panic
	//fmt.Println(resultInt)

}
