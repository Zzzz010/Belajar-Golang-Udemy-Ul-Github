package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`r([a-z])`)

	fmt.Println(regex.MatchString("raul"))
	fmt.Println(regex.MatchString("raun"))
	fmt.Println(regex.MatchString("rAug"))

	fmt.Println(regex.FindAllString("raul, ,rau, raun, raug, aul", 2))
}
