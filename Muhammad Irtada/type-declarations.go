package main

import "fmt"

func main()  {
	//dibuat untuk membuat alias dari tipe data yang sudah ada
	//bertujuan agar lebih mudah dimengerti
	type noKTP string
	type married bool

	var noKTPYanto noKTP = "123123123412"
	var marriedStatus married = false

	fmt.Println(noKTPYanto)
	fmt.Println(marriedStatus)
}