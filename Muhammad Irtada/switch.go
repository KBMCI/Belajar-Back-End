package main

import "fmt"

func main() {
	name := "yanto"

	switch name {
	case "yanto":
		fmt.Println("halo yanto")
	case "agung":
		fmt.Println("halo agung")
	default:
		fmt.Println("sapa anda")
	}

	//short switch
	switch length := len(name); length {
	case 5:
		fmt.Println("apakah kamu yanto")
	}

	//switch tanpa kondisi
	umur := 20
	switch {
	case umur != 20:
		fmt.Println("kamu tidak cukup umur nak balek lagi taun depan")
	case len(name) < 10:
		println("namamu kependekan")
	}
}
