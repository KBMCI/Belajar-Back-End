package main

import "fmt"

// struct adalah representasi data dalam program aplikasi yang kita buat
// Data di struct disimpan dalam field, sederhananya struct adalah kumpulan field
type Customer struct {
	Name, Address string
	Age           int
}
 
func main() {
	// seperti membuat objek klo di OOP
	var cust Customer
	cust.Name = "Ahmad Fathoni"
	cust.Address = "Indonesia"
	cust.Age = 20

	fmt.Println(cust)
	fmt.Println(cust.Name)
	fmt.Println(cust.Address)
	fmt.Println(cust.Age)

	// Struct Literals (beberapa cara yang bisa digunakan untuk membuat data dr struct)
	aldy := Customer{
		Name: "Aldy",
		Address: "Lowokwaru",
		Age: 40,
	}
	fmt.Println(aldy)

	// pengisian data spt dibawah ini harus sama dengan tata letak data structnya
	faza := Customer{"Faza", "Tidar", 30}
	fmt.Println(faza)
}