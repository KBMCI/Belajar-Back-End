package book

import "gorm.io/gorm"

type Repository interface {
	Get() ([]Book, error)
	GetById(ID int) (Book, error)
	Save(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error
	
	return books, err
}

func (r *repository) GetById(ID int) (Book, error) {
	var book Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Save(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}