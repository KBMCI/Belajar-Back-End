package main

import "fmt"

type DataCustomer struct {
	Nama, kota string
	umur int
}

// ini yang disebut struct function atau struct method
func (dataCustomer DataCustomer) sayHello(nama string)  {
	fmt.Println("Halo lama tidak berjumpa", nama, ",", dataCustomer.Nama, " dan kalian berasal dari kota", dataCustomer.kota)
	// kota akan diambil di dataCustomer1 yang sudah dideklarasikan di func main
}

func main() {
	dataCustomer1 := DataCustomer {
		Nama: "Aldy Imam Wijaya",
		kota: "Malang",
		umur: 20,
	}	

	dataCustomer1.sayHello("Toni")
}


