package main;

import "fmt";

func main() {

	angka := 1;

	for angka <= 10 {
		
		fmt.Println("Ini angka ke - ", angka);
		angka++
	};

	// for dengan statement 

	for i := 0; i < 10; i++ {
		fmt.Println("Ini angka ke - ", i);
	};

	// for loop dengan slice 
	slice := []string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	};

	for j := 0; j < len(slice); j++ {
		fmt.Println(slice[j]);
	};
	
	//for range
	for k, bulan := range slice {
		fmt.Println("Bulan ke", k, " = ", bulan);
	};
	
	// tanda " _ " setelah for berguna untuk memberitahu sistem agar variable tidak digunakan 
	
	for _, bulan := range slice {
		fmt.Println(bulan);
	};

	for k, _ := range slice {
		fmt.Println("Bulan ke", k);
	};

	// for loop dengan map 

	person := make(map[string]string);

	person["nama"] = "Aldy";
	person["kota"] = "Malang";

	for key, v := range person {
		fmt.Println(key, "=", v);
	};

};

