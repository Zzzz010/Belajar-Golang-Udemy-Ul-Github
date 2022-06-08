package main

import "fmt"

func main() {
	var nilai1 bool = true
	var nilai2 bool = false
	var result = !nilai1 && nilai2

	fmt.Println(result)
	fmt.Println(nilai1 || nilai2)

	var nilaiujian = 100
	var absenkehadiran = 45

	var lulusujian = nilaiujian > 70
	var lulusabsen = absenkehadiran > 60
	var lulus = lulusujian && lulusabsen

	fmt.Println(lulus)

}
