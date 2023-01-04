package main

import "fmt"

func main() {

	name := "Ahmad"

	if name == "Toni" {
		fmt.Println("Halo Toni")
	} else if name == "Ahmad" {
		fmt.Println("Halo Ahmad")
	} else {
		fmt.Println("Hai Kenalan Dong")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")		
	}else{
		fmt.Println("Nama sudah benar")		
	}

}