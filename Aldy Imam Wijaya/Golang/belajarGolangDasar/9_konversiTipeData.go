package main;


import "fmt";

func main() {

	// konversi type data 1 

	var nilai32 int32 = 200;
	var nilai16 int16 = int16(nilai32);	

	fmt.Println(nilai32);
	fmt.Println(nilai16);

	//nilai int yang besar tidak dapat ditampung pada int yang kecil 
	
	var nilai64 int64 = 215150700111039;
	var nilai8	int8 = int8(nilai64);

	fmt.Println(nilai64);
	fmt.Println(nilai8);

	// konversi type data 2 

	var namaDepan string = "Aldy";
	var e byte = namaDepan[0];
	var eString string = string(e);

	fmt.Println(namaDepan);
	fmt.Println(eString);
};