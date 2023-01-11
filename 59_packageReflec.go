package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10" ` // reflec.StructTag
}

type CobaLagi struct {
	Name string
	Description string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}

func main59() {
	
	sample := Sample{
		Name: "Aldy",
	}	

	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField()) // menghitung jumlah field 
	fmt.Println(sampleType.Field(0).Name)

	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min")) // output kosong
	
	sample.Name = "" // output akan false jika sample.Name tidak sesuai dengan field yang ada pada struct
	fmt.Println(IsValid(sample))

	contoh := CobaLagi{
		Name: "Aldy",
		Description: "Sangat oke", 
	}

	// output akan selalu true karena pada field struct CobaLagi tidak ada structTag nya 
	fmt.Println(IsValid(contoh))
}