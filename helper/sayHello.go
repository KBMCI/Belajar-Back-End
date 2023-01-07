package helper

import "fmt"

var Application = "Belajar Golang"
var version int = 1

// version := 1 // ini error gatau kenapa, jadi harus dideskripsikan pake kata kunci var, type datanya ga wajib

func SayHello(name string) {
	fmt.Println("hello", name)
}

func sayGoodbye(name string)  {
	fmt.Println("Goodbye", name)
}