package main

import "fmt"

func pertama()  {
	fmt.Println("Ini sudah pasti dieksekusi")
}


func kedua()  {
	defer pertama()
	fmt.Println("Setelah ini dieksekusi maka...")
}
func main() {
	kedua()
}