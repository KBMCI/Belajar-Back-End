package main

import "fmt"

func main()  {
	sayHello()
	sayHello2("yanto", "agung")

	dataHello := getHello("agung")
	fmt.Println(dataHello)	

	firstName, lastName := getFullName()
	println("fullname =", firstName, lastName)

	a, b := getFullName2()
	fmt.Println(a, b)

	fmt.Println(sumAll("test aja", 32, 5, 6, 46, 75))

	slice := []int{
		13,
		24,
		5,
	}
	fmt.Println(sumAll("test", slice...))

	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("yanto"))

	sayHelloWithFilter("anjing", spamFilter)

	blacklist := func (name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("yanto", blacklist)
	registerUser("yanto", func(name string) bool {
		return name == "yanto"
	})

	fmt.Println(factorialRecursive(5))
}

func sayHello()  {
	fmt.Println("hello world")
}
func sayHello2(firstName string, lastName string)  {
	fmt.Println("hello", firstName, lastName)
}
func getHello(name string) string {
	return "hello " + name
}
func getFullName() (string, string) {
	return "yanto", "agung"
}
func getFullName2() (firstName, lastName string) {
	firstName = "yanto"
	lastName = "agung"
	return
}

//variadic - variable arguments
//variadic hanya bisa diletakkan di akhir parameter dan hanya bisa dibuat satu saja
//variadic merupakan slice sehingga bisa menerima parameter slice yang dirubah ke variadic terlebih dahulu
func sumAll(name string, numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

//function value
//digunakan saat menjadikan suatu fungsi sebagai parameter
func getGoodBye(name string) string {
	return "good bye " + name
}

//penerapan function value
type Filter func(string) string
func sayHelloWithFilter(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}else {
		return name
	}
}

//anonymous function
//function dapat langsung di assign ke variabel atau di program
type Blacklist func(string) bool
func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println(name)
	}else {
		println("welcome", name)
	}
}

//recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}else {
		return value * factorialRecursive(value - 1)
	}
}