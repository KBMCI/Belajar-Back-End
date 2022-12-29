package main;

import "fmt";

func main() {

	// constant tidak wajib digunakan tidak seperti variable
	// nilai constant tidak dapat diubah 

	const namaDepan = "Aldy";
	const namaTengah = "Imam";
	const namaBelakang = "Wijaya";
	const nim = 215150700111039;
	

	fmt.Println(namaDepan);
	fmt.Println(namaTengah);
	fmt.Println(namaBelakang);

	// nim tidak digunakan tapi tidak error karena memakai constant

	// multiple constant 
	// tidak sama dengan multiple variable yang menggunakan " := " diawal deklarasi setelah variable

	const(

		namaDepanSaya = "Aldy";
		namaTengahSaya = "Imam";
		namaBelakangSaya = "Wijaya";
		nomorIndukMahasiswa = 215150700111039;

	);


	
};