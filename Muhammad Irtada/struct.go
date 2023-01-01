package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHi(name string)  {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	var yanto Customer
	yanto.Name = "yanto"
	yanto.Address = "bekasi"
	yanto.Age = 20

	fmt.Println(yanto)

	//posisi tidak berpengaruh
	agung := Customer{
		Name:    "agung",
		Age:     22,
		Address: "pluto",
	}

	fmt.Println(agung)

	//posisi atau tambahan field dapat berpengaruh
	eko := Customer{"eko", "uranus", 24}
	fmt.Println(eko)
	
	yanto.sayHi("agung")

}
