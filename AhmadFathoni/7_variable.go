package main

import "fmt"

func main() {
	var name string
	var age = 20
	// inisialisasi tanpa var
	country := "Indonesia"
	var (
		fname = "Ahmad"
		lname = "Fathoni"
	)

	name = "Ahmad Fathoni"
	fmt.Println(name)

	name = "Toni"
	fmt.Println(name)

	fmt.Println(age)
	fmt.Println(country)
	fmt.Println(fname)
	fmt.Println(lname)
}