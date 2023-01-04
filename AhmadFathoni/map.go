package main

import "fmt"

func main() {

	// cara buat map = map[key]value
	// map sama spt array assosiasi
	person := map[string]string{
		"name":    "Toni",
		"address": "Malang",
	}

	// menambah data baru yang awalnya tdk ada
	person["jk"] = "L" 
	
	fmt.Println(person)
	// len : panjang dari map
	fmt.Println(len(person))
	// cara mengambil data di map dg key
	fmt.Println("nama : ", person["name"])
	fmt.Println("alamat : ", person["address"])
	fmt.Println("jenis kelamin : ", person["jk"])
	
	book := make(map[string]string)
	book["title"] = "Judul Buku"
	book["author"] = "Penulis"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}