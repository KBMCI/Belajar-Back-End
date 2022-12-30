package main

import "fmt"

func main()  {
	runApplication(0)
}

//deffer digunakan untuk memanggil fungsi lain meskipun fungsi yang memanggil terdapat 
//panic digunakan untuk mengeluarkan error
//recover untuk menangkap pesan error
func logging()  {
	message := recover()
	if message != nil {
		fmt.Println("error dengan pesan :", message)
	}
	fmt.Println("selesai memanggil function")
}
func runApplication(value int)  {
	defer logging()
	fmt.Println("Run application")
	if value == 0 {
		panic("Tidak dapat dibagi dengan 0")
	}
	result := 10 / value
	fmt.Println("result =", result)
}