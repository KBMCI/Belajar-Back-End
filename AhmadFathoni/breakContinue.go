package main

import "fmt"

func main() {
	// break untuk menghentikan perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Angka ke ", i)
	}

	fmt.Println()
	// continue untuk skip perulangan yang sedang berlangsung dengan kondisi yang sudah ditentukan contoh di baris 18
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println("Angka ke ", i)
	}
}