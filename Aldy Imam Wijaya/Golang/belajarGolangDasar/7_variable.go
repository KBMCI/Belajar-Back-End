package main;

import "fmt";

func main() {

	// variable wajib digunakan jika tidak digunakan maka deklarasi akan memunculkan error

	// var dengan Menyebutkan Type Data
	var name string;

	name = "Aldy";
	fmt.Println(name);
	
	name = "Aldy Imam";
	fmt.Println(name);
	
	name = "Aldy Imam Wijaya";
	fmt.Println(name);

	// var tanpa Menyebutkan Type Data

	var nim = 215150700111039;
	fmt.Println(nim);

	// var tanpa Menyebutkan Type Data

	var kelas string= "215150700111039";
	fmt.Println(kelas);

	// var tidak wajib ditulis dengan syarat setelah variable pertama 
	// maka WAJIB menuliskan tanda " := " (tanpa tanda petik) dan
	// hanya berlaku pada saat deklarasi awal saja 

	kota := "Malang";
	fmt.Println(kota);

	kota = "Surabaya";
	fmt.Println(kota);
	
	// multiple deklarasi 
	// type data didepan variable tidak wajib 

	var (
		namaDepan string = "Aldy";
		namaTengah string = "Imam";
		namaBelakang string = "Wijaya";
	)

	fmt.Println(namaDepan);
	fmt.Println(namaTengah);
	fmt.Println(namaBelakang);
};