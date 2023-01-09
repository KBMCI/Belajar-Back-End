package main

import (
	"fmt"
	"strings"
)

// strings.Trim(string, cutset) 		--> Memotong cutset di awal dan akhir string
// strings.ToLower(string) 				--> Membuat semua karakter string menjadi lower case
// strings.ToUpper(string) 				--> Membuat semua karakter string menjadi upper case
// strings.Split(string, separator) 	--> Memotong string berdasarkan separator
// strings.Contains(string, search) 	--> Mengecek apakah string mengandung string lain
// strings.ReplaceAll(string, from, to) --> Mengubah semua string dari from ke to

func main52() {
	
	// strings.Contains
	fmt.Println(strings.Contains("Aldy Imam Wijaya", "Aldy")) // hasil berupa boolean true
	fmt.Println(strings.Contains("Aldy Imam Wijaya", "Toni")) // hasil berupa boolean false

	// strings.Split 
	fmt.Println(strings.Split("Aldy Imam Wijaya", " ")) // hasil berupa type data slice 

	// strings.ToLower & strings.ToUpper
	fmt.Println(strings.ToLower("Aldy Imam Wijaya")) // hasil berupa "aldy imam wijaya"
	fmt.Println(strings.ToUpper("Aldy Imam Wijaya")) // hasil berupa "ALDY IMAM WIJAYA"

	//strings.Trim
	fmt.Println(strings.Trim("111 Aldy Imam Wijaya 111", "1")) // angka 1 hilang namun spasi setelah dan sebelum 1 tidak hilang
	fmt.Println(strings.Trim("111 Aldy Imam Wijaya 111", "1 ")) // angka 1 dan spasi hilang 

	//strings.ReplaceAll
	fmt.Println(strings.ReplaceAll("Aldy Aldy Aldy Aldy", "Aldy", "Toni")) // tulisan Aldy akan menjadi Toni semua
	fmt.Println(strings.ReplaceAll("Aldy Faza Budi Aldy", "Aldy", "Toni")) // tulisan Aldy akan menjadi Toni tapi tidak untuk faza dan Budi

}