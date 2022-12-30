package main

import "fmt"

func main()  {
	var nilai32 int32 = 128000
	var nilai64 int64 = int64(nilai32)
	//bisa tapi balik ke nilai terendah int8
	var nilai8 int8 = int8(nilai32) 

	fmt.Println(nilai64)
	fmt.Println(nilai8)

	name := "yanto"
	var e = name[0]
	eString := string(e)

	fmt.Println(name)
	fmt.Println(eString)
}