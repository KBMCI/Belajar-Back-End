package main;

import "fmt";

func getBiodata()(string, uint64, string)  {
	return "Aldy Imam Wijaya", 215150700111039, "Malang";
}

func main() {
	namaLengkap, nim, _:= getBiodata();
	fmt.Println("NAMA \t: ", namaLengkap);
	fmt.Println("NIM \t: ", nim);
	fmt.Println("KOTA \t:", ) // tidak bisa dieksekusi karena tanda " _ " berarti meng ignore
};