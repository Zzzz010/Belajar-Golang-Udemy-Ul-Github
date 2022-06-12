package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(7)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value" + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}