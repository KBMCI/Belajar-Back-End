package main;

import "fmt";

// perbedaan dengan bab 24 ada di baris 5
func getBiodata2() (namaLengkap string, nim int64){
	namaLengkap = "Aldy Imam Wijaya";
	nim = 215150700111039;

	return 
	
	// kata return wajib ditulis 
	// setelah kata return tidak wajib ditulis 
};

func main() {
	var namaLengkap, nim = getBiodata2();

	fmt.Println("NAMA \t: ", namaLengkap);
	fmt.Println("NIM \t: ", nim);

	
};