package main

import "fmt"

type Address struct{
	kota, provinsi, negara string
}

func Kota1(alamat Address) {
	alamat.negara = "Indonesia"
	
}

func Kota2(alamat *Address) {
	alamat.negara = "Malaysia"
	
}

func main() {

	var alamat1 = Address{
		kota: "Malang",
		provinsi: "Jawa Timur",
		negara: "",
	}
	
	Kota1(alamat1)
	fmt.Println(alamat1)

	var alamat2 = &Address{
		kota: "Konoha",
		provinsi: "Serawak",
		negara: "",
	}

	Kota2(alamat2)
	fmt.Println(alamat2)
}