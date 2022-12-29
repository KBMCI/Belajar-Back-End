package main;

import "fmt";

func main() {

	// membuat array manual 
	var daftarNama [3]string;

	daftarNama[0] = "Aldy";
	daftarNama[1] = "Imam";
	daftarNama[2] = "Wijaya";

	fmt.Println(daftarNama[0]);
	fmt.Println(daftarNama[1]);
	fmt.Println(daftarNama[2]);

	// membuat array langsung 

	var values = [5] int {
		10,
		20,
		30,
	};

	fmt.Println(values);

	// function array 
	// " len(array) 		   --> menghitung jumlah panjang array"
	// " array[index]          --> mendapatkan data di posisi index"
	// " array[index] = value  --> mengubah data di posisi index"

	fmt.Println(len(daftarNama));
	fmt.Println(len(values));
	
	values[4] = 100;
	fmt.Println(values);

	
};