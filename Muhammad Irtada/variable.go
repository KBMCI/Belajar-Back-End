package main

import "fmt"

func main()  {
	//deklrasi variabel
	var name string

	//var name hanya bisa di assign dengan tipe data string
	name = "yanto"
	fmt.Println(name)

	//deklarasi auto detect tipe data
	var umur = 20
	fmt.Println(umur)

	//buat variabel tanpa var tapi harus langsung di assign datanya
	isKawin := false
	fmt.Println(isKawin)

	//multiple deklarasi
	var (
		firstName = "yanto"
		lastName = "agung"
	)

	fmt.Println(firstName,lastName)
}