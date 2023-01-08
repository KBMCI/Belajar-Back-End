package database2

import "fmt"
var connection string

func init() {

	fmt.Println("Test")
	// mengecek functioin init dapat berjalan tanpa mengeksekusi function lainya 
	connection = "My SQL Workbench"
}

func GetDatabase2() string {
	return connection	
}