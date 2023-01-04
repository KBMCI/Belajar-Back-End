package main

import "fmt"

func main() {
	angka := 1

	// seperti while
	for angka <= 10 {
		fmt.Println("Angka ke ", angka)
		angka++
	}

	fmt.Println()
	// for statement biasa
	for i := 1; i <= 10; i++ {
		fmt.Println("Angka ke ", i)
	}
	
	fmt.Println()
	// for range
	slice := []string{"Ahmad", "Fathoni", "Toni", "Budi", "Joko"}
	
	// cara for biasa
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	
	fmt.Println()
	// cara for range
	// key, value := range array/slice
	// _ untuk menandakan variable tdk dibutuhkan
	for _, value := range slice {
		// fmt.Println("index", i, "=", value)
		fmt.Println(value)
	}

	fmt.Println()
	// for range untuk map
	person := make(map[string]string)
	person["name"] = "Fathoni"
	person["title"] = "Mahasiswa"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}