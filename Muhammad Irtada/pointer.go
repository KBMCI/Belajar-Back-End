package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address)  {
	address.Country = "Indonesia"
}

type Man struct{
	name string
}

func (man *Man) Married()  {
	man.name = "Mr. " + man.name
	fmt.Println("di method", man.name)
}

func main()  {
	var address1 Address = Address{
		City:     "Subang",
		Province: "Jawa barat",
		Country:  "indonesia",
	}

	//pass by value
	// address2 := address1

	//pass by refference
	address2 := &address1
	var address3 *Address = &address1

	address2.City = "bandung"

	// yang dirubah bukan address2 saja tapi semua yang memiliki alamat sama dengan address2
	*address2 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//membuat address tanpa data
	var address4 *Address = new(Address)
	fmt.Println(address4)


	//pointer di function
	alamat := Address{
		City:     "Bogor",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)


	//pointer di method
	yanto := Man{"yanto"}
	yanto.Married()
}