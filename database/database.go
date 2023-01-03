package database

import "fmt"

var connection string

// fungsi init digunakan untuk menjalankan fungsi saat package pertama kali digunakan
func init() {
	connection = "MySQL"
	fmt.Println("init terpanggil")
}

func GetDatabase() string {
	return connection
}