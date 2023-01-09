package main

import (
	"Belajar-Back-End/database"
	_ "Belajar-Back-End/database2" 
	// beri tanda " _ " untuk blank identifier karena secara default golang akan melakukan complain jika ada package yg diimpor namun tidak digunakan
	"fmt"
)

func main49() {
	hasil := database.GetDatabase()
	fmt.Println(hasil)
}