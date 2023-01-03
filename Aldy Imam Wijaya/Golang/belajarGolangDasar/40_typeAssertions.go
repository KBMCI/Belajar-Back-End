package main

import "fmt"


// type assertions digunakan untuk mengubah tipe data yang ada pada interface
func TypeAssertions() interface{} {
	return true
}

func main() {
	hasil := TypeAssertions()

	// cara 1 (kurang aman)
	// hasilKonversi := hasil.(string)

	// fmt.Println(hasilKonversi)

	// cara 2 (main aman)

	switch value := hasil.(type) {
	case string:
		fmt.Println("Hasilnya", value, " yang berupa string")
	case int:
		fmt.Println("Hasilnya", value, " yang berupa int")
	case bool:
		fmt.Println("Hasilnya", value, " yang berupa boolean")
	default:
		fmt.Println("Hasilnya", value, " yang berupa default")

	}
}