package book

import "time"

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