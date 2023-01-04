package main

import "fmt"

type Man struct{
	nama string
}

type Woman struct{
	nama string
}

func (woman *Woman) getName2(){
	woman.nama = "Mrs." + woman.nama
}

func (man Man) getName1() {
	man.nama = "Mr." + man.nama
	
}

func main() {
	nama1 := Man{"Aldy"}
	nama1.getName1()

	fmt.Println(nama1)

	nama2 := Woman{"Dewi"}
	nama2.getName2()
	fmt.Println(nama2)
}