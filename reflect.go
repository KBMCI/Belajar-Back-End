package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
	Name string `required:"true" max:"10"` // digunakan untuk membuat tag
	Umur int
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main59()  {
	sample := Sample{"yanto", 42}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType)

	// melihat ada berapa field
	fmt.Println(sampleType.NumField())
	// melihat nama filed ke 1
	fmt.Println(sampleType.Field(1).Name)

	// mengambil data dari tag
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	// contoh kasus penggunaan reflection
	fmt.Println(IsValid(sample))
	sample.Name = ""
	fmt.Println(IsValid(sample))
}