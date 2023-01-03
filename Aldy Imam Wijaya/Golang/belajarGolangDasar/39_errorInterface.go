package main

import (
	"errors"
 	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	
	if pembagi == 0 {
		return 0, errors.New("karena pembagi tidak boleh nol")
	}else {
		hasil := nilai / pembagi
		return hasil, nil
	}
	
}

func main() {

	hasil, err := Pembagian(100,0)
	if err == nil {
		fmt.Println("Hasilnya", hasil)
	}else {
		fmt.Println("Hasilnya Error", err.Error())
	}
}