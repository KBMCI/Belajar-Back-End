package main

import "fmt";

func main() {

	nama := "Faza";

	switch nama {
	case "Aldy" : 
		fmt.Println("Halo Aldy");
	case "Toni" :
		fmt.Println("Halo Konton");
	case "Faza" :
		fmt.Println("Halo Risol");
	case "Bariq" :
		fmt.Println("Halo Babi");
	default :
		fmt.Println("Sabilah kenalan hehe...")
	}

	// switch case short statement 

	var kota string = "Surabaya";

	switch panjangKota := len(kota); panjangKota > 5{
	case true :
		fmt.Println("Nama kotanya panjang");
	case false :
		fmt.Println("Nama kotanya pendek");
	} 

	// switch tanpa kondisi 

	var nim uint64 = 215150700111039;

	switch  {
	case nim > 200000000000000 :
		fmt.Println("Anak FILKOM");
	case nim < 200000000000000 :
		fmt.Println("Bukan anak FILKOM");
	default : 
		fmt.Println("Anak mana lu ?? ")

	}
};