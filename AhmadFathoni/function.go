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
// variadic hanya bisa dideklarasikan 1x saja
// varaiadic harus diletakkan di parameter paling terakhir 
func sumAll(str string, angka ...int) int  {
	total := 0

	for _, value := range angka {
		total += value
	}

	return total
}

// function value
// maksudnya yaitu function bisa disimpan di sebuah variable dan akan dianggap sbg valuenya
func getGoodBye(name string) string  {
	return "Good Bye " + name
}

// Function as Parameter
// Function type declarations
type Filter func (string) string

func sayHelloWithFilter(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string  {
	if name == "Anjing" {
		return "***"
	}else {
		return name
	}
}

// Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

// Function recursive
func factorialRecursive(value int) int  {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

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
	
	// implementasi func variadic
	fmt.Println("==Function Variadic==")
	total := sumAll("coba str", 10, 20, 30, 40)
	fmt.Println(total)

	// jika sudah ada slice dan ingin digunakan di variable argumen/variadic bisa menggunakan "(nama slice)..."
	slice := []int{10, 10, 10, 10}
	fmt.Println(sumAll("Halo", slice...))

	// implementasi func value
	// jika ingin passing func ke dalam variable maka jg gunakan tanda kurung 
	fmt.Println("==Function Value==")
	sayGoodBye := getGoodBye 
	fmt.Println(sayGoodBye("Toni"))
	
	// implementasi func as parameter
	fmt.Println("==Function as Parameter==")
	sayHelloWithFilter("Toni", spamFilter)
	
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
	
	// implementasi func anonymous
	fmt.Println("==Anonymous Function==")
	// anonymous func bisa ditaruh di sebuah variable
	blacklist := func(name string) bool  {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Toni", blacklist)
	
	// atau anonymous func juga bisa langsung di deklarasikan di parameternya
	registerUser("root", func(name string) bool {
		return name == "root"
	}) 
	
	registerUser("Toni", func(name string) bool {
		return name == "root"
	}) 
			
	// implementasi Function recursive
	fmt.Println("==Function Recursive==")
	factorial := factorialRecursive(5)
	fmt.Println(factorial)
	fmt.Println(factorialRecursive(10))

}