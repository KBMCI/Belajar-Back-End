package main

import "fmt"

// masih bingung 
func endApp()  {
	fmt.Println("Aplikasi sudah dijalankan")
}

func startApp()  {
	defer endApp()

	fmt.Println("Aplikasi mulai")
}

func main() {
	
}