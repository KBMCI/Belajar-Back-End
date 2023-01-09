package main

import (
	"flag"
	"fmt"
)

func main51() {
	
	host := flag.String("host", "aldy", "Put Your Host")
	username := flag.String("username", "root", "Put Your username")
	password := flag.String("password", "Localhost", "Put Your password")
	var number *int = flag.Int("Number", 100, "Put your number") 

	flag.Parse()

	fmt.Println("host :",*host)
	fmt.Println("username :",*username)
	fmt.Println("password :",*password)
	fmt.Println("number :",*number)
}