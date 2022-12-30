package main

import "fmt"

func main() {
	const nama = "yanto"
	//ERROR
	//nama = "agung"
	fmt.Println(nama)

	// multi const
	const (
		firstName = "yanto"
		lastName  = "agung"
	)

}
