package main

import "fmt";

func getPenjumlahan(nama, kota string, angka ...int) int {
	
	total := 0;

	for _, angka := range angka {
		total += angka;
	};

	return total;

}

func main() {
	
	total := getPenjumlahan("Aldy", "Malang", 10, 10, 10, 10, )
	fmt.Println(total)
}