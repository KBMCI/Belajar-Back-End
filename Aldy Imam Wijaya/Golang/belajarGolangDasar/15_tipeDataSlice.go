package main

import "fmt";

func main() {
	
	// membuat slice dari array 
	// array[low:high]  --> membuat slice dari array dimulai dari index low sampai index sebelum high 
	// array[low:]      --> membuat slice dari array dimulai dari index low sampai index akhir di array 
	// array[:high]     --> membuat slice dari array dimulai dari index 0 sampai index sebelum high 
	// array[:]         --> membuat slice dari array dimulai dari index 0 sampai index akhir di array  
	
	// jika tidak tahu berapa panjang array yg akan ditulis maka diisi " ... "
	// contoh var jumlahRambut = [...] int {}

	var bulan = [12] string {
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

	var slice1 = bulan[4:7];
	fmt.Println(slice1);

	// function slice 
	// len(slice) 							--> mendapatkan panjang
	// cap(slice) 							--> mendapatkan kapasitas
	// append(slice,data) 					--> membuat slice baru dengan menambah data ke posisi terakhir slice, 
	//                        					jika kapasitas sudah penuh, maka akan membuat array baru
	// make([]TypeData, length, capacity) 	--> membuat slice baru
	// copy(destination, source) 			--> menyalin slice dari source ke destination

	fmt.Println(len(slice1));
	fmt.Println(cap(slice1));
	
	// mengubah array akan merubah slice 
	// contoh : bulan[5];
	//          fmt.Println(slice1);
	// hasil  : [Mei Diubah Juli]

	// merubah slice akan merubah array 
	// contoh : slice[0] = "Mei lagi";
	// 			fmt.Println(bulan);
	// hasil  : [Januari Februari Maret April Mei lagi Juni Juli Agustus September Oktober November Desember]

	var slice2 = bulan[10:];
	fmt.Println(slice2); // [November Desember]

	slice3  := append(slice2, "ALDY");
	fmt.Println(slice3); // [November Desember ALDY]

	var slice4 = bulan[2:6];
	fmt.Println(slice4); // [Maret April Mei Juni]

	slice4[1] = "ALDY";
	fmt.Println(slice4); // [Maret ALDY Mei Juni]

	// make slice
	// make([]TypeData, length, capacity) --> membuat slice baru
	// newSlice := ([]string, 2, 5)

	newSlice := make([]string, 2, 5);
	newSlice [0] = "Aldy";
	newSlice [1] = "Imam";

	fmt.Println(newSlice);
	fmt.Println(len(newSlice));
	fmt.Println(cap(newSlice));

	// copy slice
	// copy(destination, source) --> menyalin slice dari source ke destination
	
	fromSlice := bulan[0:5];
	toSlice := make([]string, len(fromSlice), cap(fromSlice));

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	
	// beda Array dan Slice 
	iniArray := [...] int {1,2,3,4,5}; // panjang array diisi jika tahu, jika tidak diisi " ... "
	iniSlice := [] int {1,2,3,4,5}; // panjang array tidak diisi 
	
	fmt.Println(iniArray);
	fmt.Println(iniSlice);
}