package book

import "encoding/json"

// untuk menerima data yang dirikim melalui post dapat diterima menggunakan struct
type BookRequest struct {
	// Pemberian nama property harus sama dengan data yang di POST

	// json: -> Untuk mengidentifikasi nama pada json API
	// binding:"required" -> untuk menandakan field tersebut harus diisi
	// Jika saat POST data yang terdapat required pada attributenya tidak memberikan nilai maka akan terjadi error
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string
	Rating      int
	// Price    int    `json:"price" binding:"required,number"`
	// SubTitle string `json:"sub_title"` // jika beda dapat menggunakan cara ini
	// SubTitle string -> tidak bisa digunakan
	// Sub_Title string -> bisa digunakan
}
