package main

import "fmt";

func main() {
				
	nama := "Aldy Imam Wijaya"
	kota := "Malang"
	
	fungsiBebas := func ()  {
		nama = "Ahmad Fathoni"
		kota := "Surabaya"
		fmt.Println(nama)
		fmt.Println(kota)

		
	}
	
	fungsiBebas()
	fmt.Println(nama) 	// variable nama akan menjadi ahmad fathoni karena variable nama aldy imam wijaya tertumpuk
						// pada anonymous function di fungsi Bebas 
	fmt.Println(kota)
 }