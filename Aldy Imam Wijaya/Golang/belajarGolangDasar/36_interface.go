package main
import "fmt"

type HasName interface{
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Halo", hasName.GetName())
}

type Person struct{
	Nama string
}

func (person Person) GetName() string  {
	return person.Nama
}

func main() {
	var aldy Person 
	aldy.Nama = "Aldy"

	sayHello(aldy)

	// bingung juga si sksksksksk
}