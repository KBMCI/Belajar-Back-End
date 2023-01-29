package book

import (
	"encoding/json"
)

//  buat struct 
type BookInput struct {
	Title string 		`json:"title" binding:"required"` 			// binding:"required" 			--> artinya wajib diisi
	Price json.Number 	`json:"price" binding:"required,number"` 	// binding:"required,number" 	--> artinya selain wajib diisi juga harus berupa angka
	SubTitle string 	`json:"sub_title"` 							// SubTitle digunakan untuk menangkap json sub_title
	
	// json.Number berfungsi untuk mengkonversi string menjadi number 
	// contoh 1 : "20" 	--> 20
	// contoh 2 : "ada" --> error (karena bukan angka)
	
}	