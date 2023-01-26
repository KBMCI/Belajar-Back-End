package book

import (
	"encoding/json"
	"time"
)

// struct untuk menerima data dari body POST books postman
type BookInput struct {
	Title 		string		`json:"title" binding:"required"`
	// json.Number untuk validasi error jika yg diinputkan bukan number/untuk menghindari internal server error
	Price 		json.Number	`json:"price" binding:"required"`
	// untuk menangkap body json SubTitle
	SubTitle 	string 		`json:"sub_title" binding:"required"`
}

type Book struct {
	ID 			int
	Title 		string `gorm:"type:varchar(100)"` //merubah ke varchar 
	Description string `gorm:"size:255"` // set field varchar size to 255
	Price 		int
	Rating 		int
	Discount	int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}