package main

import "fmt";

func getPenjumlahan(angka ...int) int {
	
	total := 0;

	for _, angka := range angka {
		total += angka;
	};

	return total;

}

func main() {
	
	total := getPenjumlahan(10, 10, 10, 10, )
	fmt.Println(total)
}