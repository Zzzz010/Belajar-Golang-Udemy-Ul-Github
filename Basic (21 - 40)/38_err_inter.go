package main

import (
	"errors"
	"fmt"
	//"errors"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	hasil, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}
