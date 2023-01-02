package main

import "fmt"

// nil adalah data kosong kalau di java kayak null 
// nil hanya bisa digunakan untuk func, map, slice, array 

func Example(nama string) map[string]string {
	
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama": nama,
		}
	}
	
}

func Example1(kota string) map[string]string  {
	
	if kota == "" {
		return nil
	} else {
		return map[string]string{
			"kota": kota,
		}
	}
	
}


func main() {
	personZero := Example("")
	person := Example("Aldy")

	fmt.Println(personZero)
	fmt.Println(person)

	// anjassss bingung di contoh kedua
	// contoh kedua 

	  var data1 map[string]string = Example1("Malang")

	 if data1 == nil {
		return 
	 } else {
		return 
	 }
}