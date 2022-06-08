package main

import "fmt"

func main() {
	years := [...]int16{
		2000,
		2001,
		2002,
		2003,
		2004,
		2005,
		2006,
		2007,
		2008,
		2009,
		2010,
	}

	slice1 := years[5:8]
	slice2 := years[:]

	fmt.Println(slice1[1])
	fmt.Println(slice1[2])

	fmt.Println(slice2[10])
	fmt.Println(len(years))

	years[9] = 666
	fmt.Println(slice2[9])

	tanggal := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tanglice := tanggal[1:]
	tanglice[0] = 0
	tanglice[6] = 1
	fmt.Println(tanglice)

	tangjelek := append(tanglice[:], 666)
	tangjelek[0] = 000
	fmt.Println(tangjelek)
	fmt.Println(tanggal)

	slicebaru := make([]string, 4, 10)
	slicebaru[0] = "Minggu"
	slicebaru[3] = "Bye"

	fmt.Println(slicebaru)
	fmt.Println(len(slicebaru))
	fmt.Println(cap(slicebaru))

}
