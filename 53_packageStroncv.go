package main

import (
	"fmt"
	"strconv"
)

// stroncv.parseBool(string) 		--> Mengubah string ke bool
// stroncv.parseFloat(string) 		--> Mengubah string ke float
// stroncv.parseInt(string) 		--> Mengubah string ke int64
// stroncv.FormatBool(Bool) 		--> Mengubah bool ke string
// stroncv.FormatFloat(Float,...) 	--> Mengubah float 64 ke string
// stroncv.FormatInt(int,...) 		--> Mengubah int64 ke string

func main53() {
	
	// stroncv.parseBool 
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	}else {
		fmt.Println(err.Error())
	}

	// strconv.parseFloat
	number, err := strconv.ParseFloat("1.7", 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// strconv.parseInt
	number2, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number2)
	} else {
		fmt.Println(err.Error())
	}

	//strconv.FormatInt
	// bite int
	// 2 	--> binary
	// 8 	--> oktal
	// 10 	--> decimal
	// 16 	--> hexadecimal

	number3 := strconv.FormatInt(100000, 10)
	fmt.Println(number3)
	

}