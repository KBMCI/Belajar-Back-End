package main

import "fmt"

func main()  {
	person := map[string]string{
		"name": "Yanto",
		"address": "Jogja",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// tambah map
	person["title"] = "programmar"
	fmt.Println(person)

	//buat map kosongan
	var book map[string]string = make(map[string]string)
	book["tittle"] = "belajar golang"
	book["author"] = "yanto"

	//hapus data
	delete(book, "author") // hati-hati jika menghapus key yang tidak ada maka tidak terdapat error
	fmt.Println(book)
	fmt.Println(len(book))
}