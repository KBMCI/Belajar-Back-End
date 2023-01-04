package main

import "fmt"

// function biasa
func sayHello() {
	fmt.Println("Halo")
}

// function parameter
func sayHelloTo(fName string, lName string){
	fmt.Println("Nama Depan :", fName)
	fmt.Println("Nama Belakang :", lName)
}

// function return value
func tambah(a int, b int) int {
	hasil := a+b
	return hasil
}

// function return multiple value/ bisa mengembalikan nilai lebih dari satu 
func getFullName() (string, string, string)  {
	return "Ahmad", "Fathoni", "Gans"
}

// name return value
func getBiodata() (fname, lname string, umur int)  {
	fname = "Ahmad"
	lname = "Fathoni"
	umur = 20

	// return saja disini berarti sudah bisa bisa diartikan mereturn semua value yang dideklarasikan di awal
	// atau bisa jg di buat spt ini : return fname, lname, umur
	return 
}

// variadic function

func main() {

	fmt.Println("==Function biasa==")
	for i := 0; i < 3; i++ {
		sayHello()
	}
	
	// call function parameter
	fmt.Println("==Function parameter==")
	sayHelloTo("Ahmad", "Fathoni")
	
	// call function return value
	fmt.Println("==Function return value==")
	a := 3
	b := 10
	fmt.Println("Hasil", a, "+", b, "=", tambah(a, b))
	
	// implementasi func return multiple value
	fmt.Println("==Function multiple value==")
	fname, lname, _ := getFullName()
	fmt.Println(fname)
	fmt.Println(lname)

	// implementasi func name return value
	fmt.Println("==Function multiple value==")
	fname2, lname2, umur := getBiodata()
	fmt.Println(fname2)
	fmt.Println(lname2)
	fmt.Println(umur)

}