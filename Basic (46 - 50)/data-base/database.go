package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init dipanggil")
	connection = "MySql"

}

func GetDatabase() string {
	return connection
}
