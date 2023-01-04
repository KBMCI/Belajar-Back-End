package main

import "fmt"

type Address struct {
	kota, provinsi, negara string
}

func main() {
	alamat1 := Address{
		kota:     "Malang",
		provinsi: "Jawa Timur",
		negara:   "Indonesia",
	}

	fmt.Println(alamat1)

	alamat2 := alamat1
	alamat2.kota = "Surabaya"
	fmt.Println(alamat2)

	// tanda " & " digunakan untuk mereference ke alamat yg dituju
    alamat3 := &alamat1
	fmt.Println(alamat3)

	alamat3 = &Address{
		kota: "Lampung",
		provinsi: "Sumatra Barat",
		negara: "Indonesia",
	}

	fmt.Println(alamat3)

	alamat4 := alamat3
	alamat5 := alamat3
	alamat6 := alamat3

	// tanda " * " digunakan untuk menunjukkan bahwa alamat 4,5,6 ikut pada alamat 3 yang baru
	*alamat3 = Address{
		kota: "Jogja",
		provinsi: "DIY",
		negara: "Indonesia",
	}

	fmt.Println(alamat4)
	fmt.Println(alamat5)
	fmt.Println(alamat6)

	// function new 

	alamat7 := new(Address)
	fmt.Println(alamat7)

	alamat8 := new(Address)
	alamat8.kota = "Bandung"
	fmt.Println(alamat8)

}