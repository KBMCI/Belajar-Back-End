package main
				
import "fmt"

// deklarasi structnya
type BiodataCustomer struct {
	Nama, Kota, JenisKelamin string
	Umur, Nim int64

}

func main() {
	// menggunakan cara ke 1 
	var biodataCustomer BiodataCustomer
	biodataCustomer.Nama = "Aldy Imam Wijaya"
	biodataCustomer.Kota = "Malang"
	biodataCustomer.JenisKelamin = "Laki - laki"
	biodataCustomer.Umur = 20
	biodataCustomer.Nim = 215150700111039

	fmt.Println(biodataCustomer)
	// jika diprint satu satu 
	fmt.Println(biodataCustomer.Nama)
	fmt.Println(biodataCustomer.Kota)
	fmt.Println(biodataCustomer.JenisKelamin)
	fmt.Println(biodataCustomer.Umur)
	fmt.Println(biodataCustomer.Nim)

	// menggunakan cara ke 2 
	var biodataCustomer2 BiodataCustomer = BiodataCustomer{
		Nama: "Ahmad Fathoni",
		Kota: "Malang",
		JenisKelamin: "Laki - laki",
		Umur: 21,
		Nim: 215150700111040,
	} 
	
	fmt.Println(biodataCustomer2)

	// menggunakan cara ke 3 (tidak disarankan, karena rentan error)
	biodataCustomer3 := BiodataCustomer{"Faza", "Malang", "Laki - laki", 20, 215150700111050}

	fmt.Println(biodataCustomer3)
}