package main

import "fmt"

func main() {
	name := "yanto"

	if name == "yanto" {
		fmt.Println("hallo yanto")
	} else if name == "agung" {
		fmt.Println("ternyata agung yang datang")
	} else {
		fmt.Println("wah kamu bukan yanto")
	}

	//short if
	if length := len(name); length < 8 {
		fmt.Println("pendek kali la namamu")
	}

	
}
