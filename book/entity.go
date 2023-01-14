package book

import "time"

// Pembuatan tabel menggunakan struct
// Nama tabel menyesuaikan nama struct. Jika ditulis singular maka di db menjadi plural
type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
