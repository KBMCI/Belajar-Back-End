package main;

import "fmt";

func main() {
	
	var nama string = "Aldy";
	kota := "Malang";

	if nama == "Aldy" {

		fmt.Println("ANJAY ALDY");

	} else if nama == "Imam" {

		fmt.Println("ANJAY IMAM");

	} else {
		
		fmt.Println("ALDY BOHONG");

	};

	if kota != "Surabaya" {

		fmt.Println("Berarti Malang");
	};

	// if short statement

	if panjangKota := len(kota); panjangKota > 5 {

		fmt.Println("Kepanjangan nama kotanya")

	} else {

		fmt.Println("Nama kotanya singkat")
		
	}
}