package main;

import "fmt";

func getGoodBye(nama string) string  {
	return "Selamat tinggal" + nama
}

func main() {
	
	a := getGoodBye;
	fmt.Println(a(" Aldy"))
};