package main

import (
	"fmt"
	"os"
)

// Args digunakan untuk mengambil data
// Hostname digunakan untuk mengambil data host sistem operasi

func main50() {
	args := os.Args
	fmt.Println("Argument : ")

	fmt.Println(args)

	// buka terminal dulu lalu ketik " go run 50_packageOs.go Aldy Imam Wijaya "
	// baru eksekusi program 
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	nama, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", nama)
	}else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}