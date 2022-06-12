package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Raung")
	data.PushBack("Kawi")
	data.PushBack("Jayan")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}
}
