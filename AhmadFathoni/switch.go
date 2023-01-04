package main

import "fmt"

func main() {

	name := "Toni"

	switch name {
	case "Toni":
		fmt.Println("Halo Toni")
	case "Ahmad":
		fmt.Println("Halo Ahmad")
	default:
		fmt.Println("Kenalan yuk")
	}

	// short statement switch
	switch panjang := len(name); panjang > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")	
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Sesuai")
	}
}