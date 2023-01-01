package main

import "fmt"

type HasName interface{
	GetName() string
}

func SayHello(hasname HasName)  {
	fmt.Println("hello", hasname.GetName())
}

// untuk menghubungkan dengan interface dengan cara menulis nama method dengan return type yang sama seperti method yang ada di interface
type Person struct{
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// interface kosong
func Ups(i int) interface{} {
		if i == 1 {
			return 1
		}else if i == 2{
			return true
		}else {
			return "ups"
		}
}

func main()  {
	yanto := Person{
		Name: "Yanto",
	}

	SayHello(yanto)

	// var ups int = Ups(1) error karena Ups satu mereturn interface kosong bukan int
	var ups interface{} = Ups(1)
	fmt.Println(ups)
}