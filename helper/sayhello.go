package helper

import "fmt"

var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("hello", name)
}

func sayGoodbye(name string)  {
	fmt.Println("Goodbye", name)
}