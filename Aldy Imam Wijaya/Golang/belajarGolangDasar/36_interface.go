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

type Animal struct{
	Nama string
}

func (animal Animal) GetName() string {
	return animal.Nama
}

func main() {
	// cara 1 
	var aldy Person 
	aldy.Nama = "Aldy"

	sayHello(aldy)

	// cara 2
	cat := Animal{
		Nama: "Oyen",
	}

	sayHello(cat)
	// bingung juga si sksksksksk
}