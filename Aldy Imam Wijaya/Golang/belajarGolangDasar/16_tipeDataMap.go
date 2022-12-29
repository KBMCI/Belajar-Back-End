package main;

import "fmt";

func main() {

		var biodata = map[string]string{
			"Nama" : "Aldy",
			"Alamat" : "Malang",
		};
		
		//menambahkan map
		biodata["Gelar"] = "S1";

	fmt.Println(biodata);
	fmt.Println(biodata["Nama"]);
	fmt.Println(biodata["Alamat"]);
	fmt.Println(biodata["Gelar"]);

	// function map 
	// len(map) 					--> mendapatkan jumlah data di map 
	// map[key] 					--> mengambil data di map dengan key 
	// map[key] 					--> value  mengubah data di map dengan key 
	// make(map[TypeKey]TypeValue) 	--> membuat map baru 
	// delete(map, key)			 	--> menghapus data di map dengan key
	
	buku := make(map[string]string)
		 buku["Judul"] = "Jejak Si Fajar";
		 buku["Author"] = "Fajar";
		 buku["Ups"] = "Salah";
	
	fmt.Println(buku["Judul"]);
	fmt.Println(buku["Author"]);

	delete(buku, "Ups")
	
	
};