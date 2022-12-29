package main

import "fmt"

func main() {
	// multiple constant
	const (
		// tipe data const wajib di inisialisasi
		fname string = "Ahmad"
		// deklarasi tanpa tipe data
		lname = "Fathoni" 
		// const tidak bisa dibuat dengan :=, jika dibuat spt itu maka jatuhnya variable biasa
	)

	// error karena const nilainya tidak bisa dirubah
	// fname = "Toni"

	fmt.Println(fname)
	fmt.Println(lname)
}